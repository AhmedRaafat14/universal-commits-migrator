package main

import (
	"os"
	"universal-commits-migrator/github"
	"universal-commits-migrator/gitlab"
	"universal-commits-migrator/interfaces"
	"universal-commits-migrator/logger"
)

func main() {
	// Configuration: Determine the source and destination based on user input or configuration
	var source interfaces.Source
	var destination interfaces.Destination

	// Example: Choose the source and destination based on a command line argument or environment variable
	switch os.Getenv("SOURCE") {
	case "gitlab":
		source = gitlab.NewGitLabSource()
	// Add cases for other sources like bitbucket
	default:
		logger.Log.Fatalf("Unsupported source")
	}

	switch os.Getenv("DESTINATION") {
	case "github":
		destination = github.NewGitHubDestination()
	// Add cases for other destinations
	default:
		logger.Log.Fatalf("Unsupported destination")
	}

	logger.Log.Println("Fetching commits from the source...")
	commits, err := source.FetchAllCommits()
	if err != nil {
		logger.Log.Fatalf("Error fetching commits: %s", err)
	}

	var dates []string
	for _, commit := range commits {
		dates = append(dates, commit.Date)
	}

	logger.Log.Println("Pushing commits to the destination...")
	if err := destination.PushCommits(dates); err != nil {
		logger.Log.Fatalf("Error pushing commits: %s", err)
	}

	logger.Log.Println("✳️ Commits have been pushed to the destination.")
}
