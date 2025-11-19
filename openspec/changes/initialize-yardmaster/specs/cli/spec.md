# Spec: CLI

## ADDED Requirements

#### Requirement: Initialize Config
The `init` command must setup the global configuration.

#### Scenario: First run
Given no config exists
When I run `yard init`
Then `~/.config/yardmaster/config.yaml` should be created with default paths.

#### Requirement: Create Ticket
The `ticket new` command must create a new workspace.

#### Scenario: New ticket
When I run `yard ticket new PROJ-123 --repos repo-a,repo-b`
Then a workspace `tickets/PROJ-123` should be created
And it should contain worktrees for `repo-a` and `repo-b`.

#### Requirement: List Tickets
The `ticket list` command must show active workspaces.

#### Scenario: List
Given active tickets `PROJ-1` and `PROJ-2`
When I run `yard ticket list`
Then I should see both tickets in the output.
