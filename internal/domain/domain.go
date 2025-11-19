package domain

// Repo represents a git repository
type Repo struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

// Ticket represents a work item
type Ticket struct {
	ID    string `yaml:"id"`
	Repos []Repo `yaml:"repos"`
}
