# Proposal: Initialize Yardmaster

## Goal
Establish the foundational architecture and feature set for **Yardmaster**, a ticket-centric workspace manager. This includes the CLI structure, configuration management, git engine, workspace engine, and a basic TUI.

## Key Changes
1.  **Scaffold Project**: Go module, directory structure.
2.  **Core Engines**: Config (Viper), Logging, Git (go-git), Workspace.
3.  **CLI**: `init`, `ticket new`, `ticket list`, `ticket close`, `status`.
4.  **TUI**: Interactive ticket list and details.

## Rationale
This initialization sets up the "walking skeleton" of the application, allowing us to verify the core value proposition (ticket-based worktrees) before adding advanced features like JIRA integration.
