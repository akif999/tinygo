// SPI loopback test for nucleo-f401re.
// Connect MOSI (PA7/D11) to MISO (PA6/D12) with a jumper wire.
// Sends a test pattern and verifies the received data matches.
// Results are printed via ST-Link VCP.
package main

import (
	"machine"
	"time"
)

func main() {
	spi := machine.SPI0
	err := spi.Configure(machine.SPIConfig{
		SCK:       machine.SPI0_SCK_PIN, // PA5 / D13
		SDO:       machine.SPI0_SDO_PIN, // PA7 / D11 (MOSI)
		SDI:       machine.SPI0_SDI_PIN, // PA6 / D12 (MISO)
		Frequency: 1_000_000,            // 1 MHz
		Mode:      0,
	})
	if err != nil {
		println("SPI configure error:", err.Error())
		return
	}

	tx := []byte{0x01, 0x02, 0x03, 0x04, 0xAA, 0x55, 0xFF, 0x00}
	rx := make([]byte, len(tx))

	for {
		err = spi.Tx(tx, rx)
		if err != nil {
			println("SPI Tx error:", err.Error())
		} else {
			match := true
			for i := range tx {
				if tx[i] != rx[i] {
					match = false
					break
				}
			}
			if match {
				println("SPI loopback OK: tx==rx")
			} else {
				print("SPI mismatch! tx:")
				for _, b := range tx {
					print(" ", b)
				}
				print("  rx:")
				for _, b := range rx {
					print(" ", b)
				}
				println()
			}
		}
		time.Sleep(time.Second)
	}
}
