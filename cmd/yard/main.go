package main

import (
	"fmt"
	"os"

	"github.com/alexisbeaulieu97/yard/internal/config"
	"github.com/alexisbeaulieu97/yard/internal/logging"
	"github.com/spf13/cobra"
)

var (
	cfg    *config.Config
	logger *logging.Logger
)

var rootCmd = &cobra.Command{
	Use:   "yard",
	Short: "Ticket-centric workspaces",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.Load()
		if err != nil {
			return err
		}
		logger = logging.New()
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
