package main

import (
	"fmt"
	"runtime/debug"
)

var version = "dev"

func main() {

	fmt.Println("hello world", getVersion())

	//fmt.Println("hello world")
}
func getVersion() string {

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown version: error reading build info"
	}
	if info.Main.Version != "(devel)" {
		return info.Main.Version
	}

	return version
}
