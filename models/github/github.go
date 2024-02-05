package github

// Structs for making requests and parsing responses
type treeResponse struct {
	SHA string `json:"sha"`
}

// LatestCommit Structs to parse the latest commit data
type LatestCommit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Tree struct {
			SHA string `json:"sha"` // This is the tree SHA
		} `json:"tree"`
	} `json:"commit"`
}

type CommitResponse struct {
	SHA string `json:"sha"`
}

type ReferenceUpdate struct {
	SHA string `json:"sha"`
}

type CommitData struct {
	Message string     `json:"message"`
	Tree    string     `json:"tree"`
	Parents []string   `json:"parents"`
	Author  AuthorData `json:"author"`
}

type AuthorData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}
