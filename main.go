package main

import (
	"fmt"
	"runtime"
)

const version = "0.0.1"

func main() {
	// OS判定(mac, windows, linux)
	if isMac() {
		fmt.Println("Darwin")
	} else if isWindows() {
		fmt.Println("Windows")
	} else if isLinux() {
		fmt.Println("Linux")
	} else {
		fmt.Println("Unknown")
	}
}

func isMac() bool {
	return runtime.GOOS == "darwin"
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func isLinux() bool {
	return runtime.GOOS == "linux"
}
