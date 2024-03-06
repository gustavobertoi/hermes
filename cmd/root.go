package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var hermesFigure = figure.NewColorFigure("hermes", "colossal", "white", true)
var hermesBinaryFigure = figure.NewColorFigure("hermes", "binary", "yellow", true)

var rootCmd = &cobra.Command{
	Use:     "hermes",
	Short:   "A way to control files, encrypt them and sync with external sources",
	Version: "1.0.0-beta",
	Run: func(cmd *cobra.Command, args []string) {
		hermesBinaryFigure.Print()
		cmd.Println()
		hermesFigure.Print()
		cmd.Println()
		hermesBinaryFigure.Print()
		cmd.Println()
		cmd.Printf("Version: %s", cmd.Version)
	},
}

func Execute() {
	initEncrypt()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
