package interfaces

// Commit represents a commit from any source.
type Commit struct {
	Date string
}

// Source is an interface that all source control providers implement
type Source interface {
	FetchAllCommits() ([]Commit, error)
}
