package main

import (
	"fmt"
	"os"
)

type Target struct {
	Build func(ConfigItem) error
	Run func(ConfigItem) error
}

func CallTarget(command string, name string, item ConfigItem, target Target) {
	switch command {
	case "build":
		if target.Build != nil {
			target.Build(item)
		} else {
			fmt.Printf("Target cannot build")
		}
	case "run":
		if target.Run != nil {
			target.Run(item)
		} else {
			fmt.Printf("Target cannot run")
		}
	default:
		fmt.Print("Unknown command: " + command)
		fmt.Print("Usage: dev " + name + " ")
		HelpTarget(target)
		fmt.Printf("\n")
		os.Exit(1)
	}
}