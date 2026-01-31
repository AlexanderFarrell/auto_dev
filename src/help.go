package main

import (
	"fmt"
)

func Help(cfg Config) {
	var output string = "Usage: dev [help|"
	for key, _ := range cfg {
		output += key + "|"
	}
	output = output[0:len(output)-1]
	output += "]\n"
	fmt.Printf(output)
}

func HelpTarget(target Target) {
	var output string = "["
	if target.Build != nil {
		output += "build|"
	}
	if target.Run != nil {
		output += "run|"
	}

	output = output[0:len(output)-1]
	output += "]"
	fmt.Printf(output)
}