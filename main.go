package main

import "github.com/SecurityForEveryone/software-composition-scanner/cmd"

var version string

func main() {
	version = "0.1"
	cmd.Execute()
}
