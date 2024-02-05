# Configuration Guide

This document provides detailed instructions on how to configure the Universal Commits Migrator (UCM). Proper configuration is crucial for the successful execution of the migrator tool.

The configuration involves setting up environment variables that the tool uses to authenticate and interact with the GitLab and GitHub APIs.

## Configuring Environment Variables

The migrator tool uses a `.env` file to manage configuration settings. Below are the steps to properly set up your `.env` file:

### 1. Create `.env` File

If not already present, create a `.env` file in the root directory of your project. You can start by copying the contents from the `.env.template` file provided in the project repository.

### 2. GitLab Configuration

Configure the following variables related to GitLab:

- `GITLAB_API_URL`: Set this to your GitLab instance's API URL. For GitLab.com users, this will be `https://gitlab.com/api/v4`.
- `GITLAB_TOKEN`: Your personal access token from GitLab with `read_api` scope. [Create a GitLab access token](https://gitlab.com/-/profile/personal_access_tokens).
- `CONTRIBUTED_PROJECTS`: (Optional) A comma-separated list of project IDs you want to fetch commits from. If left empty, the tool will fetch commits from all projects you have access to.

Example:
```plaintext
GITLAB_API_URL=https://gitlab.example.io/api/v4
GITLAB_TOKEN=your_gitlab_token
CONTRIBUTED_PROJECTS=123,456,789
```

### 3. GitHub Configuration

Configure the following variables related to GitHub:

- `GITHUB_API_URL`: Set this to your GitHub repository's API URL. Typically, it should be in the format of `https://api.github.com/repos/your_username/your_repo`.
- `GITHUB_API_VERSION`: The version of the GitHub API you are targeting. Usually, this is set to the latest version.
- `GITHUB_BRANCH_TO_COMMIT_TO`: The branch in your GitHub repository where the commits will be pushed. Default is `master`.
- `GITHUB_USERNAME`: Your GitHub username.
- `GITHUB_EMAIL`: The email associated with your GitHub account.
- `GITHUB_TOKEN`: Your personal access token from GitHub with `repo` scope. [Create a GitHub access token](https://github.com/settings/tokens).

Example:
```plaintext
GITHUB_API_URL=https://api.github.com/repos/your_username/your_repo
GITHUB_API_VERSION=2022-11-28
GITHUB_BRANCH_TO_COMMIT_TO=master
GITHUB_USERNAME=your_github_username
GITHUB_EMAIL=your_email@example.com
GITHUB_TOKEN=your_github_token
```

### 4. Source and Destination Configuration

Specify the source from where the commits are fetched and the destination where the commits are pushed:

- `SOURCE`: The source control provider. Default is `gitlab`. Change it if you implement a different provider like `bitbucket`.
- `DESTINATION`: The destination control provider. Default is `github`. Change it if you implement a different provider.

Example:
```plaintext
SOURCE=gitlab
DESTINATION=github
```

### 5. Save the `.env` File

After setting the variables, save the `.env` file. The migrator tool will automatically read the configurations from this file when it runs.

## Validating the Configuration

It's essential to validate your configuration to ensure that the migrator tool can successfully authenticate and interact with the GitLab and GitHub APIs:

- Ensure that the `.env` file is in the root directory of the project.
- Check that all the tokens and URLs are correct and have the necessary permissions.
- Validate that the `.env` file format is correct, and there are no syntax errors.

After following these steps and validating your configuration, the Universal Commits Migrator (UCM) should be correctly configured and ready to use. Proceed to the usage documentation (`/usage/usage.md`) to learn how to run the migrator tool.