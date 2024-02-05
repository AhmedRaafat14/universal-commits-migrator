package gitlab

import (
	"fmt"
	"strconv"
	"strings"
	"universal-commits-migrator/interfaces"
	"universal-commits-migrator/logger"
	"universal-commits-migrator/models/gitlab"
)

// Ensure that gitlab package implements the Source interface
var _ interfaces.Source = (*GitLabSource)(nil)

// GitLabSource is the struct implementing the Source interface for GitLab
type GitLabSource struct{}

// NewGitLabSource initializes a new GitLabSource instance
func NewGitLabSource() interfaces.Source {
	return &GitLabSource{}
}

// FetchAllCommits fetches all commits from GitLab and implements interfaces.Source
func (g *GitLabSource) FetchAllCommits() ([]interfaces.Commit, error) {
	var projects []gitlab.Project
	var allCommits []interfaces.Commit // Update the slice type to interfaces.Commit

	if contributedProjects != "" {
		projectIDs := strings.Split(contributedProjects, ",")
		logger.Log.Printf("🔍 Found these projects in your environment to fetch: %s\n", contributedProjects)
		for _, idStr := range projectIDs {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				formattedErr := fmt.Errorf("🚨 Invalid project ID in CONTRIBUTED_PROJECTS: %s", idStr)
				logger.Log.Println(formattedErr)
				return nil, formattedErr
			}
			projects = append(projects, gitlab.Project{ID: id})
		}
	} else {
		logger.Log.Printf("🧷 Fetching the projects you contributed to on GitLab.\n")
		var err error
		projects, err = FetchProjects() // Make sure the function name is exported
		if err != nil {
			return nil, err
		}
	}

	for _, project := range projects {
		logger.Log.Printf("🧷 Fetching your commits from project: %d\n", project)
		commits, err := FetchCommitsForProject(project.ID) // Make sure the function name is exported
		if err != nil {
			return nil, err
		}
		logger.Log.Printf("🔍 Found %d commits for you on project %d\n", len(commits), project)
		for _, commit := range commits {
			allCommits = append(allCommits, interfaces.Commit{Date: commit.Date})
		}
	}

	return allCommits, nil
}
