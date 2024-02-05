package github

import (
	"fmt"
	"os"
	"strings"
	"time"
	"universal-commits-migrator/interfaces"
	"universal-commits-migrator/logger"
)

// Ensure that GitHubDestination implements the Destination interface
var _ interfaces.Destination = (*GitHubDestination)(nil)

// GitHubDestination is the struct implementing the Destination interface for GitHub
type GitHubDestination struct{}

// NewGitHubDestination initializes a new GitHubDestination instance
func NewGitHubDestination() interfaces.Destination {
	return &GitHubDestination{}
}

// PushCommits receives a slice of commits (mainly dates) and pushes empty commits to the repository
func (g *GitHubDestination) PushCommits(dates []string) error {
	// Read the already pushed dates
	pushedDates := readPushedCommits()

	for _, date := range dates {
		logger.Log.Printf("🏁 Check if we pushed commit on date %s before.\n", date)
		if pushedDates[date] {
			logger.Log.Printf("🛂 Commit on date %s has already been pushed, skipping...\n\n\n", date)
			continue
		}

		// Fetch the latest commit SHA and tree SHA
		lastCommitSHA, latestTreeSHA, err := GetLatestCommitInfo()
		if err != nil {
			formattedError := fmt.Errorf("🚨 failed to get the latest commit info: %s", err.Error())
			logger.Log.Println(formattedError)
			return formattedError
		}

		logger.Log.Printf("➕ Committing commit done on date: %s\n", date)

		// Push the empty commit using the latest commit SHA as the parent and the latest tree SHA
		newCommitSHA, err := PushEmptyCommit(date, lastCommitSHA, latestTreeSHA)
		if err != nil {
			return err
		}

		logger.Log.Printf("📌 Pushing commit %s to the master branch.\n", date)

		// Update the reference to point to the new commit
		err = UpdateReference(newCommitSHA)
		if err != nil {
			return err
		}

		// Mark this commit as pushed by adding to the local file and data structure
		logger.Log.Printf("✅ Mark the commit on date %s as pushed.\n\n\n", date)
		markCommitAsPushed(date, pushedDates)
		logger.Log.Printf("####################################################\n\n")

		// Introduce a delay of 3 seconds before the next commit
		logger.Log.Printf("⌛ Waiting for 3 seconds before the next operation to ensure consistency & give GitHub time to update the cache...\n\n")
		time.Sleep(3 * time.Second) // Sleep for 3 seconds
		logger.Log.Printf("####################################################\n\n")
	}
	return nil
}

// readPushedCommits fetch pre-saved commits dates from file.
func readPushedCommits() map[string]bool {
	pushedDates := make(map[string]bool)
	data, err := os.ReadFile("pushed_commits.txt")
	if err != nil {
		// Handle error (e.g., file not found), may decide to ignore if the file doesn't exist
		logger.Log.Printf("🚨 Error reading pushed_commits.txt:%s\n", err)
		return pushedDates
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line != "" {
			pushedDates[line] = true
		}
	}
	return pushedDates
}

// markCommitAsPushed add the pushed commit date to the commits tracking file
func markCommitAsPushed(date string, pushedDates map[string]bool) {
	pushedDates[date] = true
	file, err := os.OpenFile("pushed_commits.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Log.Printf("🚨 Error opening pushed_commits.txt:%s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(date + "\n")
	if err != nil {
		logger.Log.Printf("🚨 Error writing to pushed_commits.txt:%s\n", err)
	}
}
