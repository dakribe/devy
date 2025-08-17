package cmd

import (
	"devy/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <name> <path>",
	Short: "Creates a new environment",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		env := &config.EnvironmentConfig{
			Name:       args[0],
			ProjectDir: args[1],
		}
		err := config.CreateEnvironment(env)
		if err != nil {
			return fmt.Errorf("Unable to create environment: %v", err)
		}

		cmd.Printf("✅ Successfully created environment '%s' at '%s'\n", args[0], args[1])

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
