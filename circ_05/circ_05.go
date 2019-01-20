package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

const LOW = 0
const HIGH = 1
const OFF = LOW
const ON = HIGH

func main() {
	firmataAdapter := firmata.NewAdaptor("COM3")
	dataPin := "2"
	clockPin := "3"
	latchPin := "4"

	data := gpio.NewDirectPinDriver(firmataAdapter, dataPin)
	clock := gpio.NewDirectPinDriver(firmataAdapter, clockPin)
	latch := gpio.NewDirectPinDriver(firmataAdapter, latchPin)

	value := uint8(0)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			fmt.Println("value: ", value)
			updateLEDs(value, data, clock, latch)
			value++
			if value > 255 {
				fmt.Println("Limit reached, reset to zero")
				value = 0
			}
		})
	}

	robot := gobot.NewRobot("icbot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{data},
		[]gobot.Device{clock},
		[]gobot.Device{latch},
		work,
	)

	robot.Start()
}

// Manually calculate each bit state as I don't see anything in GoBot that matches the Arduino shiftOut function
func updateLEDs(value uint8, data *gpio.DirectPinDriver, clock *gpio.DirectPinDriver, latch *gpio.DirectPinDriver) {
	latch.DigitalWrite(LOW)

	for i := 0; i < 8; i++ {
		var bit = value & 0x80
		value = value << 1

		if bit == 0x80 {
			data.DigitalWrite(HIGH)
		} else {
			data.DigitalWrite(LOW)
		}

		// Pulse the clock pin
		clock.DigitalWrite(HIGH)
		time.Sleep(1 * time.Millisecond)
		clock.DigitalWrite(LOW)
	}

	latch.DigitalWrite(HIGH)
}
