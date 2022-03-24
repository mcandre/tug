package tug

import (
	"fmt"
	"os"
	"os/exec"
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

// Clean empties the active buildx image cache
// and removes the tug builder,
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

	return status
}
