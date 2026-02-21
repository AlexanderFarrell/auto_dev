package main

import (
	"fmt"
	"os"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	// Look for an arg
	if len(os.Args) < 2 {
		Help(cfg)
		os.Exit(1)
	}

	arg1 := os.Args[1]
	if arg1 == "help" {
		Help(cfg)
		os.Exit(0)
	}

	commands := AllCommands(cfg)
	command := ""
	targetName := ""

	if isCommand(arg1, commands) {
		command = arg1
		if len(os.Args) >= 3 {
			targetName = os.Args[2]
		} else {
			targetName = "all"
		}
	} else {
		if len(os.Args) < 3 {
			Help(cfg)
			os.Exit(1)
		}
		targetName = arg1
		command = os.Args[2]
	}

	if err := RunCommand(cfg, command, targetName); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		Help(cfg)
		os.Exit(1)
	}
}

func isCommand(value string, commands []string) bool {
	for _, cmd := range commands {
		if value == cmd {
			return true
		}
	}
	return false
}

func RunCommand(cfg ConfigFile, command string, targetName string) error {
	targets, err := ResolveTargets(cfg, targetName)
	if err != nil {
		return err
	}
	for _, name := range targets {
		item := cfg.Targets[name]
		if err := ExecuteCommand(command, name, item); err != nil {
			return err
		}
	}
	return nil
}
