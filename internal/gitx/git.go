package gitx

import (
	"fmt"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// GitEngine wraps git operations
type GitEngine struct {
	ProjectsRoot string
}

// New creates a new GitEngine
func New(projectsRoot string) *GitEngine {
	return &GitEngine{ProjectsRoot: projectsRoot}
}

// EnsureCanonical ensures the repo is cloned in ProjectsRoot (bare)
func (g *GitEngine) EnsureCanonical(repoURL, repoName string) (*git.Repository, error) {
	path := filepath.Join(g.ProjectsRoot, repoName)

	// Check if exists
	r, err := git.PlainOpen(path)
	if err == nil {
		return r, nil
	}

	// Clone if not exists
	r, err = git.PlainClone(path, true, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to clone %s: %w", repoURL, err)
	}

	return r, nil
}

// CreateWorktree creates a worktree for a ticket branch
func (g *GitEngine) CreateWorktree(repoName, worktreePath, branchName string) error {
	canonicalPath := filepath.Join(g.ProjectsRoot, repoName)
	
	// Clone from local canonical
	_, err := git.PlainClone(worktreePath, false, &git.CloneOptions{
		URL: canonicalPath, 
        ReferenceName: plumbing.NewBranchReferenceName(branchName),
	})
    
    if err != nil {
         // Try cloning default branch then checkout
         r, err := git.PlainClone(worktreePath, false, &git.CloneOptions{
            URL: canonicalPath,
         })
         if err != nil {
             return fmt.Errorf("failed to clone from canonical: %w", err)
         }
         
         w, err := r.Worktree()
         if err != nil {
             return err
         }
         
         // Create branch
         err = w.Checkout(&git.CheckoutOptions{
             Branch: plumbing.NewBranchReferenceName(branchName),
             Create: true,
         })
         if err != nil {
             return fmt.Errorf("failed to checkout branch %s: %w", branchName, err)
         }
    }

	return nil
}

// Status returns isDirty and unpushedCommits count
func (g *GitEngine) Status(path string) (bool, int, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return false, 0, fmt.Errorf("failed to open repo: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return false, 0, fmt.Errorf("failed to get worktree: %w", err)
	}

	status, err := w.Status()
	if err != nil {
		return false, 0, fmt.Errorf("failed to get status: %w", err)
	}

	isDirty := !status.IsClean()

	// Check unpushed commits (ahead of origin)
    // For MVP, we might skip this or just check HEAD vs origin/HEAD
    // This requires fetching first usually.
    
    // Simplified: just return dirty status for now.
	return isDirty, 0, nil
}
