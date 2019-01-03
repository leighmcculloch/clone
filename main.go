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
		fmt.Println("Usage: githubclone <repository-name> [repository-name] ...")
		flag.PrintDefaults()
	}
	flag.Parse()

	repos := flag.Args()

	if len(repos) == 0 {
		flag.Usage()
		return
	}

	u, err := user.Current()
	if err != nil {
		color.Red("Error: %s\n", err)
		return
	}

	for _, r := range repos {
		url := fmt.Sprintf("git@github.com:%s/%s", u.Username, r)

		cmd := exec.Command("git", "clone", url)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		color.Green("Cloning repo %s:\n", url)
		err := cmd.Run()
		if err != nil {
			color.Red("Error: %s\n", err)
		}
	}
}
