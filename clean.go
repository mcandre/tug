package tug

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

// RemoveBuildxImageCache deletes any images/layers in the active buildx cache.
func RemoveBuildxImageCache() error {
	cmd := exec.Command("docker")
	cmd.Args = []string{"docker", "buildx", "prune", "--force", "--all"}
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// RemoveTugBuilder deletes the tug buildx builder.
func RemoveTugBuilder() error {
	cmd := exec.Command("docker")
	cmd.Args = []string{"docker", "buildx", "rm", TugBuilderName}
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// TerminateQemuProcesses ends any qemu processes.
func TerminateQemuProcesses() error {
	processes, err := ps.Processes()

	if err != nil {
		return err
	}

	for _, process := range processes {
		executable := process.Executable()

		if strings.HasPrefix(executable, "qemu-") {
			p, err := os.FindProcess(process.Pid())

			if err != nil {
				return err
			}

			if err2 := p.Kill(); err2 != nil {
				return err2
			}
		}
	}

	return nil
}

// Clean empties the active buildx image cache,
// removes the tug builder,
// and terminates any qemu processes.
//
// Returns zero on successful operation. Otherwise, returns non-zero.
func Clean() int {
	var status = 0

	if err := RemoveBuildxImageCache(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		status = 1
	}

	if err := RemoveTugBuilder(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		status = 1
	}

	if err := TerminateQemuProcesses(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		status = 1
	}

	return status
}
