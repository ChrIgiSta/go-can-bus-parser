package main

import (
	"GMCanDecoder/gui"
	"math/rand"
	"time"
)

func main() {
	g := gui.NewOriginalGui()

	go func() {
		consumption := 5.5
		currentCons := 10.3
		rangeV := 400
		fan := 0
		acTemp := 18
		coolantTemp := 60
		acMode := 0
		eco := true

		for true {
			g.UpdateTime(time.Now())
			g.UpdateConsumtion(float32(consumption), "l/100km")
			g.UpdateCurrentConsumtion(float32(currentCons), "l/h")
			g.UpdateRange(rangeV, "km")
			g.UpdateClimaFanSpeed(fan)
			g.UpdateClimaTemperature(acTemp, "°C")
			g.UpdateCoolantTemp(coolantTemp, "°C")
			g.UpdateEcoMode(eco)

			g.UpdateClimaMode(gui.AcMode(acMode))

			if rand.Float32() > 0.5 {
				eco = true

				coolantTemp++
				acTemp++
			} else {
				eco = false
				rangeV--
				fan++
				acMode++
			}

			if acMode > 7 {
				acMode = 0
			}
			if rangeV == 0 {
				rangeV = 400
			}
			if fan > 6 {
				fan = 0
			}
			if acTemp > 25 {
				acTemp = 18
			}
			if coolantTemp > 105 {
				coolantTemp = 87
			}

			consumption = rand.Float64() * 10.3
			currentCons = rand.Float64() * 20.8

			time.Sleep(1 * time.Second)
		}
	}()

	g.ShowDefaultView()
}
