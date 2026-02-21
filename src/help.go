package main

import (
	"fmt"
	"sort"
)

func Help(cfg ConfigFile) {
	commands := AllCommands(cfg)
	if len(commands) == 0 {
		commands = []string{"build", "run"}
	}
	fmt.Printf("Usage: dev <command> [target|group|all]\n")
	fmt.Printf("       dev <target> <command> (legacy)\n")
	fmt.Printf("Commands: %s\n", joinPipe(commands))
	fmt.Printf("Targets: %s\n", joinPipe(sortedKeys(cfg.Targets)))
	if len(cfg.Groups) > 0 {
		fmt.Printf("Groups: %s\n", joinPipe(sortedKeys(cfg.Groups)))
	}
}

func joinPipe(items []string) string {
	if len(items) == 0 {
		return "-"
	}
	out := ""
	for _, item := range items {
		out += item + "|"
	}
	return out[:len(out)-1]
}

func sortedKeys[T any](items map[string]T) []string {
	out := make([]string, 0, len(items))
	for key := range items {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}
