# Usage Guide

This document explains how to use the Universal Commits Migrator (UCM) once you have successfully installed it on your local machine.

The tool is designed to fetch your commit history from a source control management platform (like GitLab) and migrate it to your GitHub repository as empty commits, preserving the original commit timestamps.

## Running the Migrator

Follow these steps to run the Universal Commits Migrator (UCM):

### 1. Prepare Your Environment

Ensure that your `.env` file is correctly configured with your source (GitLab) and destination (GitHub) credentials, as well as any other necessary configuration details.

### 2. Execute the Migration Script

Open a terminal in the root directory of the project and execute the following command:

```bash
go run main.go
```

This command starts the migration process. The tool performs the following actions:

- **Determine Source & Destination**: Reads the `SOURCE` and `DESTINATION` environment variables to determine the source (e.g., GitLab) and destination (e.g., GitHub) for migration.
- **Fetch Commits**: Connects to the source platform using the provided API token and fetches your commit history.
- **Process Commits**: Checks each commit's date against the `pushed_commits.txt` file to determine if it has already been migrated.
- **Push Empty Commits**: For each new commit, creates an empty commit with the same timestamp and pushes it to the specified GitHub repository.
- **Update `pushed_commits.txt`**: After a commit is successfully pushed to GitHub, its date is recorded in the `pushed_commits.txt` file to prevent reprocessing in future runs.

### 3. Monitor the Progress

The migration process can take some time, depending on the number of commits to migrate. Monitor the `app.log` for log messages. The tool provides detailed logs for each operation, making it easy to track the progress and troubleshoot if needed.

### 4. Verify the Migration

Once the script completes, you can verify that the commits have been successfully pushed to your GitHub repository by:

- Checking the commit history in your GitHub repository's web interface.
- Looking for the completion message in the `app.log` file.

### 5. Check the Log File

The tool logs all operations to an `app.log` file. If you encounter any issues or need to review the actions performed by the tool, you can find detailed logs in this file.

## Best Practices

- **Run Incrementally**: The tool is designed to be idempotent, meaning you can run it multiple times without duplicating commits. It's a good practice to run the tool periodically to keep your GitHub repository up-to-date with your latest contributions.
- **Backup**: Before running the tool for the first time, consider backing up your GitHub repository. This ensures that you can restore the repository to its previous state if needed.
- **Review Logs**: Regularly review the `app.log` file for any errors or unexpected behavior.

## Troubleshooting

If you encounter any issues while running the Universal Commits Migrator (UCM), refer to the `/troubleshooting/troubleshooting.md` document for guidance on common issues and their solutions.

## Support and Contributions

If you need help, have suggestions, or would like to contribute to the project, please feel free to open an issue or submit a pull request on the GitHub repository.

Thank you for using the Universal Commits Migrator (UCM). Your contributions and feedback help improve this tool for everyone!