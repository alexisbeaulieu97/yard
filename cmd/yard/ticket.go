package main

import (
	"fmt"
	"strings"

	"github.com/alexisbeaulieu97/yard/internal/domain"
	"github.com/alexisbeaulieu97/yard/internal/gitx"
	"github.com/alexisbeaulieu97/yard/internal/workspace"
	"github.com/spf13/cobra"
)

var ticketCmd = &cobra.Command{
	Use:   "ticket",
	Short: "Manage ticket workspaces",
}

var ticketNewCmd = &cobra.Command{
	Use:   "new <TICKET-ID>",
	Short: "Create a new ticket workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ticketID := args[0]
		reposStr, _ := cmd.Flags().GetString("repos")
		repoNames := strings.Split(reposStr, ",")

		// Initialize engines
		wsEngine := workspace.New(cfg.TicketsRoot)
		gitEngine := gitx.New(cfg.ProjectsRoot)

		var repos []domain.Repo
		for _, val := range repoNames {
			if val == "" {
				continue
			}
            
            name := val
            url := "https://github.com/example/" + name
            
            // If input is like "owner/repo", use that for URL and repo name
            if strings.Contains(val, "/") {
                parts := strings.Split(val, "/")
                if len(parts) == 2 {
                    name = parts[1]
                    url = "https://github.com/" + val
                }
            }

			repos = append(repos, domain.Repo{Name: name, URL: url})
		}

		ticket := &domain.Ticket{
			ID:    ticketID,
			Repos: repos,
		}

		// Create workspace
		if err := wsEngine.Create(ticket); err != nil {
			return err
		}
		fmt.Printf("Created workspace for %s\n", ticketID)

		// Setup worktrees
		for _, repo := range repos {
			fmt.Printf("Setting up %s...\n", repo.Name)
            // Ensure canonical exists (mocked URL for now)
            _, err := gitEngine.EnsureCanonical(repo.URL, repo.Name)
            if err != nil {
                fmt.Printf("Warning: failed to ensure canonical for %s: %v\n", repo.Name, err)
                continue
            }
            
            // Create worktree
            worktreePath := fmt.Sprintf("%s/%s/%s", cfg.TicketsRoot, ticketID, repo.Name)
            if err := gitEngine.CreateWorktree(repo.Name, worktreePath, ticketID); err != nil {
                 fmt.Printf("Warning: failed to create worktree for %s: %v\n", repo.Name, err)
            }
		}

		return nil
	},
}

var ticketListCmd = &cobra.Command{
	Use:   "list",
	Short: "List active tickets",
	RunE: func(cmd *cobra.Command, args []string) error {
		wsEngine := workspace.New(cfg.TicketsRoot)
		tickets, err := wsEngine.List()
		if err != nil {
			return err
		}

		for _, t := range tickets {
			fmt.Printf("%s (%d repos)\n", t.ID, len(t.Repos))
		}
		return nil
	},
}

var ticketCloseCmd = &cobra.Command{
	Use:   "close <TICKET-ID>",
	Short: "Close a ticket workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ticketID := args[0]
		wsEngine := workspace.New(cfg.TicketsRoot)
        
        // TODO: Check git status before deleting (Safety)
        
		if err := wsEngine.Delete(ticketID); err != nil {
			return err
		}
		fmt.Printf("Closed ticket %s\n", ticketID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ticketCmd)
	ticketCmd.AddCommand(ticketNewCmd)
	ticketCmd.AddCommand(ticketListCmd)
	ticketCmd.AddCommand(ticketCloseCmd)

	ticketNewCmd.Flags().String("repos", "", "Comma-separated list of repositories")
}
