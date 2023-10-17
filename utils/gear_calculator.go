/**
 * Copyright © 2023, Staufi Tech - Switzerland
 * All rights reserved.
 *
 *   ________________________   ___ _     ________________  _  ____
 *  / _____  _  ____________/  / __|_|   /_______________  | | ___/
 * ( (____ _| |_ _____ _   _ _| |__ _      | |_____  ____| |_|_
 *  \____ (_   _|____ | | | (_   __) |     | | ___ |/ ___)  _  \
 *  _____) )| |_/ ___ | |_| | | |  | |     | | ____( (___| | | |
 * (______/  \__)_____|____/  |_|  |_|     |_|_____)\____)_| |_|
 *
 *
 *  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 *  AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 *  IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 *  ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 *  LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 *  CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 *  SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 *  INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 *  CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 *  ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 *  POSSIBILITY OF SUCH DAMAGE.
 */

package utils

import (
	"errors"
)

// Opel Astra H OPC - Predefinitions // ToDo -> to a "database"
const GEAR_CALC_RATIO_TO_RPM_PER_KMH = 31 // Astra H OPC // rmp / km/h = radio * 31

var (
	AstraHOpcGears []Gear = []Gear{
		{Gear: 1, Name: "1", Ratio: 3.82},
		{Gear: 2, Name: "2", Ratio: 2.16},
		{Gear: 3, Name: "3", Ratio: 1.48},
		{Gear: 4, Name: "4", Ratio: 1.07},
		{Gear: 5, Name: "5", Ratio: 0.88},
		{Gear: 6, Name: "6", Ratio: 0.74},
		{Gear: -1, Name: "R", Ratio: 3.55},
	}
)

// END Opel Astra H OPC - Predefinitions

const (
	GEAR_CALC_UNKNOWN_GEAR = -2
	GEAR_CALC_UNKNOWN_NAME = "unknown"
)

type Gear struct {
	Gear  int
	Name  string
	Ratio float32
}

type GearCalculator struct {
	gears           []Gear
	tolerance       float32
	rpmPerKmhRation []float32
	currentGear     string
}

var (
	NeutralGear Gear = Gear{
		Gear:  0,
		Name:  "N",
		Ratio: 0,
	}
)

func NewGearCalculator(gears []Gear, ratioToRpmPerKmh float32, tolerance float32) *GearCalculator {
	var rpmPerKmhRatio []float32 = []float32{}

	for _, gear := range gears {
		rpmPerKmhRatio = append(rpmPerKmhRatio, gear.Ratio*ratioToRpmPerKmh)
	}
	return &GearCalculator{
		gears:           gears,
		tolerance:       tolerance / 2,
		rpmPerKmhRation: rpmPerKmhRatio,
	}
}

func (g *GearCalculator) Get(rpm float32, speed float32) (*Gear, error) {

	if speed == 0 {
		return &NeutralGear, nil
	}

	rpmPerKmh := rpm / speed

	// fmt.Println("calc: ", rpmPerKmh)

	for gearIndex, fRatio := range g.rpmPerKmhRation {
		// fmt.Println("gear ", gearIndex, " ratio is from ", fRatio-g.tolerance, "to ", fRatio+g.tolerance)

		if fRatio-g.tolerance < rpmPerKmh && fRatio+g.tolerance > rpmPerKmh {
			return &g.gears[gearIndex], nil
		}
	}

	return nil, errors.New("gear not solvable")
}
