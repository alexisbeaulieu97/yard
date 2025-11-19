# Yardmaster Roadmap

## Phase 1: Initialization (MVP)
**Goal**: Establish the "Walking Skeleton" and core ticket-based workflow.
- [ ] **Core Architecture**: Config, Logging, Git Engine (go-git), Workspace Engine.
- [ ] **CLI**: `init`, `ticket new`, `ticket list`, `ticket close`, `status`.
- [ ] **TUI**: Interactive ticket list and details view.
- [ ] **Git Integration**: Canonical repo management, worktree creation/deletion.

## Phase 2: Enhanced Productivity
**Goal**: Reduce friction and automate manual setup steps.
- [ ] **JIRA Integration**:
    - Fetch ticket title and status from JIRA.
    - Auto-generate slug from title.
    - Link back to JIRA in `ticket info`.
- [ ] **Templates**:
    - `ticket.md` templates based on ticket type (e.g., "Bug", "Feature").
    - Pre-defined repo sets for specific ticket patterns (e.g., `MIDDLEWARE-*`).

## Phase 3: Bulk Operations
**Goal**: Manage multiple repos simultaneously.
- [ ] **Bulk Push**: `yard push <ticket>` pushes branches for all involved repos.
- [ ] **Bulk Check**: `yard check <ticket>` runs status checks or commands (e.g., lint, test) across all repos.

## Phase 4: Advanced Features
- [ ] **Multi-machine Awareness**: Sync workspace metadata to allow reconstructing worktrees on another machine.
- [ ] **Trash Can**: Soft delete for closed tickets to prevent accidental data loss.
