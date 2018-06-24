package main

import (
	"fmt"
	"os/exec"
)

func main() {
	res, err := exec.Command("git", "log", "--format=%h").Output()
	fmt.Printf("%#v | %#v\n", res, err.Error())
}
