# Spec: CLI

## ADDED Requirements

#### Requirement: Initialize Config
The `init` command must setup the global configuration.

#### Scenario: First run
Given no config exists
When I run `canopy init`
Then `~/.config/canopy/config.yaml` should be created with default paths.

#### Requirement: Create Workspace
The `workspace new` command must create a new workspace.

#### Scenario: New workspace
When I run `canopy workspace new PROJ-123 --repos repo-a,repo-b`
Then a workspace `workspaces/PROJ-123` should be created
And it should contain worktrees for `repo-a` and `repo-b`.

#### Requirement: List Workspaces
The `workspace list` command must show active workspaces.

#### Scenario: List
Given active workspaces `PROJ-1` and `PROJ-2`
When I run `canopy workspace list`
Then I should see both workspaces in the output.
