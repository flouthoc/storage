//go:build darwin
// +build darwin

package fsutils


var (
	// Slice of drivers that should be used in an order
	DriverPriority = []string{
		"zfs",
		"vfs",
	}
)
