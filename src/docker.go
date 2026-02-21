package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	DockerTarget = Target{
		Commands: map[string]func(ConfigItem) error{
			"build":    BuildDocker,
			"deploy":   DeployDocker,
			"redeploy": RedeployDocker,
		},
	}
)

func BuildDocker(item ConfigItem) error {
	compose, err := dockerComposeFile(item)
	if err != nil {
		return err
	}
	return runDocker(item.Path, "docker", "compose", "-f", compose, "build")
}

func DeployDocker(item ConfigItem) error {
	compose, err := dockerComposeFile(item)
	if err != nil {
		return err
	}
	stack, err := dockerStack(item)
	if err != nil {
		return err
	}
	return runDocker(item.Path, "docker", "stack", "deploy", "-c", compose, stack)
}

func RedeployDocker(item ConfigItem) error {
	compose, err := dockerComposeFile(item)
	if err != nil {
		return err
	}
	if item.Service != nil && *item.Service != "" {
		if err := runDocker(item.Path, "docker", "compose", "-f", compose, "build"); err != nil {
			return err
		}
		return runDocker(item.Path, "docker", "service", "update", "--force", *item.Service)
	}
	stack, err := dockerStack(item)
	if err != nil {
		return err
	}
	if err := runDocker(item.Path, "docker", "stack", "rm", stack); err != nil {
		return err
	}
	if err := runDocker(item.Path, "docker", "compose", "-f", compose, "build"); err != nil {
		return err
	}
	return runDocker(item.Path, "docker", "stack", "deploy", "-c", compose, stack)
}

func dockerComposeFile(item ConfigItem) (string, error) {
	if item.Compose == nil || *item.Compose == "" {
		return "", fmt.Errorf("docker target requires compose")
	}
	return *item.Compose, nil
}

func dockerStack(item ConfigItem) (string, error) {
	if item.Stack == nil || *item.Stack == "" {
		return "", fmt.Errorf("docker target requires stack")
	}
	return *item.Stack, nil
}

func runDocker(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
