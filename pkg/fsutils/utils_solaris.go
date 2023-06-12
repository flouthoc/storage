//go:build solaris && cgo
// +build solaris,cgo

package fsutils


var (
	// Slice of drivers that should be used in an order
	DriverPriority = []string{
		"zfs",
	}
)
