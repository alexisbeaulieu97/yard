# Proposal: Initialize Canopy

## Goal
Establish the foundational architecture and feature set for **Canopy**, a workspace-centric workspace manager. This includes the CLI structure, configuration management, git engine, workspace engine, and a basic TUI.

## Key Changes
1.  **Scaffold Project**: Go module, directory structure.
2.  **Core Engines**: Config (Viper), Logging, Git (go-git), Workspace.
3.  **CLI**: `init`, `workspace new`, `workspace list`, `workspace close`, `status`.
4.  **TUI**: Interactive workspace list and details.

## Rationale
This initialization sets up the "walking skeleton" of the application, allowing us to verify the core value proposition (workspace-based worktrees) before adding advanced features like JIRA integration.
