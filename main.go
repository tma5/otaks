package main

import (
	"runtime"

	"github.com/tma5/otaks/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}