package cmd

import (
	"fmt"
	"os"

	"github.com/rahulp0817/konflint/internal/detector"
	"github.com/spf13/cobra"
)

var (
	autofix bool
	verbose bool
)

var validateCmd = &cobra.Command{
	Use:   "validate [file]",
	Short: "Validate a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		// Read file
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("❌ Cannot read file: %s\n", filePath)
			os.Exit(1)
		}

		// detect the file type
		fileType := detector.Detect(filePath, data)

		if fileType == detector.Unknown {
			fmt.Printf("\n  ❌ Unknown file type: %s\n", filePath)
			fmt.Printf("  Konflint supports: Kubernetes, Docker Compose, GitHub Actions\n\n")
			os.Exit(1)
		}

		// Print detected type
		fmt.Printf("✅ Detected: %s\n", fileType)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
