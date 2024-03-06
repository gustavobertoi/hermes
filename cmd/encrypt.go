package cmd

import (
	"os"
	"path"

	"github.com/gustavobertoi/hermes/internal/files"
	"github.com/gustavobertoi/hermes/internal/signatures"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Short:   "Encrypt a file using assimetric signatures",
	Aliases: []string{"e", "enc"},
	Run:     encrypt,
}

func initEncrypt() {
	encryptCmd.Flags().StringP("input", "i", "", "Input file path")
	encryptCmd.Flags().StringP("output", "o", "", "Output file path")
	encryptCmd.Flags().StringP("algorithm", "a", signatures.RSA, "Algorithm to use for encryption")
	encryptCmd.MarkFlagRequired("input")
	encryptCmd.MarkFlagRequired("output")
	encryptCmd.MarkFlagRequired("algorithm")
	rootCmd.AddCommand(encryptCmd)
}

func encrypt(cmd *cobra.Command, args []string) {
	inputPath, err := cmd.Flags().GetString("input")
	if err != nil {
		cmd.Printf("Error reading input path: %s", err)
		os.Exit(1)
		return
	}

	outputPath, err := cmd.Flags().GetString("output")
	if err != nil {
		cmd.Printf("Error reading output path: %s", err)
		os.Exit(1)
		return
	}

	algorithm, err := cmd.Flags().GetString("algorithm")
	if err != nil {
		cmd.Printf("Error reading algorithm: %s", err)
		os.Exit(1)
		return
	}

	targetFile := files.NewFile(inputPath)
	if err := targetFile.Load(); err != nil {
		cmd.Printf("Error loading file: %s", err)
		os.Exit(1)
		return
	}

	signature := signatures.NewSignature(algorithm)
	if signature == nil {
		cmd.Print("Signature algorithm not supported")
		os.Exit(1)
		return
	}

	if err := signature.Generate(); err != nil {
		cmd.Printf("Error generating signature: %s", err)
		os.Exit(1)
		return
	}

	rawEncrypt, err := signature.Encrypt(targetFile.Content())
	if err != nil {
		cmd.Printf("Error encrypting file: %s", err)
		os.Exit(1)
		return
	}

	cmd.Printf("File %s (%s) is encrypted with algorithm %s\n", targetFile.Name(), targetFile.ID, algorithm)
	cmd.Println()
	cmd.Printf("Saving file in %s...", outputPath)
	cmd.Println()

	encryptedFileFullPath := path.Join(outputPath, targetFile.ID+".pem")
	encryptedFile := files.NewFile(encryptedFileFullPath)
	encryptedFile.SetContent(rawEncrypt)

	if err := encryptedFile.Save(); err != nil {
		cmd.Printf("Error saving file: %s", err)
		os.Exit(1)
		return
	}

	cmd.Printf("File saved in %s", encryptedFile.Path())
}
