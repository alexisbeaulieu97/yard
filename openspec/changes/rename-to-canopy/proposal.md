# Change: Rename Project from Yardmaster to Canopy

## Why
The name "Yardmaster" uses a railroad switching yard metaphor that is obscure and doesn't clearly communicate the tool's purpose. Additionally, the current documentation incorrectly implies the tool is specifically for ticket/JIRA workflows, when it's actually a general-purpose workspace manager. "Canopy" provides a superior metaphor: it represents the bird's-eye view from above the forest that lets you see and manage all your git workspaces and branches. The name is memorable, sounds professional, and captures both the oversight perspective (TUI dashboard) and the tree/branch concepts inherent in git.

The canopy metaphor works on multiple levels:
- **Oversight perspective**: The canopy is your vantage point above the complexity
- **Management layer**: You operate from the canopy level to organize branches below
- **Bird's-eye view**: The TUI provides a literal canopy view of all workspaces
- **Git alignment**: Natural connection to branches, trees, and forest (repository collection)

Note: No users are currently using this tool, so we can make clean breaking changes without migration support.

## What Changes

### Breaking Changes
- Binary name changes from `yard` to `canopy`
- Configuration directory changes from `~/.yard/` to `~/.canopy/`
- Configuration keys: `tickets_root` → `workspaces_root`, `ticket_naming` → `workspace_naming`
- Environment variable prefix changes from `YARD_` to `CANOPY_`
- Metadata file changes from `ticket.yaml` to `workspace.yaml`
- Command aliases: remove "ticket"/"t" aliases entirely

### Code Changes
- Update go.mod module path from `yard` to `canopy`
- Rename `cmd/yard/` directory to `cmd/canopy/`
- Update all Go import paths to reference new module name
- Replace "ticket" terminology with "workspace" throughout codebase:
  - Variable names: `tickets` → `workspaces`, `ticket` → `workspace`
  - Function comments and documentation
  - User-facing error messages and help text
  - CLI command descriptions (main.go: "Ticket-centric" → "Workspace-centric")
  - Integration test configuration keys
  - Command aliases: remove "ticket"/"t" entirely

### Documentation Changes
- Update all documentation to use Canopy branding
- Add metaphor explanation to README introduction
- Update project.md with new project name and purpose
- Remove "per-ticket" and "ticket-centric" language
- Emphasize general-purpose workspace management
- Update all command examples (yard → canopy)
- Update repository name and GitHub references

## Impact
- **Affected specs**: core-architecture (project naming, branding, configuration paths)
- **Affected code files**:
  - `cmd/yard/` → `cmd/canopy/` (entire directory)
  - `go.mod` module path
  - `cmd/canopy/main.go` - root command description
  - `cmd/canopy/init.go` - default config template
  - `cmd/canopy/workspace.go` - command aliases
  - `cmd/canopy/repo.go` - flag descriptions
  - `internal/config/config.go` - config paths and defaults
  - `internal/workspace/workspace.go` - metadata file handling
  - `internal/workspaces/service.go` - variable names and error messages
  - `internal/gitx/git.go` - function comments
  - `test/integration/integration_test.go` - config key names
  - All documentation files (README.md, docs/, openspec/)
