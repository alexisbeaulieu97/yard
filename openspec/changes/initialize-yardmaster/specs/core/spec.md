# Spec: Core Engine

## ADDED Requirements

#### Requirement: Manage Canonical Repos
The system must maintain a directory of canonical git repositories (bare or mirror clones) to serve as the source for worktrees.

#### Scenario: Clone missing repo
Given a requested repo URL that is not in `projects_root`
When a workspace is created using that repo
Then the system should clone it into `projects_root` first.

#### Requirement: Create Workspace Worktrees
The system must be able to create a git worktree for a specific workspace branch.

#### Scenario: Create worktree
Given a canonical repo `repo-a`
When I create a workspace `TICKET-1` involving `repo-a`
Then a worktree should be created at `workspaces_root/TICKET-1/repo-a`
And it should be on branch `TICKET-1`.

#### Requirement: Safe Deletion
The system must prevent accidental data loss when closing workspaces.

#### Scenario: Block deletion on dirty state
Given a workspace `TICKET-1` with uncommitted changes in `repo-a`
When I try to close the workspace
Then the operation should fail with a warning
Unless I provide a force flag.
