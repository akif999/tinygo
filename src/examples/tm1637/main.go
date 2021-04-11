package main

import (
	"machine"
	"time"

	"github.com/sago35/tinygo/src/examples/tm1637/tm1637"
)

func main() {
	// Initial TM1637
	device := tm1637.New(machine.ADC5, machine.ADC4)
	device.Configure()
	device.Init()
	device.Set(0x40, 0xC0, 0x02)

	// Initial LED
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Initial siwtches
	swGrn := machine.D4
	swGrn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	swBlu := machine.D0
	swBlu.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	swBlk := machine.D2
	swBlk.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	swWht := machine.D3
	swWht.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	swRed := machine.D5
	swRed.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	//Initial digital out port
	relaySwitchOut := machine.D6
	relaySwitchOut.Configure(machine.PinConfig{Mode: machine.PinOutput})
	relaySwitchOut.High()

	device.Point(true)
	// 10ms
	device.DisplayWithBitAddr(0, 1)
	device.DisplayWithBitAddr(1, 0)
	// 5times
	device.DisplayWithBitAddr(2, 0)
	device.DisplayWithBitAddr(3, 5)

	segParam := struct {
		energizationTime  uint8 // lsb: 10ms
		numOfEnergization uint8
		editingDigit      uint8
	}{
		10,
		5,
		0,
	}
	for {
		led.Low()

		device.DisplayWithBitAddr(0, segParam.energizationTime/10)
		device.DisplayWithBitAddr(1, segParam.energizationTime%10)
		device.DisplayWithBitAddr(2, segParam.numOfEnergization/10)
		device.DisplayWithBitAddr(3, segParam.numOfEnergization%10)

		time.Sleep(time.Millisecond * 200)

		led.High()

		device.ClearDisplayWithBitAddr(segParam.editingDigit)

		time.Sleep(time.Millisecond * 200)

		// shift flashing digti right
		if !swBlu.Get() {
			if segParam.editingDigit < 3 {
				segParam.editingDigit++
			} else if segParam.editingDigit == 3 {
				segParam.editingDigit = 0x00
			} else {
				// nothing to do
			}
		}
		// shift flashing digti left
		if !swBlk.Get() {
			if 0 < segParam.editingDigit {
				segParam.editingDigit--
			} else if segParam.editingDigit == 0 {
				segParam.editingDigit = 0x03
			} else {
				// nothing to do
			}
		}

		// increment prameter
		if !swWht.Get() {
			switch segParam.editingDigit {
			case 0x00:
				if segParam.energizationTime < 90 {
					segParam.energizationTime += 10
				} else if 90 <= segParam.energizationTime {
					segParam.energizationTime -= 90
				} else {
					// nothing to do
				}
			case 0x01:
				if segParam.energizationTime < 99 {
					segParam.energizationTime++
				} else if segParam.energizationTime == 99 {
					segParam.energizationTime = 0
				} else {
					// nothing to do
				}
			case 0x02:
				// never editied
			case 0x03:
				if segParam.numOfEnergization < 9 {
					segParam.numOfEnergization++
				} else if segParam.numOfEnergization == 9 {
					segParam.numOfEnergization = 0
				} else {
					// nothing to do
				}
			default:
				panic("invalid digit")
			}
		}

		// decrement prameter
		if !swGrn.Get() {
			switch segParam.editingDigit {
			case 0x00:
				if 9 < segParam.energizationTime {
					segParam.energizationTime -= 10
				} else if segParam.energizationTime <= 9 {
					segParam.energizationTime += 90
				} else {
					// nothing to do
				}
			case 0x01:
				if 0 < segParam.energizationTime {
					segParam.energizationTime--
				} else if segParam.energizationTime == 0 {
					segParam.energizationTime = 99
				} else {
					// nothing to do
				}
			case 0x02:
				// never editied
			case 0x03:
				if 0 < segParam.numOfEnergization {
					segParam.numOfEnergization--
				} else if segParam.numOfEnergization == 0 {
					segParam.numOfEnergization = 9
				} else {
					// nothing to do
				}
			default:
				panic("invalid digit")
			}
		}

		// energize
		if !swRed.Get() {
			// param 0 is max number of parameter
			t := uint16(uint16(segParam.energizationTime) * 10)
			n := segParam.numOfEnergization
			if t == 0 {
				t = 1000
			}
			if n == 0 {
				n = 10
			}
			println(segParam.energizationTime)
			println(t)
			for i := uint8(0); i < n; i++ {
				device.DisplayWithBitAddr(0, segParam.energizationTime/10)
				device.DisplayWithBitAddr(1, segParam.energizationTime%10)
				device.DisplayWithBitAddr(2, segParam.numOfEnergization/10)
				device.DisplayWithBitAddr(3, segParam.numOfEnergization%10)
				relaySwitchOut.Low()
				time.Sleep(time.Millisecond * time.Duration(t+15)) //15ms is adjustment
				device.ClearDisplay()
				relaySwitchOut.High()
				time.Sleep(time.Millisecond * time.Duration(t+15)) //15ms is adjustment
			}
		}
	}
}
