# GitLab API Integration Guide

This document provides details on how the Universal Commits Migrator (UCM) interacts with the GitLab API, including fetching projects and commits, authentication, and handling rate limits.

## Overview

The migrator uses the GitLab API to fetch the list of projects you have contributed to and the commits for each project. The main interactions with the GitLab API occur in the `gitlab/client.go` and `gitlab/fetcher.go` files.

## Authentication

The migrator uses a Personal Access Token to authenticate with the GitLab API. You need to provide this token in your `.env` file under the key `GITLAB_TOKEN`. The token should have the `read_api` scope to allow the tool to fetch commit data.

## Fetching Projects

If you want to fetch commits from specific projects, you can list their IDs in the `CONTRIBUTED_PROJECTS` environment variable. If this variable is left empty, the migrator fetches all the projects you have access to.

The `FetchProjects()` function in `gitlab/client.go` handles fetching the project list:

- **Endpoint**: `/projects`
- **Method**: GET
- **Parameters**:
    - `membership=true` to filter projects that you are a member of.
    - `simple=true` to get a simplified version of the projects' details.
    - Pagination parameters like `page` and `per_page`.

## Fetching Commits

The `FetchCommitsForProject()` function in `gitlab/client.go` fetches commits for a specific project:

- **Endpoint**: `/projects/{projectID}/repository/commits`
- **Method**: GET
- **Parameters**:
    - `page` and `per_page` for pagination.

The function loops through the paginated results until all commits are fetched.

## Rate Limits

GitLab imposes rate limits on API requests. If you encounter rate-limiting issues, you might see errors related to exceeding these limits. In such cases, you might need to wait or adjust the frequency of your requests.

## Error Handling

The code checks the HTTP status code of each response. Any status code other than 200 OK is considered an error, and the tool logs the error details. The logs can be found in the `app.log` file for further investigation.

## Best Practices

- Ensure your Personal Access Token is kept secure and has only the required scopes.
- Be mindful of the rate limits imposed by GitLab to avoid service disruptions.
- Regularly check the log file for any errors or unusual activity.

---

By understanding the interaction with the GitLab API, you can troubleshoot issues, extend the tool's capabilities, or even contribute to its improvement. If you have any enhancements or encounter any issues, feel free to contribute to the project repository.