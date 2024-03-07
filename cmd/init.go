package cmd

import (
	"os"
	"os/user"
	"path"

	"github.com/gustavobertoi/hermes/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hermes configuration",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	hermesFigure.Print()
	cmd.Println()

	usr, err := user.Current()
	if err != nil {
		cmd.Printf("Error reading user home: %s", err)
		os.Exit(1)
	}

	folderPath := path.Join(usr.HomeDir, ".hermes")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			cmd.Printf("Error creating folder path: %s", err)
			os.Exit(1)
		}
	}

	cmd.Println()
	cmd.Printf("Initializing hermes at %s", folderPath)
	cmd.Println()

	config := config.NewConfig(folderPath)
	if err := config.SaveConfig(); err != nil {
		cmd.Printf("Error saving config: %s", err)
		os.Exit(1)
	}

	cmd.Printf("Config saved at %s", folderPath)
	cmd.Println()
	cmd.Printf("Hermes initialized successfully!")
}
