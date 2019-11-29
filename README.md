<div align="center"><img alt="clone" src="README-clone.png" /></div>
<p align="center">
<a href="https://github.com/leighmcculloch/clone/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/leighmcculloch/clone.svg" /></a>
<a href="https://github.com/leighmcculloch/clone/actions"><img alt="Build" src="https://github.com/leighmcculloch/clone/workflows/build/badge.svg" /></a>
<a href="https://goreportcard.com/report/github.com/leighmcculloch/clone"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/leighmcculloch/clone" /></a>
</p>

Clone GitHub repositories conveniently.

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

### Homebrew (Linux; macOS)

```
brew install 4d63/clone/clone
```

### From Source

```
go get 4d63.com/clone
```

## Usage

```
clone [<username>/]<repository-name> [destination-path]
```

## Examples

When cloning repositories matching your username:

```
clone <repository-name>
```

When cloning repositories not matching your username:

```
clone <username>/<repository-name>
```

To clone to a different directory:

```
clone <username>/<repository-name> destination-path
```
