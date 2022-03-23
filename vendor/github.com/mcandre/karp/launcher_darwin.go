// +build darwin
package karp

import (
	"os/exec"
)

// Launch triggers a graphical hook for opening URIs.
func Launch(pth string) error {
	return exec.Command("open", pth).Run()
}
