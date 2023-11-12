package chspeed

import (
	"videotool/pkg/config"
	"videotool/pkg/detect"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func chspeed(in, out, speed, ffmpeg_path string) {

	os_name := detect.DetectOS()

	if len(os_name) == 0 {
		os_name = "linux"
	}

	if os_name == "windows" {
		command := fmt.Sprintf("%s -hide_banner -y -i %s --filter:v %s %s", ffmpeg_path, in, speed, out)
		cmd := exec.Command("cmd", "/C", command)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			log.Panicln(err.Error())
			log.Printf("Error converting file %s using speed %s: %v", in, speed, err)
		}
	} else {
		cmd := exec.Command(ffmpeg_path, "-hide_banner", "-y", "-i", in, "-filter:v", speed, out)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			log.Panicln(err.Error())
			log.Printf("Error converting file %s using speed %s: %v", in, speed, err)
		}
	}
}

func Handle(speed string) {
	cfg := config.LoadConfig()

	files, err := os.ReadDir(cfg.InDir)
	if err != nil {
		log.Fatal(err)
	}

	fullSpeedValueStr := "setpts=" + speed + "*PTS"

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		inputFile := cfg.InDir + "/" + file.Name()
		outputFile := cfg.OutDir + "/converted_" + file.Name()

		chspeed(inputFile, outputFile, fullSpeedValueStr, cfg.FFMPEGPath)
	}
}
