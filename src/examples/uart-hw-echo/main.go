// Hardware UART echo for xiao-rp2040.
// Echoes data received on UART0 (D7=RX/GPIO1) back on UART0 (D6=TX/GPIO0).
// Use with uart-peer-test on the peer board.
//
//go:build xiao_rp2040

package main

import "machine"

func main() {
	uart := machine.UART0
	uart.Configure(machine.UARTConfig{
		TX:       machine.UART0_TX_PIN,
		RX:       machine.UART0_RX_PIN,
		BaudRate: 115200,
	})
	for {
		if uart.Buffered() > 0 {
			b, _ := uart.ReadByte()
			uart.WriteByte(b)
		}
	}
}
