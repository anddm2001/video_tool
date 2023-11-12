package convert

import (
	"videotool/internal/convert"
	"github.com/spf13/cobra"
)

var codec string // Локальная переменная для кодека

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert videos from one format to another",
	Run: func(cmd *cobra.Command, args []string) {
		convert.Handle(codec)
	},
}

func init() {
	ConvertCmd.Flags().StringVarP(&codec, "codec", "c", "h264", "Codec to use for conversion (h264, h265, vp9, av1)")
}

