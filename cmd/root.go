package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var success = lipgloss.NewStyle().
	Foreground(lipgloss.Color("42")).
	Bold(false)

var Version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "konflint",
	Short: "konflint — Smart YAML config linter for terminal developers",
	Long: `konflint detects what kind of config file you're writing and gives you human-readable errors, plain-English explanations, and automatic fixes right in your terminal.

	konflint validate deployment.yaml
	konflint validate deployment.yaml --fix
	konflint validate deployment.yaml --fix --v`,

	Version: Version,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(success.Render("Konflint Installed Successfully! 🎉"))
		fmt.Println()
		fmt.Println("Run 'konflint --version' to see the version.")
		fmt.Println("Run 'konflint validate deployment.yaml' to lint a file.")
		fmt.Println("Run 'konflint validate deployment.yaml --fix' to fix issues automatically.")
		fmt.Println()
		fmt.Println("Run 'konflint --help' to see available commands.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetVersionTemplate("konflint {{.Version}}\n")
}
