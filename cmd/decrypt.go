package cmd

import (
	"os"

	"github.com/gustavobertoi/hermes/config"
	"github.com/gustavobertoi/hermes/files"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "Decrypt a file",
	Long:    `Decrypt a file using a personal signature.`,
	Aliases: []string{"d", "dec"},
	Run:     decryptHandler,
}

func decryptHandler(cmd *cobra.Command, args []string) {
	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		cmd.Printf("Error reading file path: %s", err)
		os.Exit(1)
	}

	algorithm, err := cmd.Flags().GetString("algorithm")
	if err != nil {
		cmd.Printf("Error reading algorithm: %s", err)
		os.Exit(1)
	}

	file := files.NewFile(filePath)
	if err := file.Load(); err != nil {
		cmd.Printf("Error loading file: %s", err)
		os.Exit(1)
	}

	c, err := config.GetConfig()
	if err != nil {
		cmd.Printf("Error getting config: %s", err)
		os.Exit(1)
	}

	signature, err := c.GetSignature(algorithm)
	if err != nil {
		cmd.Printf("Error getting signature: %s", err)
		os.Exit(1)
	}

	cmd.Printf("Decrypting file %s (%s) with algorithm %s\n", file.Name(), file.ID, algorithm)

	content := file.Content()
	decryptedFile, err := signature.Decrypt(content)
	if err != nil {
		cmd.Printf("Error decrypting file: %s", err)
		os.Exit(1)
	}

	cmd.Println()
	cmd.Printf("File %s (%s) is decrypted with algorithm %s\n", file.Name(), file.ID, algorithm)
	cmd.Println()

	cmd.Printf("Decrypted content: %s\n", string(decryptedFile))

}
