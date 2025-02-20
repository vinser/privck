//go:build windows
// +build windows

package privck

import (
	"testing"

	"golang.org/x/sys/windows"
)

func TestIsElevatedWindows(t *testing.T) {
	isAdmin := windows.GetCurrentThreadToken().IsElevated()

	if IsElevated() != isAdmin {
		t.Errorf("Expected IsElevated() to return %v, got %v", isAdmin, IsElevated())
	}
}
