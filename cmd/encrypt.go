package cmd

import (
	"fmt"
	"os"

	"github.com/gustavobertoi/hermes/files"
	"github.com/gustavobertoi/hermes/signatures"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Short:   "Encrypt a file using assimetric signatures",
	Aliases: []string{"e", "enc"},
	Run:     encryptHandler,
}

func encryptHandler(cmd *cobra.Command, args []string) {
	inputPath, err := cmd.Flags().GetString("input")
	if err != nil {
		cmd.Printf("Error reading input path: %s", err)
		os.Exit(1)
	}

	outputPath, err := cmd.Flags().GetString("output")
	if err != nil {
		cmd.Printf("Error reading output path: %s", err)
		os.Exit(1)
	}

	algorithm, err := cmd.Flags().GetString("algorithm")
	if err != nil {
		cmd.Printf("Error reading algorithm: %s", err)
		os.Exit(1)
	}

	targetFile := files.NewFile(inputPath)
	if err := targetFile.Load(); err != nil {
		cmd.Printf("Error loading file: %s", err)
		os.Exit(1)
	}

	signature := signatures.NewSignature(algorithm)
	if signature == nil {
		cmd.Print("Signature algorithm not supported")
		os.Exit(1)
	}

	if err := signature.Generate(); err != nil {
		cmd.Printf("Error generating signature: %s", err)
		os.Exit(1)
	}

	encryptedFile, err := signature.Encrypt(targetFile.Content())
	if err != nil {
		cmd.Printf("Error encrypting file: %s", err)
		os.Exit(1)
	}

	cmd.Printf("File %s (%s) is encrypted with algorithm %s\n", targetFile.Name(), targetFile.ID, algorithm)
	cmd.Println()
	cmd.Printf("Saving file in %s...", outputPath)
	cmd.Println()

	encryptedFileName := fmt.Sprintf("%s-%s.txt", targetFile.ID, algorithm)
	if err := encryptedFile.SaveContentToFile(outputPath, encryptedFileName); err != nil {
		cmd.Printf("Error saving file: %s", err)
		os.Exit(1)
	}

	encryptedHashSumName := fmt.Sprintf("%s-hash-sum.txt", targetFile.ID)
	if err := encryptedFile.SaveHashSum(outputPath, encryptedHashSumName); err != nil {
		cmd.Printf("Error saving hash sum: %s", err)
		os.Exit(1)
	}

	cmd.Printf("File saved in %s", outputPath)
}
