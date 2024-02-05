package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"universal-commits-migrator/logger"
	"universal-commits-migrator/models/github"
)

var (
	githubApiUrl     string
	githubApiVersion string
	githubToken      string
	authorName       string
	authorEmail      string
	branchToCommitTo string
)

func init() {
	err := godotenv.Load() // You may need to specify the path if .env is not in the root directory
	if err != nil {
		logger.Log.Println("Error loading .env file in github package:", err)
	}

	githubApiUrl = os.Getenv("GITHUB_API_URL")
	githubApiVersion = os.Getenv("GITHUB_API_VERSION")
	githubToken = os.Getenv("GITHUB_TOKEN")
	authorName = os.Getenv("GITHUB_USERNAME")
	authorEmail = os.Getenv("GITHUB_EMAIL")
	branchToCommitTo = os.Getenv("GITHUB_BRANCH_TO_COMMIT_TO")
}

func GetLatestCommitInfo() (commitSHA string, treeSHA string, err error) {
	url := fmt.Sprintf("%s/commits/%s", githubApiUrl, branchToCommitTo) // Replace 'master' with your branch name if it's different
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}

	request.Header.Set("Authorization", "Bearer "+githubToken)
	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("X-GitHub-Api-Version", githubApiVersion)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseBody, _ := ioutil.ReadAll(response.Body)
		formattedErr := fmt.Errorf("received non-200 status code: %d, body: %s", response.StatusCode, string(responseBody))
		logger.Log.Println(formattedErr)
		return "", "", formattedErr
	}

	var latestCommit github.LatestCommit
	if err := json.NewDecoder(response.Body).Decode(&latestCommit); err != nil {
		return "", "", err
	}

	return latestCommit.SHA, latestCommit.Commit.Tree.SHA, nil
}

func PushEmptyCommit(date string, parentCommitSHA string, latestTreeSHA string) (string, error) {
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return "", err
	}
	formattedDate := parsedDate.Format(time.RFC3339)

	commit := github.CommitData{
		Message: fmt.Sprintf("Contribution message on date: %s", date),
		Tree:    latestTreeSHA,
		Parents: []string{parentCommitSHA},
		Author: github.AuthorData{
			Name:  authorName,
			Email: authorEmail,
			Date:  formattedDate,
		},
	}

	commitJSON, err := json.Marshal(commit)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/git/commits", githubApiUrl)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(commitJSON))
	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", "Bearer "+githubToken)
	request.Header.Set("Content-Type", "application/vnd.github+json")
	request.Header.Set("X-GitHub-Api-Version", githubApiVersion)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		responseBody, _ := ioutil.ReadAll(response.Body)
		formattedErr := fmt.Errorf("received non-201 status code: %d, body: %s", response.StatusCode, string(responseBody))
		logger.Log.Println(formattedErr)
		return "", formattedErr
	}

	var commitResp github.CommitResponse
	if err := json.NewDecoder(response.Body).Decode(&commitResp); err != nil {
		return "", err
	}

	return commitResp.SHA, nil
}

func UpdateReference(newCommitSHA string) error {
	updateData := github.ReferenceUpdate{SHA: newCommitSHA}

	updateJSON, err := json.Marshal(updateData)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/git/refs/heads/%s", githubApiUrl, branchToCommitTo) // Replace 'master' with your branch name if it's different
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(updateJSON))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "Bearer "+githubToken)
	request.Header.Set("Content-Type", "application/vnd.github+json")
	request.Header.Set("X-GitHub-Api-Version", githubApiVersion)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseBody, _ := ioutil.ReadAll(response.Body)
		formattedErr := fmt.Errorf("received non-200 status code: %d, body: %s", response.StatusCode, string(responseBody))
		logger.Log.Println(formattedErr)
		return formattedErr
	}

	return nil
}
