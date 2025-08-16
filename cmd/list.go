package cmd

import (
	"devy/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		environments, err := config.ListEnvironments()
		if err != nil {
			return fmt.Errorf("Unable to list environments: %v", err)
		}

		for _, env := range environments {
			fmt.Println(env)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
