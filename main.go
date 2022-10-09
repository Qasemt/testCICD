package main

import (
	"fmt"
	"runtime/debug"
)

var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

func main() {

	fmt.Println("hello world", buildVersion(version, commit, date, builtBy))

	//fmt.Println("hello world")
}

func buildVersion(version, commit, date, builtBy string) string {
	result := "nfpm version " + version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nmodule version: %s, checksum: %s", result, version, info.Main.Sum)
	}
	return result
}
