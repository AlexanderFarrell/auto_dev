package main

import (
	"fmt"
	"sort"
)

type Target struct {
	Build    func(ConfigItem) error
	Run      func(ConfigItem) error
	Commands map[string]func(ConfigItem) error
}

func AllCommands(cfg ConfigFile) []string {
	set := map[string]struct{}{}
	for _, item := range cfg.Targets {
		target, ok := Targets[item.Type]
		if !ok {
			continue
		}
		if target.Build != nil {
			set["build"] = struct{}{}
		}
		if target.Run != nil {
			set["run"] = struct{}{}
		}
		for key := range target.Commands {
			set[key] = struct{}{}
		}
	}
	out := make([]string, 0, len(set))
	for key := range set {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}

func ResolveTargets(cfg ConfigFile, name string) ([]string, error) {
	if name == "all" {
		out := make([]string, 0, len(cfg.Targets))
		for key := range cfg.Targets {
			out = append(out, key)
		}
		sort.Strings(out)
		return out, nil
	}
	if cfg.Groups != nil {
		if group, ok := cfg.Groups[name]; ok {
			return group, nil
		}
	}
	if _, ok := cfg.Targets[name]; ok {
		return []string{name}, nil
	}
	return nil, fmt.Errorf("unknown target or group: %s", name)
}

func ExecuteCommand(command string, name string, item ConfigItem) error {
	target, ok := Targets[item.Type]
	if !ok {
		return fmt.Errorf("target not supported: %s", item.Type)
	}
	if target.Commands != nil {
		if cmd, ok := target.Commands[command]; ok {
			return cmd(item)
		}
	}
	switch command {
	case "build":
		if target.Build == nil {
			return fmt.Errorf("target cannot build: %s", name)
		}
		return target.Build(item)
	case "run":
		if target.Run == nil {
			return fmt.Errorf("target cannot run: %s", name)
		}
		return target.Run(item)
	default:
		return fmt.Errorf("command not supported: %s", command)
	}
}
