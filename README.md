# newyear

Update the copyright year in all of your GitHub repositories.

## How does it work?

1. Clone all repositories owned by your account in a temporary directory. newyear will only clone repositories that aren't forks and aren't archived.
2. Change 2021 to 2022 in any files with the word "license" in them (not case sensitive).
3. If changed, commit the change and push it.

Pretty simple but trust me, when you have a lot of repositories it saves you a ton of time.

## Installing

- macOS: `brew install gleich/homebrew-tap/newyear`.
- Linux and Windows: download a binary from the [releases page](https://go.mattglei.ch/newyear/releases/latest).

## Basic Usage

### Creating a PAT (personal access token)

Before you can use newyear you need to [create a personal access token](https://github.com/settings/tokens/new) from GitHub. Please select the `repo` scope. To be able to load private repositories, this top-level scope needs to be selected. If you don't feel comfortable using this top-level scope you can tick none of the boxes to have newyear just update public repositories.

### Running the CLI

Open up your terminal and run the following terminal command:

```bash
newyear 2021 2022
```

This will update all copyright years from 2021 to 2022

## Options

### Cloning with SSH

To clone with SSH instead of HTTPS please add the following flag to the end of the command like so:

```bash
newyear 2021 2022 --ssh
```
