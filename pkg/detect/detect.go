package detect

import (
	"log"
	"runtime"
)

func DetectOS() (res string) {

	os := runtime.GOOS

	switch os {
	case "windows":
		res = "windows"
	case "linux":
		res = "linux"
	case "darwin":
		res = "macos"
	default:
		res = ""
		log.Println("Unknown operation system")
	}

	return
}
