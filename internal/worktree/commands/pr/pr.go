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

			return createWorktree(branchName, path)
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

func createWorktree(branchName string, path string) error {
	cmdList := []string{"git", "worktree", "add", path, branchName}

	exe, err := safeexec.LookPath(cmdList[0])
	if err != nil {
		return err
	}

	cmd := exec.Command(exe, cmdList[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
