# Implementation Tasks

## 1. Module and Package Renaming
- [x] 1.1 Update `go.mod` module path from `yard` to `canopy`
- [x] 1.2 Rename `cmd/yard/` directory to `cmd/canopy/`
- [x] 1.3 Update all Go import paths to reference new module name
- [x] 1.4 Run `go mod tidy` to verify module changes

## 2. Configuration
- [x] 2.1 Update config directory constant from `.yard` to `.canopy` in internal/config/config.go
- [x] 2.2 Update environment variable prefix from `YARD_` to `CANOPY_`
- [x] 2.3 Update default config in init.go: `tickets_root` → `workspaces_root`, `ticket_naming` → `workspace_naming`
- [x] 2.4 Update config paths in internal/config/config.go (line 42-43)
- [x] 2.5 Update metadata file: `ticket.yaml` → `workspace.yaml` in internal/workspace/workspace.go

## 3. Code Terminology Cleanup
- [x] 3.1 Update variable names in internal/workspaces/service.go: `tickets` → `workspaces` (lines 441-458)
- [x] 3.2 Update error messages: "tickets" → "workspaces" in service.go
- [x] 3.3 Update comment in internal/gitx/git.go:46: "ticket branch" → "workspace branch"
- [x] 3.4 Update integration test: `ticket_patterns` → `workspace_patterns` in test/integration/integration_test.go:90
- [x] 3.5 Search all Go files for remaining "ticket" references and update

## 4. CLI Updates
- [x] 4.1 Update root command in main.go: "Ticket-centric workspaces" → "Workspace-centric development"
- [x] 4.2 Remove "ticket"/"t" aliases from workspace.go:18
- [x] 4.3 Update repo.go:397: "active tickets" → "active workspaces"
- [x] 4.4 Update all command help text to use "Canopy" branding
- [x] 4.5 Update example commands in CLI help (yard → canopy)

## 5. Documentation
- [x] 5.1 Update README.md: already done (title, metaphor, examples)
- [x] 5.2 Update openspec/project.md: already done
- [x] 5.3 Search and replace "Yardmaster" with "Canopy" in all remaining markdown files
- [x] 5.4 Search and replace "yard" with "canopy" in code examples
- [x] 5.5 Update docs/usage.md and docs/roadmap.md (if they exist)

## 6. Repository Configuration
- [x] 6.1 Update GitHub repository name (if applicable) (not changed here; repository rename is external)
- [x] 6.2 Update .gitignore if it references yard-specific paths

## 7. Testing
- [x] 7.1 Run all tests: `go test ./...`
- [x] 7.2 Build binary and verify it's named `canopy`
- [x] 7.3 Test `canopy init` creates `~/.canopy/` directory
- [x] 7.4 Verify help text displays correctly
- [x] 7.5 Run `openspec validate rename-to-canopy --strict`
