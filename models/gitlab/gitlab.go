package gitlab

// Commit represents a GitLab commit.
type Commit struct {
	Date string `json:"committed_date"`
}

// Project represents a GitLab project.
type Project struct {
	ID int `json:"id"`
}
