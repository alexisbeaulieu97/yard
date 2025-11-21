# Usage Guide

## Workflow

1.  **Create a Workspace**:
    When you pick up work (e.g., `PROJ-123`), run:
    ```bash
    canopy workspace new PROJ-123 --repos backend,frontend
    ```
    This creates a workspace at `~/workspaces/PROJ-123` (using the default config).

2.  **Work**:
    ```bash
    cd ~/workspaces/PROJ-123
    # Edit files in backend/ and frontend/
    ```
    You are automatically on branch `PROJ-123` in both repos.

3.  **Check Status**:
    ```bash
    canopy status
    ```

4.  **Finish**:
    Push your changes using standard git commands inside the worktrees.
    ```bash
    cd backend && git push origin PROJ-123
    ```

5.  **Cleanup**:
    ```bash
    canopy workspace close PROJ-123
    ```
    This removes the directory. It will warn you if you have unpushed changes (once fully implemented).
