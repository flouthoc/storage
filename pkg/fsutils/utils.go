package fsutils

import (
	"path/filepath"
	"os"
)


var Drivers = []string{
		"overlay",
		"aufs",
		"btrfs",
		"zfs",
		"vfs",
	}

// scanPriorDrivers returns an un-ordered scan of directories of prior storage drivers
func ScanPriorDrivers(root string) map[string]bool {
	driversMap := make(map[string]bool)

	for _, driver := range Drivers {
		p := filepath.Join(root, driver)
		if _, err := os.Stat(p); err == nil {
			driversMap[driver] = true
		}
	}
	return driversMap
}

