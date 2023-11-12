package resize

import (
	"videotool/internal/resize"
	"github.com/spf13/cobra"
)

var size string // Локальная переменная для набора разрешений

var ConvertCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize videos from one resolution to another",
	Run: func(cmd *cobra.Command, args []string) {
		resize.Handle(size)
	},
}

func init() {
	ConvertCmd.Flags().StringVarP(&size, "size", "s", "HD", "Video resolutions (LD, SD, HD, FHD, QHD, UHD4K, UHD8K, UW1080, UW1440, C4K)")
}
