package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func fixRemoteRepoName(repo string) (string, error) {
	// TODO: まじめに書く
	repo = strings.TrimPrefix(repo, "https://")
	return strings.TrimSuffix(repo, ".git"), nil
}

func getRemoteRepos() ([]string, error) {
	b, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		return nil, err
	}

	repoMap := map[string]bool{}

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		if err := s.Err(); err != nil {
			return nil, err
		}

		t := s.Text()
		ss := strings.Fields(t)

		if len(ss) != 3 {
			return nil, errors.New("error on parsing: " + t)
		}

		repo, err := fixRemoteRepoName(ss[1])
		if err != nil {
			return nil, err
		}

		repoMap[repo] = true
	}

	repos := make([]string, 0, len(repoMap))
	for repo := range repoMap {
		repos = append(repos, repo)
	}

	return repos, nil
}

func selectRepo(repos []string) (string, error) {
	prompt := promptui.Select{
		Label:    "repository",
		Items:    repos,
		HideHelp: true,
	}

	_, res, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return res, nil
}

func goModInit(repo string) error {
	fmt.Println("--")
	defer fmt.Println("--")

	fmt.Println("$ go mod init", repo)

	cmd := exec.Command("go", "mod", "init", repo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func run() error {
	repos, err := getRemoteRepos()
	if err != nil {
		return err
	}

	repo, err := selectRepo(repos)
	if err != nil {
		return err
	}

	return goModInit(repo)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
