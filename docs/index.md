# Vault CLI 🔐

Vault CLI is a command-line interface tool designed to streamline the management and interaction with your passwords locally. It provides a user-friendly way to store, retrieve, and manage sensitive passwords directly from your terminal.

## Features

- **Secure Storage**: Safely store sensitive data with robust encryption mechanisms.
- **Easy Retrieval**: Quickly access your stored secrets when needed.
- **User-Friendly Commands**: Intuitive commands that simplify vault operations.

## Installation

### For macOS and Linux (using Homebrew)

To install Vault CLI using Homebrew:

```sh
brew tap fagom/vault https://github.com/fagom/homebrew-vault
```

```sh
brew install fagom/vault/vault
```

### For Windows

#### Manual Installation

For manual installation:

Download the latest release (**vault-windows-amd64.tar.gz**) for your operating system from the [releases page](https://github.com/fagom/vault-cli/releases).
Place the executable in a directory included in your system's `PATH`.

## Usage

After installation, you can use the `vault` command followed by specific subcommands to interact with your vaults. For example:

```sh
vault add secret_name secret_value
vault get secret_name
vault list
```

For a complete list of commands and options, run:

```sh
vault help
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes. For major changes, open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/fagom/vault-cli/blob/main/LICENSE) file for more details.
