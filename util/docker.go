package util

import (
	"fmt"
	"os"
	"os/exec"
)

func DockerImageTag(hash string, suffix string) string {
	tag := fmt.Sprintf("kbst:%s", hash)
	if suffix != "" {
		tag = fmt.Sprintf("%s-%s", tag, suffix)
	}

	return tag
}

func DockerBuildCommand(path string, args []string) (cmd exec.Cmd) {
	buildArgs := append([]string{"build"}, args...)
	cmd = *exec.Command("docker", buildArgs...)
	cmd.Env = append(os.Environ(), "DOCKER_BUILDKIT=1")
	cmd.Dir = path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}

func DockerRunCommand(args []string) (cmd exec.Cmd) {
	runArgs := append([]string{"run"}, args...)
	cmd = *exec.Command("docker", runArgs...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}
