package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"

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
	target := ""
	if len(args) > 1 {
		target = args[1]
	}

	url := fmt.Sprintf("git@github.com:%s/%s", u.Username, repo)

	cloneArgs := []string{"clone", url}
	if target != "" {
		cloneArgs = append(cloneArgs, target)
	}

	cmd := exec.Command("git", cloneArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	color.Green("Cloning repo %s:\n", url)
	err = cmd.Run()
	if err != nil {
		color.Red("Error: %s\n", err)
	}
}
