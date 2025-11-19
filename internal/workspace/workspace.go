package workspace

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexisbeaulieu97/yard/internal/domain"
	"gopkg.in/yaml.v3"
)

// Engine manages ticket workspaces
type Engine struct {
	TicketsRoot string
}

// New creates a new Workspace Engine
func New(ticketsRoot string) *Engine {
	return &Engine{TicketsRoot: ticketsRoot}
}

// Create creates a new ticket workspace
func (e *Engine) Create(ticket *domain.Ticket) error {
	path := filepath.Join(e.TicketsRoot, ticket.ID)
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create ticket dir: %w", err)
	}

	// Save metadata
	metaPath := filepath.Join(path, "ticket.yaml")
	f, err := os.Create(metaPath)
	if err != nil {
		return fmt.Errorf("failed to create metadata file: %w", err)
	}
	defer f.Close()

	enc := yaml.NewEncoder(f)
	if err := enc.Encode(ticket); err != nil {
		return fmt.Errorf("failed to encode metadata: %w", err)
	}

	return nil
}

// List returns all active tickets
func (e *Engine) List() ([]domain.Ticket, error) {
	entries, err := os.ReadDir(e.TicketsRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read tickets root: %w", err)
	}

	var tickets []domain.Ticket
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		metaPath := filepath.Join(e.TicketsRoot, entry.Name(), "ticket.yaml")
		f, err := os.Open(metaPath)
		if err != nil {
			// Skip if no metadata (might be a random dir)
			continue
		}
		defer f.Close()

		var t domain.Ticket
		if err := yaml.NewDecoder(f).Decode(&t); err != nil {
			continue
		}
		tickets = append(tickets, t)
	}

	return tickets, nil
}

// Delete removes a ticket workspace
func (e *Engine) Delete(ticketID string) error {
	path := filepath.Join(e.TicketsRoot, ticketID)
	return os.RemoveAll(path)
}
