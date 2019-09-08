package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"

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

	u, err := user.Current()
	if err != nil {
		color.Red("Error: %s\n", err)
		return
	}

	repo := args[0]
	target := repo
	if len(args) > 1 {
		target = args[1]
	}

	repoPath := path.Join(u.Username, repo)

	httpsURL := fmt.Sprintf("https://github.com/%s", repoPath)
	cloneCmd := exec.Command("git", "clone", httpsURL, target)
	cloneCmd.Stdin = os.Stdin
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	color.Green("Cloning repo %s:\n", httpsURL)
	err = cloneCmd.Run()
	if err != nil {
		color.Red("Error: %s\n", err)
		os.Exit(1)
	}

	sshURL := fmt.Sprintf("git@github.com:%s", repoPath)
	setURLCmd := exec.Command("git", "remote", "set-url", "--add", "--push", "origin", sshURL)
	setURLCmd.Dir = filepath.Join(".", target)
	setURLCmd.Stdin = os.Stdin
	setURLCmd.Stdout = os.Stdout
	setURLCmd.Stderr = os.Stderr
	err = setURLCmd.Run()
	if err != nil {
		color.Red("Error: %s\n", err)
		os.Exit(1)
	}
}
