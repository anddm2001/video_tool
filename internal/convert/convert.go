package convert

import (
	"videotool/pkg/config"
	"videotool/pkg/detect"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func convert(in, out, codec, ffmpeg_path string) {

	os_name := detect.DetectOS()

	if len(os_name) == 0 {
		os_name = "linux"
	}

	if os_name == "windows" {
		command := fmt.Sprintf("%s -hide_banner -y -i %s -vcodec %s -crf 30 %s", ffmpeg_path, in, codec, out)
		cmd := exec.Command("cmd", "/C", command)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err := cmd.Run()

		if err != nil {
			log.Panicln(err.Error())
			log.Printf("Error converting file %s using codec %s: %v", in, codec, err)
		}
	} else {
		cmd := exec.Command(ffmpeg_path, "-hide_banner", "-y", "-i", in, "-vcodec", codec, "-crf", "30", out)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err := cmd.Run()

		if err != nil {
			log.Panicln(err.Error())
			log.Printf("Error converting file %s using codec %s: %v", in, codec, err)
		}
	}
}

func convertCUDA(in, out, codec, ffmpeg_path string) {

	cmd := exec.Command(ffmpeg_path, "-hide_banner", "-hwaccel", "cuda", "-y", "-i", in, "-vcodec", codec, "-crf", "30", out)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error converting file %s using codec %s: %v", in, codec, err)
	}
}

func Handle(codecKey string) {
	cfg := config.LoadConfig()

	files, err := os.ReadDir(cfg.InDir)
	if err != nil {
		log.Fatal(err)
	}

	codecMap := codecMap()

	codec, ok := codecMap[codecKey]
	if !ok {
		log.Printf("Unsupported codec: %s", codecKey)
		return
	}

	cudaSupported := checkCUDASupport()

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		inputFile := cfg.InDir + "/" + file.Name()
		outputFile := cfg.OutDir + "/converted_" + file.Name()

		if cudaSupported {
			convertCUDA(inputFile, outputFile, codec, cfg.FFMPEGPath)
		} else {
			convert(inputFile, outputFile, codec, cfg.FFMPEGPath)
		}
	}
}

func codecMap() map[string]string {
	cudaSupported := checkCUDASupport()

	// Определяем доступные кодеки в зависимости от поддержки CUDA
	var codecMap map[string]string
	if cudaSupported {
		codecMap = map[string]string{
			"h264": "h264_nvenc", // CUDA-supported codec for H.264
			"h265": "hevc_nvenc", // CUDA-supported codec for H.265
		}
	} else {
		codecMap = map[string]string{
			"h264": "libx264",
			"h265": "libx265",
			"vp9":  "libvpx-vp9",
			"av1":  "libaom-av1",
		}
	}

	return codecMap
}
