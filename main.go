package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
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

	systemUser, err := user.Current()
	if err != nil {
		color.Red("Error: %s\n", err)
		return
	}
	if username == "" {
		username = systemUser.Username
	}

	target := repo
	if username != systemUser.Username {
		target = username + "--" + repo
	}
	if len(args) > 1 {
		target = args[1]
	}

	repoPath := path.Join(username, repo)

	httpsURL := fmt.Sprintf("https://github.com/%s", repoPath)
	sshURL := fmt.Sprintf("git@github.com:%s", repoPath)

	public := true
	headResp, err := http.Head(httpsURL)
	if err != nil {
		color.Yellow("Warning: Unable to check if project is public/private: %s\n", err)
		color.Yellow("Assuming repository is public.\n")
	} else {
		public = headResp.StatusCode == http.StatusOK
	}

	cloneURL := httpsURL
	if !public {
		cloneURL = sshURL
	}
	cloneCmd := exec.Command("git", "clone", cloneURL, target)
	cloneCmd.Stdin = os.Stdin
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	color.Green("Cloning repository %s:\n", cloneURL)
	if err := cloneCmd.Run(); err != nil {
		color.Red("Error cloning: %s\n", err)
		os.Exit(1)
	}

	setURLCmd := exec.Command("git", "remote", "set-url", "--add", "--push", "origin", sshURL)
	setURLCmd.Dir = filepath.Join(".", target)
	setURLCmd.Stdin = os.Stdin
	setURLCmd.Stdout = os.Stdout
	setURLCmd.Stderr = os.Stderr
	if err := setURLCmd.Run(); err != nil {
		color.Red("Error setting origin: %s\n", err)
		os.Exit(1)
	}

	parentCloneURL := ""
	parentSSHURL := ""
	if public {
		apiURL := fmt.Sprintf("https://api.github.com/repos/%s", repoPath)
		apiResp, err := http.Get(apiURL)
		if err != nil {
			color.Yellow("Warning: Unable to check if fork: %s\n", err)
			os.Exit(0)
		}
		if apiResp.StatusCode != http.StatusOK {
			color.Yellow("Warning: Unable to check if fork: %s\n", apiResp.Status)
			os.Exit(0)
		} else {
			defer apiResp.Body.Close()
			apiRespDec := struct {
				Parent struct {
					CloneURL string `json:"clone_url"`
					SSHURL   string `json:"ssh_url"`
				} `json:"parent"`
			}{}
			err := json.NewDecoder(apiResp.Body).Decode(&apiRespDec)
			if err != nil {
				color.Yellow("Warning: Unable to check if fork: %s\n", err)
				os.Exit(0)
			}
			if apiRespDec.Parent.SSHURL != "" {
				parentCloneURL = apiRespDec.Parent.CloneURL
				parentSSHURL = apiRespDec.Parent.SSHURL
			}
		}
	}

	if !public && strings.Contains(repo, "--") {
		repoParts := strings.SplitN(repo, "--", 2)
		parentRepoPath := path.Join(repoParts...)
		parentURL := fmt.Sprintf("git@github.com:%s", parentRepoPath)
		parentCloneURL = parentURL
		parentSSHURL = parentURL
	}

	if parentCloneURL != "" && parentSSHURL != "" {
		setUpstreamCmd := exec.Command("git", "remote", "add", "upstream", parentCloneURL)
		setUpstreamCmd.Dir = filepath.Join(".", target)
		setUpstreamCmd.Stdin = os.Stdin
		setUpstreamCmd.Stdout = os.Stdout
		setUpstreamCmd.Stderr = os.Stderr
		if err := setUpstreamCmd.Run(); err != nil {
			color.Red("Error setting upstream: %s\n", err)
			os.Exit(1)
		}
		setUpstreamPushCmd := exec.Command("git", "remote", "set-url", "--add", "--push", "upstream", parentSSHURL)
		setUpstreamPushCmd.Dir = filepath.Join(".", target)
		setUpstreamPushCmd.Stdin = os.Stdin
		setUpstreamPushCmd.Stdout = os.Stdout
		setUpstreamPushCmd.Stderr = os.Stderr
		if err := setUpstreamPushCmd.Run(); err != nil {
			color.Red("Error setting upstream: %s\n", err)
			os.Exit(1)
		}
	}
}
