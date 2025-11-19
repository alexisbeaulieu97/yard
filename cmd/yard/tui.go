package main

import (
	"fmt"

	"github.com/alexisbeaulieu97/yard/internal/tui"
	"github.com/alexisbeaulieu97/yard/internal/workspace"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Interactive ticket manager",
	RunE: func(cmd *cobra.Command, args []string) error {
		wsEngine := workspace.New(cfg.TicketsRoot)
		p := tea.NewProgram(tui.NewModel(wsEngine))
		if _, err := p.Run(); err != nil {
			return fmt.Errorf("failed to run tui: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
