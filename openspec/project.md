# Project Context

## Purpose
**Yardmaster** (binary `yard`) is a ticket-centric workspace manager for engineers. It solves the problem of juggling multiple JIRA tickets and repositories by creating per-ticket directories containing git worktrees, while keeping canonical clones centralized.

## Tech Stack
- **Language**: Go
- **CLI**: Cobra
- **Config**: Viper
- **TUI**: Bubble Tea, Bubbles, Lipgloss
- **Git**: go-git
- **Logging**: charmbracelet/log
- **Filesystem**: afero (for testing)

## Project Conventions

### Code Style
- Standard Go formatting (`gofmt`).
- Linter: `golangci-lint`.

### Architecture Patterns
- **Hexagonal/Clean Architecture**:
    - `cmd/`: Entry points.
    - `internal/domain/`: Core logic (Ticket, Repo entities).
    - `internal/gitx/`: Git adapter.
    - `internal/workspace/`: Filesystem adapter.
    - `internal/config/`: Configuration adapter.
    - `internal/tui/`: UI adapter.

### Testing Strategy
- Unit tests for domain and adapters.
- `afero` for filesystem mocking.
- Integration tests for git operations.

### Git Workflow
- Feature branches.
- Conventional Commits.

## Domain Context
- **Canonical Repo**: A bare or mirror clone managed by Yardmaster in `projects_root`.
- **Ticket Workspace**: A directory in `tickets_root` named after the ticket (e.g., `tickets/PROJ-123`).
- **Worktree**: A git worktree checked out to a branch named after the ticket.

## Important Constraints
- **No shelling out to git**: Use `go-git` for core operations to ensure portability and testability.
- **Safe Deletion**: `ticket close` must verify no unpushed/uncommitted changes before deletion.
