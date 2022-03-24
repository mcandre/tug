package tug

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

// TugBuilderName denotes the name used for the buildx builder.
const TugBuilderName = "tug"

// TugBuilderPattern denotes the pattern used to search for the tug builder within the buildx builder list.
var TugBuilderPattern = regexp.MustCompile(`^tug\W+`)

// DefaultPlatformsPattern denotes the pattern used to extract supported buildx platforms.
var DefaultPlatformsPattern = regexp.MustCompile(`Platforms:\W+(?P<platforms>.+)$`)

// PlatformPattern denotes the pattern used to extract operating system and architecture variants from buildx platform strings.
var PlatformPattern = regexp.MustCompile(`(?P<os>[^/]+)/(?P<arch>.+)$`)

// Platform models a targetable Docker image platform.
type Platform struct {
	// Os denotes a buildx operating system, e.g. "linux".
	Os string

	// Arch denotes a buildx architecture, e.g. "arm64".
	Arch string
}

// ParsePlatform extracts metadata from a buildx platform string.
func ParsePlatform(s string) (*Platform, error) {
	match := PlatformPattern.FindStringSubmatch(s)

	if len(match) < 3 {
		return nil, fmt.Errorf("invalid buildx platform: %v", s)
	}

	return &Platform{Os: match[1], Arch: match[2]}, nil
}

// Format renders a buildx platform string.
func (o Platform) Format() string {
	return fmt.Sprintf("%s/%s", o.Os, o.Arch)
}

// Platforms models a slice of platform(s).
type Platforms []Platform

// Len calculates the number of elements in a Platforms collection,
// in service of sorting.
func (o Platforms) Len() int {
	return len(o)
}

// Swap reverse the order of two elements in a Platforms collection,
// in service of sorting.
func (o Platforms) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// Less returns whether the elements a Platforms collection,
// identified by their indices,
// are in ascending order or not,
// in service of sorting.
func (o Platforms) Less(i int, j int) bool {
	return o[i].Format() < o[j].Format()
}

// EnsureTugBuilderExists creates a tug buildx builder when necessary.
func EnsureTugBuilderExists() error {
	cmd := exec.Command("docker")
	cmd.Args = []string{"docker", "buildx", "ls"}
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	stdoutChild, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdoutChild)

	var foundTugBuilder bool

	for scanner.Scan() {
		line := scanner.Text()

		if TugBuilderPattern.MatchString(line) {
			foundTugBuilder = true
			break
		}
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	if !foundTugBuilder {
		cmd := exec.Command("docker")
		cmd.Args = []string{"docker", "buildx", "create", "--name", TugBuilderName}
		cmd.Env = os.Environ()
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	return nil
}

// EnsureTugBuilderInUse activates the tug buildx builder.
func EnsureTugBuilderInUse() error {
	cmd := exec.Command("docker")
	cmd.Args = []string{"docker", "buildx", "use", TugBuilderName}
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// EnsureTugBuilderBootstrapped prepares the tug buildx builder for building.
func EnsureTugBuilderBootstrapped() error {
	cmd := exec.Command("docker")
	cmd.Args = []string{"docker", "buildx", "inspect", "--bootstrap"}
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// EnsureTugBuilder fully initializes the tug buildx builder.
func EnsureTugBuilder() error {
	if err := EnsureTugBuilderExists(); err != nil {
		return err
	}

	if err := EnsureTugBuilderInUse(); err != nil {
		return err
	}

	return EnsureTugBuilderBootstrapped()
}

// AvailablePlatforms initializes tug and reports the available buildx platforms.
func AvailablePlatforms() ([]Platform, error) {
	if err := EnsureTugBuilder(); err != nil {
		return nil, err
	}

	cmd := exec.Command("docker")
	cmd.Args = []string{"docker", "buildx", "inspect", TugBuilderName}
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	stdoutChild, err := cmd.StdoutPipe()

	if err != nil {
		return []Platform{}, err
	}

	if err2 := cmd.Start(); err2 != nil {
		return []Platform{}, err
	}

	scanner := bufio.NewScanner(stdoutChild)

	var platforms []Platform

	for scanner.Scan() {
		line := scanner.Text()
		match := DefaultPlatformsPattern.FindStringSubmatch(line)

		if len(match) < 2 {
			continue
		}

		platformsText := match[1]
		platformPairsText := strings.Split(platformsText, ", ")

		for _, platformPairText := range platformPairsText {
			platform, err2 := ParsePlatform(platformPairText)

			if err2 != nil {
				return platforms, err
			}

			platforms = append(platforms, *platform)
		}
	}

	if err2 := cmd.Wait(); err2 != nil {
		return platforms, err
	}

	if platforms == nil {
		return platforms, fmt.Errorf("no platforms detected")
	}

	sort.Sort(Platforms(platforms))

	return platforms, nil
}
