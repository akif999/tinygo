package main

import (
	"machine"
	"time"

	"github.com/sago35/tinygo/src/examples/tm1637/tm1637"
)

func main() {
	device := tm1637.New(machine.ADC5, machine.ADC4)
	device.Configure()
	device.Init()

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		device.Point(true)
		device.DisplayWithBitAddr(0, 2)
		device.DisplayWithBitAddr(1, 3)
		device.DisplayWithBitAddr(2, 5)
		device.DisplayWithBitAddr(3, 5)
		time.Sleep(time.Millisecond * 500)

		led.High()
		device.Point(false)
		device.ClearDisplay()
		time.Sleep(time.Millisecond * 500)
	}
}
