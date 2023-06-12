//go:build freebsd
// +build freebsd

package fsutils


var (
	// Slice of drivers that should be used in an order
	DriverPriority = []string{
		"zfs",
		"vfs",
	}
)
