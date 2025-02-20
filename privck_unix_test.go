//go:build !windows
// +build !windows

package privck

import (
	"os"
	"testing"
)

func TestIsElevatedUnix(t *testing.T) {
	// Test running as root
	if os.Geteuid() == 0 {
		if !IsElevated() {
			t.Error("Expected IsElevated() to return true when running as root")
		}
	} else {
		if IsElevated() {
			t.Error("Expected IsElevated() to return false when not running as root")
		}
	}
}
