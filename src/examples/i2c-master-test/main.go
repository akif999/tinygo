// I2C master (controller) test for nucleo-f401re.
// Connects to an I2C target at address 0x11 via I2C0 (PB8=SCL/D15, PB9=SDA/D14).
// Reads a message from the target, then writes and reads back a value.
// Results are printed via ST-Link VCP.
package main

import (
	"machine"
	"time"
)

const targetAddress = 0x11

func main() {
	i2c := machine.I2C0
	err := i2c.Configure(machine.I2CConfig{
		SCL: machine.I2C0_SCL_PIN,
		SDA: machine.I2C0_SDA_PIN,
	})
	if err != nil {
		println("I2C configure error:", err.Error())
		return
	}

	time.Sleep(500 * time.Millisecond)

	for {
		// read 12 bytes from target (address register 0)
		buf := make([]byte, 12)
		err = i2c.Tx(targetAddress, []byte{0}, buf)
		if err != nil {
			println("read error:", err.Error())
		} else {
			println("read:", string(buf))
		}

		// write value 42 to register 1
		err = i2c.Tx(targetAddress, []byte{1, 42}, nil)
		if err != nil {
			println("write error:", err.Error())
		} else {
			println("wrote 42 to reg 1")
		}

		time.Sleep(100 * time.Millisecond)

		// read back register 1
		rbuf := make([]byte, 1)
		err = i2c.Tx(targetAddress, []byte{1}, rbuf)
		if err != nil {
			println("readback error:", err.Error())
		} else {
			println("readback reg 1:", rbuf[0])
		}

		time.Sleep(2 * time.Second)
	}
}
