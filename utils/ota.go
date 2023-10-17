package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"strings"

	log "github.com/ChrIgiSta/go-utils/logger"
	update "github.com/inconshreveable/go-update"
	"golang.org/x/net/html"
)

const (
	CURRENT_VERSION = "v0.0.2"
	DIST_URL        = "https://distro.easiot.ch/ota/apps/opel-bc/"
	BIN_ENDING      = "bin"
)

var (
	osArch     string
	goOs       string
	newVersion string
)

func init() {
	osArch = runtime.GOARCH
	goOs = runtime.GOOS
}

func IsUpdatable() bool {

	currentVSplit := strings.Split(strings.ReplaceAll(CURRENT_VERSION, "v", ""), ".")
	currentMajor, _ := strconv.Atoi(currentVSplit[0])
	currentMinor, _ := strconv.Atoi(currentVSplit[1])
	currentPatch, _ := strconv.Atoi(currentVSplit[2])

	log.Info("ota", "OS: %s, ARCH: %s", goOs, osArch)
	log.Info("ota", "not implemented")

	listResp, err := http.Get(DIST_URL + goOs + "/" + osArch + "/")
	if err != nil {
		log.Error("ota", "check for binary: %v", err)
		return false
	}
	defer listResp.Body.Close()

	if listResp.StatusCode != 200 {
		log.Error("ota", "couldn't get dist list. Code !=200")
		return false
	}

	listB, err := io.ReadAll(listResp.Body)
	if err != nil {
		log.Error("ota", "read in dist list: %v", err)
		return false
	}

	list, err := getLinksFromHtml(listB)
	if err != nil {
		log.Error("ota", "cannot extract list from html: %v", err)
		return false
	}

	fmt.Println(list)

	for _, file := range list {
		split := strings.Split(file, ".")
		if len(split) == 3 {
			if split[2] == BIN_ENDING {
				log.Info("otp", "binary found, check version: %v", file)

				versionSplit := strings.Split(strings.ReplaceAll(split[1], "v", ""), "-")

				if len(versionSplit) != len(currentVSplit) {
					log.Error("otp", "different version formats. skip... %v != %v", versionSplit, currentMajor)
					return false
				}
				// major:minor:patch
				major, _ := strconv.Atoi(versionSplit[0])
				minor, _ := strconv.Atoi(versionSplit[1])
				patch, _ := strconv.Atoi(versionSplit[2])

				if currentMajor < major {
					newVersion = file
					log.Info("ota", "major release found.")
					return true
				} else if currentMajor == major {
					if currentMinor < minor {
						newVersion = file
						log.Info("ota", "minor release found.")
						return true
					} else if currentMinor == minor {
						if currentPatch < patch {
							newVersion = file
							log.Info("ota", "patch found.")
							return true
						}
					}
				}
			}
		}
	}

	return false
}

func Update() error {

	if newVersion == "" {
		return errors.New("no check or no newer version is available")
	}

	log.Info("ota", "OS: %s, ARCH: %s", goOs, osArch)
	log.Info("ota", "not implemented")

	binaryResp, err := http.Get(DIST_URL + goOs + "/" + osArch + "/" + newVersion)
	if err != nil {
		log.Error("ota", "downloading binary: %v", err)
		return errors.New("downloading binary")
	}
	defer binaryResp.Body.Close()

	log.Info("otp", "bin loaded. apply")

	err = update.Apply(binaryResp.Body, update.Options{
		TargetPath: "", // itself
		Checksum:   nil,
		PublicKey:  nil,
		Signature:  nil,
	})
	if err != nil {
		log.Error("ota", "applying update: %v", err)
	}

	return err
}

func getLinksFromHtml(htmlContent []byte) ([]string, error) {
	var hrefs []string

	doc, err := html.Parse(strings.NewReader(string(htmlContent)))
	if err != nil {
		return nil, err
	}

	var extractHrefs func(*html.Node)
	extractHrefs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					hrefs = append(hrefs, attr.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractHrefs(c)
		}
	}

	extractHrefs(doc)

	return hrefs, err
}
