package tm1763

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

func (d *device) coding() {}

func bitDelay() {
	time.Sleep(time.Microsecond * 50)
}
