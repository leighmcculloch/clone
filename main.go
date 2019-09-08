package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: githubclone <repository-name> [target]")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 || len(args) > 2 {
		flag.Usage()
		return
	}

	repo := args[0]

	username := ""
	if strings.Contains(repo, "/") {
		components := strings.SplitN(repo, "/", 2)
		username = components[0]
		repo = components[1]
	}

	if username == "" {
		u, err := user.Current()
		if err != nil {
			color.Red("Error: %s\n", err)
			return
		}
		username = u.Username
	}

	target := repo
	if len(args) > 1 {
		target = args[1]
	}

	repoPath := path.Join(username, repo)

	httpsURL := fmt.Sprintf("https://github.com/%s", repoPath)
	cloneCmd := exec.Command("git", "clone", httpsURL, target)
	cloneCmd.Stdin = os.Stdin
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	color.Green("Cloning repo %s:\n", httpsURL)
	if err := cloneCmd.Run(); err != nil {
		color.Red("Error: %s\n", err)
		os.Exit(1)
	}

	sshURL := fmt.Sprintf("git@github.com:%s", repoPath)
	setURLCmd := exec.Command("git", "remote", "set-url", "--add", "--push", "origin", sshURL)
	setURLCmd.Dir = filepath.Join(".", target)
	setURLCmd.Stdin = os.Stdin
	setURLCmd.Stdout = os.Stdout
	setURLCmd.Stderr = os.Stderr
	if err := setURLCmd.Run(); err != nil {
		color.Red("Error: %s\n", err)
		os.Exit(1)
	}
}
