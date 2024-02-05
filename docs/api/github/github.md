# GitHub API Integration Guide

This document provides an in-depth look at how the Universal Commits Migrator (UCM) interacts with the GitHub API, including creating commits, updating references, and authentication procedures.

## Overview

The migrator uses the GitHub API to create empty commits and update references in a specified GitHub repository. The interactions with the GitHub API are primarily managed in the `github/client.go` and `github/pusher.go` files.

## Authentication

The migrator authenticates with the GitHub API using a Personal Access Token. This token should be specified in your `.env` file under the key `GITHUB_TOKEN`. Ensure that your token has the `repo` scope to allow the tool to create commits and update references in your repository.

## Creating Commits

The migrator creates empty commits on the specified branch of your GitHub repository, preserving the commit dates from the source control (e.g., GitLab). This is handled by the `PushEmptyCommit()` function in `github/client.go`.

- **Endpoint**: `/git/commits`
- **Method**: POST
- **Body**:
    - `message`: A string representing the commit message.
    - `tree`: The SHA of the tree object this commit points to.
    - `parents`: An array of SHAs of the commits that were the parents of this commit.
    - `author`: An object containing information about the commit author.

The `tree` and `parents` values are fetched using the `GetLatestCommitInfo()` function, ensuring that each new commit is added on top of the latest commit in the branch.

## Updating References

After creating a commit, the migrator updates the reference (branch) to point to the new commit. This ensures that the commit history in the repository reflects the newly created commits. This process is managed by the `UpdateReference()` function in `github/client.go`.

- **Endpoint**: `/git/refs/heads/{branch}`
- **Method**: PATCH
- **Body**:
    - `sha`: The SHA of the new commit to which the branch will point.

## Handling Rate Limits

GitHub imposes rate limits on API requests. If you exceed these limits, the API will start returning 403 Forbidden responses. The migrator logs any non-200 status responses, which can be used to diagnose rate-limiting or other issues.

## Error Handling

The migrator checks the HTTP status code of each response. Any status code other than 200 OK or 201 Created is treated as an error. These errors, along with their details, are logged to the `app.log` file for easy troubleshooting.

## Best Practices

- Protect your Personal Access Token and ensure it has the correct scope.
- Be aware of the rate limits imposed by GitHub to prevent service interruptions.
- Regularly monitor the `app.log` file for any errors or unexpected messages.

By understanding the interaction with the GitHub API, you can troubleshoot issues, extend the functionality of the tool, or contribute to its development. If you have suggestions for improvements or encounter any problems, feel free to open an issue or submit a pull request to the project repository.