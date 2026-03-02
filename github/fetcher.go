package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"universal-commits-migrator/interfaces"
	"universal-commits-migrator/logger"
)

var (
	sourceGithubToken    string
	sourceGithubUsername string
	sourceGithubApiUrl   string
)

func init() {
	sourceGithubToken = os.Getenv("SOURCE_GITHUB_TOKEN")
	sourceGithubUsername = os.Getenv("SOURCE_GITHUB_USERNAME")
	sourceGithubApiUrl = os.Getenv("SOURCE_GITHUB_API_URL")
	if sourceGithubApiUrl == "" {
		sourceGithubApiUrl = "https://api.github.com"
	}
}

// Ensure that GitHubSource implements the Source interface
var _ interfaces.Source = (*GitHubSource)(nil)

// GitHubSource is the struct implementing the Source interface for GitHub
type GitHubSource struct{}

// NewGitHubSource initializes a new GitHubSource instance
func NewGitHubSource() interfaces.Source {
	return &GitHubSource{}
}

// FetchAllCommits fetches all commits authored by the user from GitHub
func (g *GitHubSource) FetchAllCommits() ([]interfaces.Commit, error) {
	var allCommits []interfaces.Commit
	page := 1
	perPage := 100

	for {
		url := fmt.Sprintf(
			"%s/search/commits?q=author:%s&per_page=%d&page=%d&sort=committer-date&order=asc",
			sourceGithubApiUrl, sourceGithubUsername, perPage, page,
		)

		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		request.Header.Set("Authorization", "Bearer "+sourceGithubToken)
		request.Header.Set("Accept", "application/vnd.github.cloak-preview+json")
		request.Header.Set("X-GitHub-Api-Version", "2022-11-28")

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			responseBody, _ := ioutil.ReadAll(response.Body)
			formattedErr := fmt.Errorf("received non-200 status code: %d, body: %s", response.StatusCode, string(responseBody))
			logger.Log.Println(formattedErr)
			return nil, formattedErr
		}

		var result struct {
			Items []struct {
				Commit struct {
					Author struct {
						Date string `json:"date"`
					} `json:"author"`
				} `json:"commit"`
			} `json:"items"`
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		logger.Log.Printf("🔍 Found %d commits on page %d\n", len(result.Items), page)

		for _, item := range result.Items {
			allCommits = append(allCommits, interfaces.Commit{Date: item.Commit.Author.Date})
		}

		if len(result.Items) < perPage {
			break
		}
		page++
	}

	return allCommits, nil
}
