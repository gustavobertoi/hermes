package cmd

import (
	"fmt"
	"os"

	"github.com/gustavobertoi/hermes/config"
	"github.com/gustavobertoi/hermes/files"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Short:   "Encrypt a file using assimetric signatures",
	Aliases: []string{"e", "enc"},
	Run:     encryptHandler,
}

func encryptHandler(cmd *cobra.Command, args []string) {
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

	encryptedFile, err := signature.Encrypt(file.Content())
	if err != nil {
		cmd.Printf("Error encrypting file: %s", err)
		os.Exit(1)
	}

	cmd.Println()
	cmd.Printf("File %s (%s) is encrypted with algorithm %s\n", file.Name(), file.ID, algorithm)
	cmd.Println()

	encryptedFileName := fmt.Sprintf("%s-%s.txt", file.ID, algorithm)

	outputPath, err := c.GetFilesPathBySignature(algorithm)
	if err != nil {
		cmd.Printf("Error getting file path: %s", err)
		os.Exit(1)
	}

	if err := encryptedFile.Save(outputPath, encryptedFileName); err != nil {
		cmd.Printf("Error saving file: %s", err)
		os.Exit(1)
	}

	encryptedHashSumName := fmt.Sprintf("%s-hash-sum.txt", file.ID)
	if err := encryptedFile.SaveHashSum(outputPath, encryptedHashSumName); err != nil {
		cmd.Printf("Error saving hash sum: %s", err)
		os.Exit(1)
	}

	cmd.Printf("Files (%s, %s) saved in %s", encryptedFileName, encryptedHashSumName, outputPath)
}
