package tm1637

import (
	"machine"
	"time"
)

type device struct {
	Clk      machine.Pin
	Dio      machine.Pin
	Data     uint8
	Addr     uint8
	dispCtrl uint8
	point    bool
}

func New(clk, dio machine.Pin) *device {
	return &device{Clk: clk, Dio: dio}
}

func (d *device) Configure() {
	d.Clk.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.Dio.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (d *device) Init() {
	d.ClearDisplay()
}

const (
	AddrAuto  uint8 = 0x40
	AddrFixed uint8 = 0x44
)

func (d *device) Set(data, addr, brightness uint8) {
	d.Data = data
	d.Addr = addr
	d.dispCtrl = 0x88 + brightness
}

func (d *device) Display(data []uint8) {
	segData := make([]uint8, 4)
	copy(data, segData)

	d.coding(segData)
	d.start()
	d.writeByte(AddrAuto)
	d.stop()
	d.start()
	d.writeByte(d.Addr)
	for _, dt := range segData {
		d.writeByte(dt)
	}
	d.stop()
	d.start()
	d.writeByte(d.dispCtrl)
	d.stop()
}

func (d *device) DisplayWithBitAddr(addr, data uint8) {
	segData := d._coding(&data)
	d.start()
	d.writeByte(AddrFixed)
	d.stop()
	d.start()
	d.writeByte(addr | 0xC0)
	d.writeByte(segData)
	d.stop()
	d.start()
	d.writeByte(d.dispCtrl)
	d.stop()
}

func (d *device) Point(point bool) {
	d.point = point
}

func (d *device) ClearDisplay() {
	d.DisplayWithBitAddr(0x00, 0x7F)
	d.DisplayWithBitAddr(0x01, 0x7F)
	d.DisplayWithBitAddr(0x02, 0x7F)
	d.DisplayWithBitAddr(0x03, 0x7F)
}

func (d *device) ClearDisplayWithBitAddr(addr uint8) {
	d.DisplayWithBitAddr(addr, 0x7F)
}

func (d *device) writeByte(data uint8) bool {
	for i := 0; i < 8; i++ {
		d.Clk.Low()
		if (data & 0x01) == 0x01 {
			d.Dio.High()
		} else {
			d.Dio.Low()
		}
		data >>= 1
		d.Clk.High()
	}
	d.Clk.Low()
	d.Dio.High()
	d.Clk.High()
	d.Dio.Configure(machine.PinConfig{Mode: machine.PinInput})

	bitDelay()
	ack := d.Dio.Get()
	if ack == false {
		d.Dio.Configure(machine.PinConfig{Mode: machine.PinOutput})
		d.Dio.Low()
	}
	bitDelay()
	d.Dio.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bitDelay()

	return ack
}

func (d *device) start() {
	d.Clk.High()
	d.Dio.High()
	d.Dio.Low()
	d.Clk.Low()
}

func (d *device) stop() {
	d.Clk.Low()
	d.Dio.Low()
	d.Clk.High()
	d.Dio.High()
}

func (d *device) coding(dispData []uint8) {
	for i := 0; i < 4; i++ {
		_ = d._coding(&dispData[i])
	}
}

func (d *device) _coding(dispData *uint8) uint8 {
	tubeTab := []uint8{
		0x3f, 0x06, 0x5b, 0x4f,
		0x66, 0x6d, 0x7d, 0x07,
		0x7f, 0x6f, 0x77, 0x7c,
		0x39, 0x5e, 0x79, 0x71,
	} //0~9,A,b,C,d,E,F
	pointData := uint8(0x00)
	if d.point {
		pointData = 0x80
	}
	if *dispData == 0x7F {
		*dispData = 0x00 + pointData
	} else {
		*dispData = tubeTab[*dispData] + pointData
	}
	return *dispData
}

func bitDelay() {
	time.Sleep(time.Microsecond * 50)
}
