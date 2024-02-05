# Installation Guide

Welcome to the installation guide for the Universal Commits Migrator (UCM). This tool is designed to migrate your commit history from source control management platforms like GitLab to a GitHub repository as empty commits.

This document provides detailed instructions to help you set up and install the tool on your local machine.

## Prerequisites

Before you begin the installation process, ensure you have the following prerequisites:

1. **Git**: To clone the repository.
2. **Go**: The tool is written in Go, so you'll need the Go language installed to run it.
3. **Personal Access Tokens**:
    - **GitLab Personal Access Token**: Token with `read_api` scope to fetch commit data from your GitLab account. [Create one here](https://gitlab.com/-/profile/personal_access_tokens).
    - **GitHub Personal Access Token**: Token with `repo` scope to push commits to your GitHub repository. [Create one here](https://github.com/settings/tokens).

## Installation Steps

### 1. Clone the Repository

Clone the Universal Commits Migrator (UCM) repository to your local machine using the following command:

```bash
git clone https://github.com/AhmedRaafat14/universal-commits-migrator.git
```

### 2. Navigate to the Project Directory

Change into the project directory with:

```bash
cd universal-commits-migrator
```

### 3. Configure Environment Variables

The tool uses environment variables for configuration. Rename the provided `.env.template` to `.env` and fill in your details:

```plaintext
# Rename .env.template to .env
mv .env.template .env

# Edit the .env file with your favorite editor (vim, nano, etc.)
nano .env
```

Inside the `.env` file, you will need to provide the following information:

- **GitLab Configuration**:
    - `GITLAB_API_URL`: Your GitLab API URL.
    - `CONTRIBUTED_PROJECTS`: (Optional) Specify project IDs to fetch commits from, separated by commas.
    - `GITLAB_TOKEN`: Your GitLab Personal Access Token.

- **GitHub Configuration**:
    - `GITHUB_API_URL`: Your GitHub API URL (pointing to the specific repository).
    - `GITHUB_API_VERSION`: The GitHub API version (default is 2022-11-28).
    - `GITHUB_BRANCH_TO_COMMIT_TO`: The branch to which you want to push commits (default is `master`).
    - `GITHUB_USERNAME`: Your GitHub username.
    - `GITHUB_EMAIL`: Your GitHub email address.
    - `GITHUB_TOKEN`: Your GitHub Personal Access Token.

Check the Configuration docs for more details on this.

### 4. Install Dependencies

Run the following to ensure the installation of any dependencies:

```bash
go mod tidy
```

This command will download and install any required dependencies for the project.

### 5. Verify the Installation

Ensure everything is set up correctly by running a simple command to check if the application runs without errors:

```bash
go run main.go
```

If you see a log output without errors and the tool displays fetching and pushing steps, your installation is successful.

## Next Steps

After successfully installing the Universal Commits Migrator (UCM), you may proceed to run the tool following the usage instructions detailed in the `/usage/usage.md` document.

For any issues during the installation or any further assistance, please check the `/troubleshooting/troubleshooting.md` document or open an issue in the GitHub repository.

Thank you for installing the Universal Commits Migrator (UCM). Happy migrating!