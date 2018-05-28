package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdapter := firmata.NewAdaptor("COM3")
	led := gpio.NewLedDriver(firmataAdapter, "3")
	fadeLevel := uint8(0)
	fadeIn := true

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			fmt.Println("fadeLevel: ", fadeLevel)
			led.Brightness(byte(fadeLevel))

			if fadeIn {
				if fadeLevel >= 255 {
					fadeIn = !fadeIn
				} else {
					fadeLevel += uint8(5)
				}
			} else {
				if fadeLevel <= 0 {
					fadeIn = !fadeIn
				} else {
					fadeLevel -= uint8(5)
				}
			}
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
