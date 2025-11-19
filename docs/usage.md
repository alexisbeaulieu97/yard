# Usage Guide

## Workflow

1.  **Start a Ticket**:
    When you pick up a JIRA ticket (e.g., `PROJ-123`), run:
    ```bash
    yard ticket new PROJ-123 --repos backend,frontend
    ```
    This creates a workspace at `~/tickets/PROJ-123`.

2.  **Work**:
    ```bash
    cd ~/tickets/PROJ-123
    # Edit files in backend/ and frontend/
    ```
    You are automatically on branch `PROJ-123` in both repos.

3.  **Check Status**:
    ```bash
    yard status
    ```

4.  **Finish**:
    Push your changes using standard git commands inside the worktrees.
    ```bash
    cd backend && git push origin PROJ-123
    ```

5.  **Cleanup**:
    ```bash
    yard ticket close PROJ-123
    ```
    This removes the directory. It will warn you if you have unpushed changes (once fully implemented).
