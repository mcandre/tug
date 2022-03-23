package tug

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// Job models a Docker muliti-image build operation.
type Job struct {
	// Debug can enable additional logging.
	Debug bool

	// Push can push cached buildx images to the remote Docker registry
	// as a side effect during builds.
	Push bool

	// Builder denotes a buildx builder.
	Builder string

	// LoadPlatform can load the image for a given platform
	// onto the local Docker registry as a side effect during builds.
	LoadPlatform *string

	// Platforms denotes the list of targeted image platforms.
	Platforms []Platform

	// OsExclusions skips the given operating systems.
	OsExclusions []string

	// ArchExclusions skips the given architectures.
	ArchExclusions []string

	// ListImageName can query the buildx cache
	// for any multi-platform images matching the given image name,
	// of the form name[:tag].
	ListImageName *string

	// ImageName denotes the buildx image artifact name,
	// of the form name[:tag].
	ImageName *string

	// Directory denotes the Docker build directory (defaults behavior assumes the current working directory).
	Directory string
}

// NewJob generates a default Job.
func NewJob() (*Job, error) {
	platforms, err := AvailablePlatforms()

	if err != nil {
		return nil, err
	}

	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	return &Job{Builder: TugBuilderName, Platforms: platforms, Directory: cwd}, nil
}

// Run executes a Job.
func (o Job) Run() error {
	cmd := exec.Command("docker")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Args = []string{"docker", "buildx"}

	if o.ListImageName != nil {
		cmd.Args = append(cmd.Args, "imagetools", "inspect", *o.ListImageName)
		return cmd.Run()
	}

	cmd.Args = append(cmd.Args, "build")

	var platformPairs []string

	for _, platform := range o.Platforms {
		var excludedOs bool

		for _, osExclusion := range o.OsExclusions {
			if platform.Os == osExclusion {
				excludedOs = true
				break
			}
		}

		if excludedOs {
			continue
		}

		var excludedArch bool

		for _, archExclusion := range o.ArchExclusions {
			if platform.Arch == archExclusion {
				excludedArch = true
				break
			}
		}

		if excludedArch {
			continue
		}

		platformPairs = append(platformPairs, platform.Format())
	}

	cmd.Args = append(cmd.Args, "--platform")

	if o.LoadPlatform == nil {
		cmd.Args = append(cmd.Args, strings.Join(platformPairs, ","))
	} else {
		cmd.Args = append(cmd.Args, *o.LoadPlatform)
		cmd.Args = append(cmd.Args, "--load")
	}

	if o.Push {
		cmd.Args = append(cmd.Args, "--push")
	}

	cmd.Args = append(cmd.Args, "-t")
	cmd.Args = append(cmd.Args, *o.ImageName)

	cmd.Args = append(cmd.Args, o.Directory)

	if o.Debug {
		log.Printf("Command: %v\n", cmd)
	}

	return cmd.Run()
}
