# githubclone

Clone personal GitHub repositories conveniently.

For public repositories:
- Clone uses `https`. e.g. `https://github.com/leighmcculloch/githubclone`
- Push uses `ssh`. e.g. `git@github.com:leighmcculloch/githubclone`
- Automatically configures an `upstream` remote to the parent if repo is a fork.

For private repositories:
- Clone uses `ssh`.
- Push uses `ssh`.

## Install

### Binary (Linux; macOS; Windows)

Download and install the binary from the [releases](https://github.com/leighmcculloch/githubclone/releases) page.

### From Source

```
go get 4d63.com/githubclone
```

## Usage

```
githubclone [<username>/]<repository-name> [destination-path]
```

## Examples

When cloning repositories matching your username:

```
githubclone <repository-name>
```

When cloning repositories not matching your username:

```
githubclone <username>/<repository-name>
```

To clone to a different directory:

```
githubclone <username>/<repository-name> destination-path
```
