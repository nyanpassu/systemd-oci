package systemctl

import (
	"os/exec"
)

// Start .
func Start(name string) error {
	return exec.Command("systemctl", "start", "name").Run()
}

// Stop .
func Stop(string) error {
	return exec.Command("systemctl", "stop", "name").Run()
}

// Enable .
func Enable(string) error {
	return exec.Command("systemctl", "enable", "name").Run()
}

// Disable .
func Disable(string) error {
	return exec.Command("systemctl", "disable", "name").Run()
}
