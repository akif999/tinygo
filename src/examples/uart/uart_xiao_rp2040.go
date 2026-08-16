//go:build xiao_rp2040

package main

import "machine"

var (
	uart = machine.UART0
	tx   = machine.UART0_TX_PIN
	rx   = machine.UART0_RX_PIN
)
