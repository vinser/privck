// privck.go
package privck

// IsElevated checks if the program is running with elevated privileges.
func IsElevated() bool {
	return isElevated()
}
