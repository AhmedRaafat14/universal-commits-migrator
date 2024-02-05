package gitlab

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"universal-commits-migrator/logger"
	"universal-commits-migrator/models/gitlab"
)

var (
	gitlabToken         string
	gitlabApiUrl        string
	contributedProjects string
)

func init() {
	err := godotenv.Load() // Load .env file
	if err != nil {
		logger.Log.Println("Error loading .env file in gitlab package:", err)
	}

	gitlabToken = os.Getenv("GITLAB_TOKEN")
	gitlabApiUrl = os.Getenv("GITLAB_API_URL")
	contributedProjects = os.Getenv("CONTRIBUTED_PROJECTS")
}

func FetchProjects() ([]gitlab.Project, error) {
	var allProjects []gitlab.Project
	page := 1
	perPage := 100

	for {
		url := fmt.Sprintf("%s/projects?page=%d&per_page=%d&membership=true&simple=true", gitlabApiUrl, page, perPage)
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Set("PRIVATE-TOKEN", gitlabToken)

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			formattedErr := fmt.Errorf("received non-200 status code: %d", response.StatusCode)
			logger.Log.Println(formattedErr)
			return nil, formattedErr
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var projects []gitlab.Project
		if err := json.Unmarshal(body, &projects); err != nil {
			return nil, err
		}

		allProjects = append(allProjects, projects...)

		if len(projects) < perPage {
			break
		}
		page++
	}

	return allProjects, nil
}

func FetchCommitsForProject(projectID int) ([]gitlab.Commit, error) {
	var allCommits []gitlab.Commit
	baseURL := fmt.Sprintf("%s/projects/%d/repository/commits", gitlabApiUrl, projectID)

	// Start fetching from the first page
	page := 1
	perPage := 100 // Adjust per_page to your preference

	for {
		url := fmt.Sprintf("%s?page=%d&per_page=%d", baseURL, page, perPage)
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Set("PRIVATE-TOKEN", gitlabToken)

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			formattedErr := fmt.Errorf("received non-200 status code: %d", response.StatusCode)
			logger.Log.Println(formattedErr)
			return nil, formattedErr
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var commits []gitlab.Commit
		if err := json.Unmarshal(body, &commits); err != nil {
			return nil, err
		}

		allCommits = append(allCommits, commits...)

		// Check if we've reached the last page
		if len(commits) < perPage {
			break
		}
		page++
	}

	return allCommits, nil
}
