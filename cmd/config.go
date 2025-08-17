package cmd

import (
	"devy/internal/config"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var envFlag string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config <environment-name>",
	Short: "Configure an existing environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		envName := args[0]

		if envFlag != "" {
			parts := strings.SplitN(envFlag, "=", 2)
			if len(parts) != 2 {
				return fmt.Errorf("Invalid env format. Use KEY=VALUE")
			}

			key := parts[0]
			value := parts[1]

			err := config.AddEnvVariable(envName, key, value)
			if err != nil {
				return fmt.Errorf("Unable to add environment variable: %v", err)
			}

			cmd.Printf("✅ Successfully added environment variable '%s=%s' to .env file in environment '%s'\n", key, value, envName)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVar(&envFlag, "env", "", "Environment variable to set in format KEY=VALUE")
}
