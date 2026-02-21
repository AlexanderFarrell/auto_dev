package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	DatabaseTarget = Target{
		Build: BuildDatabase,
		Run:   nil,
	}
)

func BuildDatabase(item ConfigItem) error {
	if item.Output == nil || *item.Output == "" {
		return fmt.Errorf("database target requires output")
	}
	base := item.Path
	if base == "" {
		base = "."
	}

	files, err := databaseFiles(base, item)
	if err != nil {
		return err
	}

	var out strings.Builder
	for _, rel := range files {
		path := filepath.Join(base, rel)
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		out.Write(data)
		out.WriteString("\n\n")
	}

	if err := os.WriteFile(filepath.Clean(*item.Output), []byte(out.String()), 0o644); err != nil {
		return fmt.Errorf("write output: %w", err)
	}
	fmt.Printf("Wrote %s\n", *item.Output)
	return nil
}

func databaseFiles(base string, item ConfigItem) ([]string, error) {
	if item.OrderFile != nil && *item.OrderFile != "" {
		path := filepath.Join(base, *item.OrderFile)
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read order file: %w", err)
		}
		lines := strings.Split(string(data), "\n")
		out := make([]string, 0, len(lines))
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			out = append(out, line)
		}
		return out, nil
	}
	if len(item.Inputs) == 0 {
		return nil, fmt.Errorf("database target requires order_file or inputs")
	}
	return item.Inputs, nil
}
