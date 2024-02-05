## Contributing

Contributions are what make the open-source community an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

Don't forget to give the project a star! Thanks again!

### Adding New Source or Destination Providers

To contribute a new source or destination provider to the project, you'll need to:

1. **Implement the Interface**: Create a new Go file in the appropriate directory (`gitlab`, `github`, `bitbucket`, etc.) and ensure your class implements the `Source` or `Destination` interface from the `interfaces` package.

2. **Define the Model**: If your provider has unique data structures, define them in the `models` directory.

3. **Update `main.go`**: Add a new case to the appropriate `switch` statement in `main.go` for your new provider. This will allow the application to recognize and utilize your new provider based on the environment configuration.

4. **Test Your Implementation**: Ensure that your new provider works as expected and handles errors gracefully.

5. **Document Your Provider**: Update the `ReadMe.md` and any other documentation to include information about the new provider, how to set it up, and any specific configurations it may require.

6. **Submit a Pull Request**: Once you're happy with your implementation, submit a pull request with a clear list of what you've done. The more details, the better!

### Pull Request Process

1. Ensure any install or build dependencies are removed before the end of the layer when doing a build.
2. Update the `ReadMe.md` with details of changes to the interface, this includes new environment variables, and useful file locations.
3. You may merge the Pull Request in once you have the sign-off of two other developers, or if you do not have permission to do that, you may request the second reviewer to merge it for you.

Thank you for contributing to the Universal Commits Migrator (UCM)!
