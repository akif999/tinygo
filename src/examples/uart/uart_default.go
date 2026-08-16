//go:build !xiao_rp2040

package main

import "machine"

var (
	uart = machine.UART1
	tx   = machine.UART1_TX_PIN
	rx   = machine.UART1_RX_PIN
)
