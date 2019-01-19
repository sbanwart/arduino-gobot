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
	position := 0
	increment := 1
	servoPin := "9"

	servo := gpio.NewServoDriver(firmataAdapter, servoPin)
	//firmataAdapter.ServoConfig(servoPin, 0, 180)

	work := func() {
		gobot.Every(15*time.Millisecond, func() {
			fmt.Println("servo position: ", position)
			firmataAdapter.ServoWrite(servoPin, byte(position))
			position += increment

			if position > 179 {
				increment *= -1
			} else if position < 1 {
				increment *= -1
			}
		})
	}

	robot := gobot.NewRobot("servobot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{servo},
		work,
	)

	robot.Start()
}
