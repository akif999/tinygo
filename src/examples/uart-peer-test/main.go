// UART peer communication test for nucleo-f401re.
// Sends "hello\r\n" via UART2 (PA9=TX, PA10=RX) every second and prints
// any echoed data back to the ST-Link VCP (UART1, PA2/PA3).
//
// Peer board setup: flash examples/echo to xiao-rp2040 (UART0 = D6/D7).
// Wiring:
//   NUCLEO PA9  (D8) → xiao D7 (RX)
//   NUCLEO PA10 (D2) → xiao D6 (TX)
//   GND ─ GND
package main

import (
	"machine"
	"time"
)

func main() {
	peer := machine.UART2
	peer.Configure(machine.UARTConfig{
		TX:       machine.UART2_TX_PIN,
		RX:       machine.UART2_RX_PIN,
		BaudRate: 115200,
	})

	vcp := machine.Serial // ST-Link VCP

	for {
		// send test message to peer
		peer.Write([]byte("hello\r\n"))
		vcp.Write([]byte("sent: hello\r\n"))

		// collect echo for 200ms
		time.Sleep(200 * time.Millisecond)
		if peer.Buffered() > 0 {
			vcp.Write([]byte("recv: "))
			for peer.Buffered() > 0 {
				b, _ := peer.ReadByte()
				vcp.WriteByte(b)
			}
		} else {
			vcp.Write([]byte("recv: (nothing)\r\n"))
		}

		time.Sleep(800 * time.Millisecond)
	}
}
