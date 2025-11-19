# Yardmaster

> Ticket-centric worktrees for humans who live in JIRA and git.

Yardmaster (`yard`) is a CLI/TUI tool that manages per-ticket workspaces. It creates isolated directories for each ticket, containing git worktrees for all relevant repositories, while keeping canonical clones centralized.

## Features

- **Ticket Workspaces**: Create a dedicated folder for each ticket (e.g., `~/tickets/PROJ-123`).
- **Git Worktrees**: Automatically create worktrees for multiple repos on the ticket branch.
- **Centralized Storage**: Canonical repos are stored in `~/projects` (configurable) and never re-cloned.
- **TUI**: Interactive terminal UI for managing tickets.

## Installation

```bash
go install github.com/alexisbeaulieu97/yard/cmd/yard@latest
```

## Usage

### Initialization

```bash
yard init
```
Creates `~/.config/yardmaster/config.yaml`.

### Create a Ticket

```bash
yard ticket new PROJ-123 --repos repo-a,repo-b
```
Creates `~/tickets/PROJ-123` and sets up worktrees for `repo-a` and `repo-b`.

### List Tickets

```bash
yard ticket list
```

### Close a Ticket

```bash
yard ticket close PROJ-123
```
Removes the workspace and worktrees.

### TUI

```bash
yard tui
```
Launches the interactive UI.

## Configuration

Edit `~/.config/yardmaster/config.yaml`:

```yaml
projects_root: ~/projects
tickets_root: ~/tickets
ticket_naming: "{{.ID}}"
```