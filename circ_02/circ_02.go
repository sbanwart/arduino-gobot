package main

import (
	"fmt"
	"strconv"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdapter := firmata.NewAdaptor("COM3")
	pinNum := 2
	led := 0
	directionOn := true

	var leds [8]*gpio.LedDriver
	for i := 0; i < 8; i++ {
		pin := strconv.Itoa(pinNum)
		leds[i] = gpio.NewLedDriver(firmataAdapter, pin)

		pinNum = pinNum + 1
	}

	work := func() {
		gobot.Every(1*time.Second, func() {
			fmt.Println("led: ", led, ", directionOn: ", directionOn)
			leds[led].Toggle()
			if directionOn {
				if led == 7 {
					directionOn = !directionOn
				} else {
					led = led + 1
				}
			} else {
				if led == 0 {
					directionOn = !directionOn
				} else {
					led = led - 1
				}
			}
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{leds[0]},
		[]gobot.Device{leds[1]},
		[]gobot.Device{leds[2]},
		[]gobot.Device{leds[3]},
		[]gobot.Device{leds[4]},
		[]gobot.Device{leds[5]},
		[]gobot.Device{leds[6]},
		[]gobot.Device{leds[7]},
		work,
	)

	robot.Start()
}
