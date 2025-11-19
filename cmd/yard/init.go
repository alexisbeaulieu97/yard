package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		configDir := filepath.Join(home, ".yard")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return err
		}

		configFile := filepath.Join(configDir, "config.yaml")
		if _, err := os.Stat(configFile); err == nil {
			fmt.Println("Config file already exists:", configFile)
			return nil
		}

		f, err := os.Create(configFile)
		if err != nil {
			return err
		}
		defer f.Close()

		// Write defaults
		_, err = f.WriteString(`projects_root: ~/.yard/projects
tickets_root: ~/.yard/tickets
ticket_naming: "{{.ID}}"
`)
		if err != nil {
			return err
		}

		fmt.Println("Initialized config at:", configFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
