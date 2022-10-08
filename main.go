package main

import (
	"fmt"
	"runtime/debug"
	//	_ "github.com/caarlos0/svu"
)

var version = "dev"

func main() {

	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		fmt.Println(info)
	}

	fmt.Println("hello world")
}
