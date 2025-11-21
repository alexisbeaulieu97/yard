# Canopy Roadmap

## Phase 1: Initialization (MVP)
**Goal**: Establish the "Walking Skeleton" and core workspace workflow.
- [ ] **Core Architecture**: Config, Logging, Git Engine (go-git), Workspace Engine.
- [ ] **CLI**: `init`, `workspace new`, `workspace list`, `workspace close`, `status`.
- [ ] **TUI**: Interactive workspace list and details view.
- [ ] **Git Integration**: Canonical repo management, worktree creation/deletion.

## Phase 2: Enhanced Productivity
**Goal**: Reduce friction and automate manual setup steps.
- [ ] **Templates**:
    - Workspace templates based on work type (e.g., "Bug", "Feature").
    - Pre-defined repo sets for specific workspace patterns (e.g., `MIDDLEWARE-*`).

## Phase 3: Bulk Operations
**Goal**: Manage multiple repos simultaneously.
- [ ] **Bulk Push**: `canopy push <workspace>` pushes branches for all involved repos.
- [ ] **Bulk Check**: `canopy check <workspace>` runs status checks or commands (e.g., lint, test) across all repos.

## Phase 4: Advanced Features
- [ ] **Multi-machine Awareness**: Sync workspace metadata to allow reconstructing worktrees on another machine.
- [ ] **Trash Can**: Soft delete for closed workspaces to prevent accidental data loss.
