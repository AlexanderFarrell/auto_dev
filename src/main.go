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

	target := os.Args[1]
	if target == "help" {
		Help(cfg)
		os.Exit(0)
	}
	if len(os.Args) < 3 {
		Help(cfg)
		os.Exit(1)
	}
	command := os.Args[2]
	for key, item := range cfg {
		if key == target {
			for keyTarget, target
			CallTarget(command, target, )
			// fmt.Printf("TODO add this")
			os.Exit(0)
		}
	}

	fmt.Printf("Error: Not a valid command: %s\n", target)
	Help(cfg)
	os.Exit(1)
}