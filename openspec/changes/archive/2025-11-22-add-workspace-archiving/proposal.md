# Change: Add Workspace Archiving

## Why
Currently, `canopy workspace close` permanently deletes workspace directories and metadata. Users who want to reference old work or restore a workspace later have no option. Archiving provides a middle ground: remove active worktrees (freeing disk space) while preserving metadata and history for future reference or restoration.

## What Changes
- Add `canopy workspace archive <ID>` command (distinct from close)
- Create `~/.canopy/archive/` directory for archived workspace metadata
- Move workspace metadata to archive directory (preserve workspace.yaml)
- Remove worktrees but keep metadata for potential restoration
- Add `canopy workspace restore <ID>` command to recreate from archive
- Add `canopy workspace list --archived` flag to view archived workspaces
- Update `canopy workspace close` to prompt "Archive instead? [Y/n]"

## Impact
- Affected specs: `specs/workspace-management/spec.md`
- Affected code:
  - `internal/workspaces/service.go` - Add ArchiveWorkspace and RestoreWorkspace methods
  - `internal/workspace/workspace.go` - Add archive directory handling
  - `cmd/canopy/workspace.go` - Add archive, restore, and list --archived commands
  - `internal/config/config.go` - Add archive_root configuration option
