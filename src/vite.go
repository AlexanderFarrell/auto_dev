package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	ViteTarget = Target{
		Build: BuildVite,
		Run:   RunVite,
	}
)

func BuildVite(item ConfigItem) error {
	if item.Path == "" {
		return fmt.Errorf("vite target requires path")
	}
	if err := runInDir(item.Path, "npm", "install"); err != nil {
		return err
	}
	return runInDir(item.Path, "npm", "run", "build")
}

func RunVite(item ConfigItem) error {
	if item.Path == "" {
		return fmt.Errorf("vite target requires path")
	}
	if err := runInDir(item.Path, "npm", "install"); err != nil {
		return err
	}
	return runInDir(item.Path, "npm", "run", "dev")
}

func runInDir(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
