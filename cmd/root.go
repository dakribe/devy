package cmd

import (
	"devy/internal/config"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devy",
	Short: "A development environment management CLI",
	Long: `Devy is a command-line tool for managing development environments.
It allows you to create, list, and manage development environments
for your projects, helping you organize and switch between different
development setups efficiently.

Use 'devy create' to create new environments and 'devy list' to 
view all existing environments.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.CreateConfig()
}
