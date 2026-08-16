package main

import (
	"machine"
	"time"
)

func main() {
	machine.InitADC()

	sensor.Configure(machine.ADCConfig{})

	for {
		val := sensor.Get()
		println(val)
		time.Sleep(time.Millisecond * 500)
	}
}
