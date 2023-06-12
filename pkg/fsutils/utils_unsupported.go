//go:build !linux && !windows && !freebsd && !solaris && !darwin
// +build !linux,!windows,!freebsd,!solaris,!darwin

package fsutils

var (
	// Slice of drivers that should be used in an order
	DriverPriority = []string{
		"unsupported",
	}
)
