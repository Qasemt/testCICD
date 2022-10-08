package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		fmt.Println(info)
	}

	fmt.Println("hello world")
}
