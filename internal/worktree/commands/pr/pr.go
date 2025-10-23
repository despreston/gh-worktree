package pr

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/cli/go-gh"
	ghapi "github.com/cli/go-gh/pkg/api"
	"github.com/cli/safeexec"
	"github.com/spf13/cobra"
)

func New(restClient ghapi.RESTClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pr [pr number] <path>",
		Short: "worktree from PR",
		Long:  "Create a new worktree from a PR number",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a pr number")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			repo, err := gh.CurrentRepository()
			if err != nil {
				return err
			}

			branchName, err := getPullRequest(restClient, repo.Owner(), repo.Name(), args[0])
			if err != nil {
				return err
			}

			path := branchName
			if len(args) > 1 {
				path = args[1]
			}

			prNumber := args[0]
			return createWorktreeFromPR(prNumber, branchName, path)
		},
	}

	return cmd
}

func getPullRequest(rc ghapi.RESTClient, owner, repo, pr string) (string, error) {
	var response = struct {
		Head struct {
			Ref string
		}
	}{}

	url := fmt.Sprintf("repos/%s/%s/pulls/%s", owner, repo, pr)
	if err := rc.Get(url, &response); err != nil {
		return "", err
	}

	return response.Head.Ref, nil
}

func createWorktreeFromPR(prNumber, branchName string, path string) error {
	// First, fetch the PR ref from GitHub
	fetchCmd := []string{"git", "fetch", "origin", fmt.Sprintf("pull/%s/head:%s", prNumber, branchName)}

	exe, err := safeexec.LookPath(fetchCmd[0])
	if err != nil {
		return err
	}

	fetch := exec.Command(exe, fetchCmd[1:]...)
	fetch.Stdout = os.Stdout
	fetch.Stderr = os.Stderr

	if err := fetch.Run(); err != nil {
		return fmt.Errorf("failed to fetch PR: %w", err)
	}

	// Now create the worktree from the fetched branch
	worktreeCmd := []string{"git", "worktree", "add", path, branchName}

	worktree := exec.Command(exe, worktreeCmd[1:]...)
	worktree.Stdout = os.Stdout
	worktree.Stderr = os.Stderr

	return worktree.Run()
}
