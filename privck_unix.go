//go:build !windows
// +build !windows

package privck

import "os"

func isElevated() bool {
	return os.Geteuid() == 0
}
