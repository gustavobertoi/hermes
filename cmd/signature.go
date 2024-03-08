package cmd

import (
	"os"

	"github.com/gustavobertoi/hermes/config"
	"github.com/gustavobertoi/hermes/signatures"
	"github.com/spf13/cobra"
)

var signatureCmd = &cobra.Command{
	Use:     "signature",
	Short:   "Manage our signature to encrypt and decrypt files",
	Aliases: []string{"sig"},
	Run:     signatureHandler,
}

func signatureHandler(cmd *cobra.Command, args []string) {
	algorithm, err := cmd.Flags().GetString("algorithm")
	if err != nil {
		cmd.Printf("Error reading algorithm flag: %s", err)
		os.Exit(1)
	}

	hermesFigure.Print()
	cmd.Println()

	if !signatures.IsValidAlgorithm(algorithm) {
		cmd.Printf("Invalid algorithm: %s", algorithm)
		os.Exit(1)
	}

	signature := signatures.NewSignature(algorithm)
	if signature == nil {
		cmd.Printf("Error creating signature with algorithm: %s", algorithm)
		os.Exit(1)
	}

	cmd.Println()
	cmd.Printf("Generating signature with algorithm: %s", algorithm)

	if err := signature.Generate(); err != nil {
		cmd.Printf("Error generating signature: %s", err)
		os.Exit(1)
	}

	cmd.Println()
	cmd.Print("Signature generated successfully, saving...")
	cmd.Println()

	c, err := config.GetConfig()
	if err != nil {
		cmd.Printf("Error reading config: %s", err)
		os.Exit(1)
	}

	if err := c.AddSignature(algorithm, signature); err != nil {
		cmd.Printf("Error adding signature: %s", err)
		os.Exit(1)
	}

	cmd.Print("Signature saved successfully")
}
