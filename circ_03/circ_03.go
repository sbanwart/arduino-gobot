package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdapter := firmata.NewAdaptor(os.Args[1])
	motor := gpio.NewMotorDriver(firmataAdapter, "9")
	speed := byte(100)

	work := func() {
		gobot.Every(10*time.Second, func() {
			motor.On()
			fmt.Println("Setting speed...")
			motor.Speed(speed)
			fmt.Println("Sleeping...")
			duration, _ := time.ParseDuration("5s")
			time.Sleep(duration)
			fmt.Println("Shutting down...")
			motor.Off()
		})
	}

	robot := gobot.NewRobot("motorBot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{motor},
		work,
	)

	robot.Start()
}
