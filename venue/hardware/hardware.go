// Package hardware provides constants for the different IO devices.
package hardware

// Hardware defines the type of hardware.
type Hardware int

//go:generate stringer -type=Hardware

const (
	StageBox Hardware = iota
	Local
	ProTools
)
