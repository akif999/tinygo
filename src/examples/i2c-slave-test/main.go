// I2C target (slave) for xiao-rp2040.
// Listens at address 0x11 on I2C1 (D5=SCL/GPIO7, D4=SDA/GPIO6).
// Implements a simple 256-byte memory map.
//
//go:build xiao_rp2040

package main

import (
	"machine"
)

const (
	targetAddress = 0x11
	maxTxSize     = 16
)

var mem [256]byte

func main() {
	copy(mem[0:], []byte("Hello NUCLEO!"))

	i2c := machine.I2C1
	err := i2c.Configure(machine.I2CConfig{
		Mode: machine.I2CModeTarget,
		SCL:  machine.I2C1_SCL_PIN, // D5 / GPIO7
		SDA:  machine.I2C1_SDA_PIN, // D4 / GPIO6
	})
	if err != nil {
		println("I2C configure error:", err.Error())
		return
	}

	err = i2c.Listen(targetAddress)
	if err != nil {
		println("I2C listen error:", err.Error())
		return
	}

	buf := make([]byte, maxTxSize)
	var ptr uint8

	for {
		evt, n, err := i2c.WaitForEvent(buf)
		if err != nil {
			println("event error:", err.Error())
			continue
		}

		switch evt {
		case machine.I2CReceive:
			if n > 0 {
				ptr = buf[0]
			}
			for o := 1; o < n; o++ {
				mem[ptr] = buf[o]
				ptr++
			}
		case machine.I2CRequest:
			i2c.Reply(mem[ptr:])
		case machine.I2CFinish:
			// nothing
		}
	}
}
