# Design: Yardmaster Architecture

## Architecture Overview

Yardmaster follows a modular architecture to separate concerns between the CLI/TUI, the domain logic, and the underlying git/filesystem operations.

### Components

1.  **Config Layer (`internal/config`)**
    - Uses Viper to load config from `~/.config/yardmaster/config.yaml`.
    - Defines `projects_root` (canonical clones) and `tickets_root` (workspaces).

2.  **Git Engine (`internal/gitx`)**
    - Wraps `go-git`.
    - Manages "Canonical Repos" (bare/mirror clones).
    - Creates Worktrees for tickets.
    - **Constraint**: Must handle worktree creation safely, ensuring branches are named consistently (`<TICKET-ID>`).

3.  **Workspace Engine (`internal/workspace`)**
    - Manages the directory structure in `tickets_root`.
    - Persists ticket metadata (`ticket.yaml`) containing the list of repos and other context.

4.  **Domain Layer (`internal/domain`)**
    - `Ticket`: Struct representing a ticket (ID, Slug, Repos).
    - `Repo`: Struct representing a repository (URL, Name).

5.  **CLI (`cmd/yard`)**
    - Cobra commands mapping to domain actions.

6.  **TUI (`internal/tui`)**
    - Bubble Tea model.
    - **State**: List of tickets, current selection.
    - **Actions**: Open (shell), Close (delete).

## Data Flow

1.  **New Ticket**:
    - User runs `yard ticket new PROJ-123 --repos repo-a`.
    - `Workspace` creates `tickets/PROJ-123`.
    - `Gitx` ensures `repo-a` is cloned in `projects_root`.
    - `Gitx` creates worktree for `repo-a` in `tickets/PROJ-123/repo-a` on branch `PROJ-123`.
    - `Workspace` saves `ticket.yaml`.

2.  **Close Ticket**:
    - User runs `yard ticket close PROJ-123`.
    - `Gitx` checks status (dirty/unpushed).
    - If clean (or forced), `Gitx` removes worktrees and deletes branch `PROJ-123`.
    - `Workspace` removes `tickets/PROJ-123`.
