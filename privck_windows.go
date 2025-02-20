//go:build windows
// +build windows

package privck

import (
	"golang.org/x/sys/windows"
)

func isElevated() bool {
	return windows.GetCurrentThreadToken().IsElevated()
}
