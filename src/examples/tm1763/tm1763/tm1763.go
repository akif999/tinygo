package tm1763

import (
	"machine"
	"time"
)

type device struct {
	Clk        machine.Pin
	Dio        machine.Pin
	Data       uint8
	Addr       uint8
	dispCtrl   uint8
	point      bool
	brightness uint8
}

func New(clk, dio uint8) *device {
	return &device{Clk: clk, Dio: dio}
}

func (d *device) Configure() {
	d.Clk.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.Dio.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (d *device) Init() {
	d.ClearDisplay()
}

func (d *device) Set(data, addr, brightness uint8) {
	d.Data = data
	d.Addr = addr
	d.brightness = 0x88 + brightness
}
func (d *device) Display() {}

func (d *device) DisplayWithBitAddr() {}

func (d *device) Point(point bool) {
	d.point = point
}

func (d *device) ClearDisplay() {}

func (d *device) writeByte() {}

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
