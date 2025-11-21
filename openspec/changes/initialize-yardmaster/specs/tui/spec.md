# Spec: TUI

## ADDED Requirements

#### Requirement: Interactive List
The TUI must display a navigable list of workspaces.

#### Scenario: Navigation
Given a list of workspaces
When I press Down/Up arrows
Then the selection highlight should move.

#### Requirement: Detail View
The TUI must show details for the selected workspace.

#### Scenario: View Details
Given I have selected `PROJ-1`
When I press Enter
Then I should see the list of repos and their status for `PROJ-1`.
