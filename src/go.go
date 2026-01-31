package main

var (
	GoTarget = Target{
		Build: BuildGo,
		Run: RunGo,
	}
)

func BuildGo(item ConfigItem) error {
	return nil
}

func RunGo(item ConfigItem) error {
	return nil
}