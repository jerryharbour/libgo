package version

import "fmt"

var (
	Version = "1.0.0"

	Tag = "2023-02-09_Release"

	BuildTime = "2023-02-09 14:50:00"

	GitHash = "unknown"
)

func ShowVersion() {
	fmt.Printf("%s", GetVersion())
}

func GetVersion() string {
	return fmt.Sprintf("Version: %s\nTag: %s\nBuildTime: %s\nGitHash: %s\n", Version,
		Tag, BuildTime, GitHash)
}
