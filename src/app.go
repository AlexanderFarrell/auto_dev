package main

var Targets = map[string]Target{
	"go":       GoTarget,
	"vite":     ViteTarget,
	"database": DatabaseTarget,
	"docker":   DockerTarget,
}
