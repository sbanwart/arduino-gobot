package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

// LOW signal value for writing to a digital interface pin
const LOW = 0

// HIGH signal value for writing to a digital interface pin
const HIGH = 1

func main() {
	firmataAdapter := firmata.NewAdaptor("COM3")
	speakerPin := "9"

	// Song data
	songLength := 15
	notes := []string{"c", "c", "g", "g", "a", "a", "g", "f", "f", "e", "e", "d", "d", "c", " "}
	beats := []int{1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 2, 4}
	tempo := 300

	speaker := gpio.NewDirectPinDriver(firmataAdapter, speakerPin)
	position := 0

	work := func() {
		gobot.Every(time.Duration(tempo/2)*time.Millisecond, func() {
			if notes[position] == " " {
				time.Sleep(time.Duration(beats[position]*tempo) * time.Millisecond)
			} else {
				playNote(notes[position], beats[position]*tempo, speaker)
			}
			position++
			if position > songLength-1 {
				position = 0
			}
		})
	}

	robot := gobot.NewRobot("musicbot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{speaker},
		work,
	)

	robot.Start()
}

func playTone(tone int, duration int, speaker *gpio.DirectPinDriver) {
	fmt.Println("tone: ", tone, ", duration: ", duration)
	for i := 0; i < (duration * 1000); i += (int(tone) * 2) {
		speaker.DigitalWrite(HIGH)
		time.Sleep(time.Duration(tone) * time.Microsecond)
		speaker.DigitalWrite(LOW)
		time.Sleep(time.Duration(tone) * time.Microsecond)
	}
}

func playNote(note string, duration int, speaker *gpio.DirectPinDriver) {
	names := []string{"c", "d", "e", "f", "g", "a", "b", "c"}
	tones := []int{1915, 1700, 1519, 1432, 1275, 1136, 1014, 956}

	for i := 0; i < 8; i++ {
		if names[i] == note {
			playTone(tones[i], duration, speaker)
		}
	}
}
