## MODIFIED Requirements
### Requirement: Project Naming and Branding
The system SHALL be named "Canopy" with the binary named `canopy`, using forest/tree metaphors in all user-facing communication.

#### Scenario: Binary installation and invocation
- **WHEN** a user installs the tool via `go install`
- **THEN** the binary is named `canopy` (not `yard` or `yardmaster`)
- **AND** all commands are invoked as `canopy <command>`

#### Scenario: Configuration directory naming
- **WHEN** the system initializes or loads configuration
- **THEN** configuration is stored in `~/.canopy/` directory
- **AND** config file is `~/.canopy/config.yaml`

#### Scenario: Environment variables
- **WHEN** configuration is loaded from environment
- **THEN** environment variables use `CANOPY_` prefix
- **EXAMPLES**: `CANOPY_PROJECTS_ROOT`, `CANOPY_WORKSPACES_ROOT`

#### Scenario: Documentation uses consistent branding
- **WHEN** users read help text, README, or error messages
- **THEN** the project is referred to as "Canopy"
- **AND** metaphors reference canopy, forest, trees, and branches (not railroad/yard terminology)
- **AND** the metaphor explanation appears in the README introduction

## ADDED Requirements
### Requirement: Canopy Metaphor Documentation
The README SHALL include an explanation of the canopy metaphor in the introduction section.

#### Scenario: README metaphor explanation
- **WHEN** a user reads the README introduction
- **THEN** they see an explanation that canopy represents a bird's-eye view above the forest
- **AND** the explanation connects the metaphor to managing git workspaces and branches
- **AND** it clarifies that the TUI provides a literal canopy-level view of all workspaces
