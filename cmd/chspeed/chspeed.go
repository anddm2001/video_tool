package chspeed

import (
	"videotool/internal/chspeed"
	"github.com/spf13/cobra"
)

var speed string // Локальная переменная для набора разрешений

var ConvertCmd = &cobra.Command{
	Use:   "chspeed",
	Short: "Change speed for video",
	Run: func(cmd *cobra.Command, args []string) {
		chspeed.Handle(speed)
	},
}

func init() {
	ConvertCmd.Flags().StringVarP(&speed, "speed", "sp", "2.0", "Video speed value")
}
