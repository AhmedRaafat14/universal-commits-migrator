# Troubleshooting Guide

This troubleshooting guide is designed to help you diagnose and resolve common issues you may encounter while using the Universal Commits Migrator (UCM). If you're facing a problem, follow the steps and checks outlined below to resolve it.

## Common Issues and Solutions

### 1. Authentication Failures

**Symptoms**:
- Errors stating `Authentication failed` or `Invalid credentials`.
- The migrator tool cannot fetch data from GitLab or push data to GitHub.

**Checks and Solutions**:
- Ensure that the `GITLAB_TOKEN` and `GITHUB_TOKEN` in your `.env` file are correct and have the required scopes (`read_api` for GitLab, `repo` for GitHub).
- Check if the tokens have expired or been revoked. If so, generate new tokens.

### 2. API URL Misconfiguration

**Symptoms**:
- Errors stating `URL not found` or `API endpoint not reachable`.

**Checks and Solutions**:
- Verify that `GITLAB_API_URL` and `GITHUB_API_URL` are correctly set in your `.env` file.
- For GitHub, the URL should follow the pattern `https://api.github.com/repos/your_username/your_repo`.
- For GitLab, the URL should be `https://gitlab.yourdomain.io/api/v4` or `https://gitlab.com/api/v4` for GitLab.com users.

### 3. Branch Name Issues

**Symptoms**:
- Errors related to branch not found or unable to push to the branch.

**Checks and Solutions**:
- Ensure that the branch specified in `GITHUB_BRANCH_TO_COMMIT_TO` exists in your GitHub repository.
- Check for typos in the branch name.

### 4. Rate Limiting or API Quotas

**Symptoms**:
- Errors stating `Rate limit exceeded` or similar messages.

**Checks and Solutions**:
- Both GitLab and GitHub have rate limits on their APIs. If you encounter rate limiting issues, you may need to wait for the limit to reset.
- Consider reducing the frequency of your requests or distributing the migration over a longer period.

### 5. `pushed_commits.txt` File Issues

**Symptoms**:
- Commits are not being pushed, and the log indicates that the commit has already been pushed.

**Checks and Solutions**:
- Check the `pushed_commits.txt` file to see if the commit dates are recorded correctly.
- If you believe a commit has been wrongly marked as pushed, you can manually edit the `pushed_commits.txt` file, but do so with caution.

### 6. Log File for More Information

If none of the above solutions resolve your issue, refer to the `app.log` file for more detailed error messages and stack traces. The log file may provide insights into what's going wrong.

### 7. Reporting Issues

If you've gone through the above checks and are still facing issues, you may report it by:

- Opening an issue in the project's GitHub repository.
- Providing detailed information, including error messages, logs, and steps to reproduce the issue.

Remember to remove or obfuscate any sensitive information like tokens or personal email addresses before sharing logs or error messages.

---

We hope this troubleshooting guide helps you resolve any issues you encounter. If you have a problem that's not covered here, don't hesitate to reach out for help or open an issue on GitHub.