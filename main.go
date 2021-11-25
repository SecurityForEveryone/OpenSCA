package main

import (
	"fmt"
	"github.com/SecurityForEveryone/software-composition-scanner/cmd"
)

var version string

func main() {
	fmt.Println("Version from go releaser")
	fmt.Println(version)
	cmd.Execute()
}
