// +build !darwin

package karp

import (
	"io/ioutil"
	"os/exec"
	"regexp"
)

// linuxLaunch triggers a graphical hook for opening URIs.
func linuxLaunch(pth string) error {
	return exec.Command("xdg-open", pth).Run()
}

// windowsLaunch triggers a graphical hook for opening URIs.
func windowsLaunch(pth string) error {
	return exec.Command("powershell.exe", "Start", pth).Run()
}

// brandPattern identifies COMSPEC/WSL environments.
var brandPattern = regexp.MustCompile("[Mm]icrosoft")

// Launch triggers a graphical hook for opening URIs.
func Launch(pth string) error {
	versionPath := "/proc/version"

	versionBytes, err := ioutil.ReadFile(versionPath)

	if err != nil || !brandPattern.Match(versionBytes) {
		return linuxLaunch(pth)
	}

	return windowsLaunch(pth)
}
