package resize

import (
	"videotool/pkg/config"
	"videotool/pkg/detect"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func resize(in, out, resolution, ffmpeg_path string) {

	os_name := detect.DetectOS()

	if len(os_name) == 0 {
		os_name = "linux"
	}

	if os_name == "windows" {
		command := fmt.Sprintf("%s -hide_banner -y -i %s -s %s %s", ffmpeg_path, in, resolution, out)
		cmd := exec.Command("cmd", "/C", command)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			log.Panicln(err.Error())
			log.Printf("Error converting file %s using video resolution %s: %v", in, resolution, err)
		}
	} else {
		cmd := exec.Command(ffmpeg_path, "-hide_banner", "-y", "-i", in, "-s", resolution, out)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			log.Panicln(err.Error())
			log.Printf("Error converting file %s using video resolution %s: %v", in, resolution, err)
		}
	}
}

func Handle(size string) {
	cfg := config.LoadConfig()

	files, err := os.ReadDir(cfg.InDir)
	if err != nil {
		log.Fatal(err)
	}

	resolutionMap := resolutionMap()

	sizeUpper := strings.ToUpper(size)

	res, ok := resolutionMap[sizeUpper]
	if !ok {
		log.Printf("Unsupported resolution: %s", size)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		inputFile := cfg.InDir + "/" + file.Name()
		outputFile := cfg.OutDir + "/converted_" + file.Name()

		resize(inputFile, outputFile, res, cfg.FFMPEGPath)
	}
}

func resolutionMap() map[string]string {

	return map[string]string{
		"LD":    "640x360",    // Low Definition
		"SD":    "720x480",    // Standard Definition
		"HD":    "1280x720",   // High Definition
		"FHD":   "1920x1080",  // Full HD
		"QHD":   "2560x1440",  // Quad HD
		"UHD4K": "3840x2160",  // 4K Ultra HD
		"UHD8K": "7680x4320",  // 8K Ultra HD
		"UW1080": "2560x1080", // Ultra Wide 1080p
		"UW1440": "3440x1440", // Ultra Wide 1440p
		"C4K":   "4096x2160",  // Cinematic 4K
	}
}
