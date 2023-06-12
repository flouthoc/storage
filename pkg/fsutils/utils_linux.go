//go:build linux
// +build linux

package fsutils


var (
	// Slice of drivers that should be used in an order
	DriverPriority = []string{
		"overlay",
		// We don't support devicemapper without configuration
		// "devicemapper",
		"aufs",
		"btrfs",
		"zfs",
		"vfs",
	}
)
