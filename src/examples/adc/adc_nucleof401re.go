//go:build nucleof401re

package main

import "machine"

// A0 = PA0 = ADC1_IN0
var sensor = machine.ADC{Pin: machine.A0}
