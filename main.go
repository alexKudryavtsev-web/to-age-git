package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: to-age-git <repository_path> <number_of_years>")
		os.Exit(1)
	}

	repoPath := os.Args[1]
	years, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid number of years: %v\n", err)
		os.Exit(1)
	}

	if !isGitRepo(repoPath) {
		fmt.Println("The specified folder is not a git repository")
		os.Exit(1)
	}

	commits, err := getAllCommits(repoPath)
	if err != nil {
		fmt.Printf("Error getting commits: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d commits to modify\n", len(commits))

	for _, commit := range commits {
		err := changeCommitDate(repoPath, commit, years)
		if err != nil {
			fmt.Printf("Error modifying commit %s: %v\n", commit, err)
			os.Exit(1)
		}
	}

	fmt.Println("All commits successfully modified!")
}

func isGitRepo(path string) bool {
	gitPath := filepath.Join(path, ".git")
	_, err := os.Stat(gitPath)
	return !os.IsNotExist(err)
}

func getAllCommits(repoPath string) ([]string, error) {
	// git rev-list --all
	cmd := exec.Command("git", "rev-list", "--all")
	cmd.Dir = repoPath
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	commits := strings.Split(strings.TrimSpace(string(output)), "\n")
	return commits, nil
}

func changeCommitDate(repoPath, commitHash string, years int) error {
	// git show -s --format=%at
	cmd := exec.Command("git", "show", "-s", "--format=%at", commitHash)
	cmd.Dir = repoPath
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	timestamp := strings.TrimSpace(string(output))
	unixTime, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return err
	}

	commitTime := time.Unix(unixTime, 0)
	newTime := commitTime.AddDate(-years, 0, 0)
	newTimestamp := fmt.Sprintf("%d", newTime.Unix())

	envVars := []string{
		fmt.Sprintf("GIT_COMMITTER_DATE=%s", newTimestamp),
		fmt.Sprintf("GIT_AUTHOR_DATE=%s", newTimestamp),
	}

	filter := fmt.Sprintf("git filter-branch -f --env-filter "+
		"'if [ $GIT_COMMIT = %s ]; then "+
		"export GIT_AUTHOR_DATE=\"%s\"; "+
		"export GIT_COMMITTER_DATE=\"%s\"; "+
		"fi' HEAD", commitHash, newTimestamp, newTimestamp)

	cmd = exec.Command("bash", "-c", filter)
	cmd.Dir = repoPath
	cmd.Env = append(os.Environ(), envVars...)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
