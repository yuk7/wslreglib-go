package wslreg

import "os"

func getWindowsDirectory() string {
	dir := os.Getenv("SYSTEMROOT")
	if dir != "" {
		return dir
	}
	dir = os.Getenv("WINDIR")
	if dir != "" {
		return dir
	}
	return "C:\\WINDOWS"

}
