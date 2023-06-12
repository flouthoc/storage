//go:build windows
// +build windows

package fsutils


var (
	// Slice of drivers that should be used in an order
	DriverPriority = []string{
		"windowsfilter",
	}
)
