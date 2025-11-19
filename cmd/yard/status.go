package main

import (
	"fmt"

	"github.com/alexisbeaulieu97/yard/internal/workspace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status [TICKET-ID]",
	Short: "Show status of ticket workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		wsEngine := workspace.New(cfg.TicketsRoot)
		// gitEngine := gitx.New(cfg.ProjectsRoot) // Unused for now

		var tickets []string
		if len(args) > 0 {
			tickets = []string{args[0]}
		} else {
			// List all
			list, err := wsEngine.List()
			if err != nil {
				return err
			}
			for _, t := range list {
				tickets = append(tickets, t.ID)
			}
		}

		for _, tID := range tickets {
			fmt.Printf("Ticket: %s\n", tID)
			// We need to load the ticket to get repos
			// For now, just listing the directory contents as a proxy or assume we can load it.
            // The List() returns Ticket structs, but if we passed an ID arg we need to load it.
            // Let's just assume we iterate the repos found in the directory for now or load metadata.
            
            // Re-loading metadata logic (duplicated from List, should be in Workspace engine)
            // For MVP, let's just say "Status not fully implemented for individual repos yet" 
            // or try to scan the directory.
            
            // ticketDir := filepath.Join(cfg.TicketsRoot, tID)
            // ...
            
            fmt.Println("  (Status check implementation pending full metadata loading)")
            
            // Example check on a repo if we knew the name
            // isDirty, _, _ := gitEngine.Status(repoPath)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
