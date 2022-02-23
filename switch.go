package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("OS => ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("%s.\n", os)
	}
}
