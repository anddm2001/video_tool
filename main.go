package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"videotool/cmd/convert"
	converter "videotool/internal/convert"
)

var codec string // переменная для хранения выбранного кодека

func init() {
	rootCmd.PersistentFlags().StringVarP(&codec, "codec", "c", "h264", "Codec to use for conversion (h264, h265, vp9, av1)")
	rootCmd.AddCommand(convert.ConvertCmd)
}

var rootCmd = &cobra.Command{
	Use:   "converter",
	Short: "Video converter is a tool to convert videos",
	Run: func(cmd *cobra.Command, args []string) {
		// Вызов функции конвертации по умолчанию если нет других команд
		converter.Handle(codec)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}


