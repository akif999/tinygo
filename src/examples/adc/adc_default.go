//go:build !nucleof401re

package main

import "machine"

var sensor = machine.ADC{Pin: machine.ADC2}
