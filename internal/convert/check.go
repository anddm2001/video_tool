package convert

import (
	"log"
	"os/exec"
	"strings"
)

func checkCUDASupport() bool {
	cmd := exec.Command("/usr/bin/ffmpeg", "-hide_banner", "-hwaccels")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error checking CUDA support: %v", err)
		return false
	}

	return strings.Contains(string(output), "cuda")
}