# Implementation Tasks

## 1. Module and Package Renaming
- [ ] 1.1 Update `go.mod` module path from `yard` to `canopy`
- [ ] 1.2 Rename `cmd/yard/` directory to `cmd/canopy/`
- [ ] 1.3 Update all Go import paths to reference new module name
- [ ] 1.4 Run `go mod tidy` to verify module changes

## 2. Configuration
- [ ] 2.1 Update config directory constant from `.yard` to `.canopy` in internal/config/config.go
- [ ] 2.2 Update environment variable prefix from `YARD_` to `CANOPY_`
- [ ] 2.3 Update default config in init.go: `tickets_root` → `workspaces_root`, `ticket_naming` → `workspace_naming`
- [ ] 2.4 Update config paths in internal/config/config.go (line 42-43)
- [ ] 2.5 Update metadata file: `ticket.yaml` → `workspace.yaml` in internal/workspace/workspace.go

## 3. Code Terminology Cleanup
- [ ] 3.1 Update variable names in internal/workspaces/service.go: `tickets` → `workspaces` (lines 441-458)
- [ ] 3.2 Update error messages: "tickets" → "workspaces" in service.go
- [ ] 3.3 Update comment in internal/gitx/git.go:46: "ticket branch" → "workspace branch"
- [ ] 3.4 Update integration test: `ticket_patterns` → `workspace_patterns` in test/integration/integration_test.go:90
- [ ] 3.5 Search all Go files for remaining "ticket" references and update

## 4. CLI Updates
- [ ] 4.1 Update root command in main.go: "Ticket-centric workspaces" → "Workspace-centric development"
- [ ] 4.2 Remove "ticket"/"t" aliases from workspace.go:18
- [ ] 4.3 Update repo.go:397: "active tickets" → "active workspaces"
- [ ] 4.4 Update all command help text to use "Canopy" branding
- [ ] 4.5 Update example commands in CLI help (yard → canopy)

## 5. Documentation
- [ ] 5.1 Update README.md: already done (title, metaphor, examples)
- [ ] 5.2 Update openspec/project.md: already done
- [ ] 5.3 Search and replace "Yardmaster" with "Canopy" in all remaining markdown files
- [ ] 5.4 Search and replace "yard" with "canopy" in code examples
- [ ] 5.5 Update docs/usage.md and docs/roadmap.md (if they exist)

## 6. Repository Configuration
- [ ] 6.1 Update GitHub repository name (if applicable)
- [ ] 6.2 Update .gitignore if it references yard-specific paths

## 7. Testing
- [ ] 7.1 Run all tests: `go test ./...`
- [ ] 7.2 Build binary and verify it's named `canopy`
- [ ] 7.3 Test `canopy init` creates `~/.canopy/` directory
- [ ] 7.4 Verify help text displays correctly
- [ ] 7.5 Run `openspec validate rename-to-canopy --strict`
