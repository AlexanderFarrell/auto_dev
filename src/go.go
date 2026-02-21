package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	GoTarget = Target{
		Build: BuildGo,
		Run:   RunGo,
	}
)

func BuildGo(item ConfigItem) error {
	output, err := resolveOutput(item)
	if err != nil {
		return err
	}
	args := []string{"build", "-o", output, "."}
	fmt.Printf("Building go app at %s\n", item.Path)
	cmd := exec.Command("go", args...)
	if item.Path != "" {
		cmd.Dir = item.Path
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Printf("Build complete\n")
	return nil
}

func RunGo(item ConfigItem) error {
	output, err := resolveOutput(item)
	if err != nil {
		return err
	}
	runPath := output
	cmdName := output
	var cmdArgs []string
	scriptDir := ""
	if item.Script != nil && *item.Script != "" {
		scriptPath, err := resolveScript(item)
		if err != nil {
			return err
		}
		info, err := os.Stat(scriptPath)
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("script must be a file: %s", scriptPath)
		}
		runPath = scriptPath
		if info.Mode()&0111 == 0 {
			cmdName = "sh"
			cmdArgs = []string{scriptPath}
		} else {
			cmdName = scriptPath
		}
		scriptDir = filepath.Dir(scriptPath)
	}
	fmt.Printf("Running %s\n", runPath)
	cmd := exec.Command(cmdName, cmdArgs...)
	cwd, err := os.Getwd()
	if err == nil {
		cmd.Env = append(os.Environ(), "AUTO_DEV_CWD="+cwd)
	}
	if scriptDir != "" {
		cmd.Dir = scriptDir
	} else if item.Path != "" {
		cmd.Dir = item.Path
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func resolveOutput(item ConfigItem) (string, error) {
	if item.Output == nil || *item.Output == "" {
		return "", fmt.Errorf("go target requires output")
	}
	output := *item.Output
	if filepath.IsAbs(output) {
		return output, nil
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Clean(filepath.Join(cwd, output)), nil
}

func resolveScript(item ConfigItem) (string, error) {
	if item.Script == nil || *item.Script == "" {
		return "", fmt.Errorf("go target requires script")
	}
	script := *item.Script
	if filepath.IsAbs(script) {
		return script, nil
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Clean(filepath.Join(cwd, script)), nil
}
