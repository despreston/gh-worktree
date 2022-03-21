package worktree

import (
	"fmt"

	"github.com/cli/go-gh"
	"github.com/despreston/gh-worktree/internal/worktree/commands/pr"
	"github.com/spf13/cobra"
)

func New() (*cobra.Command, error) {
	rest, err := gh.RESTClient(nil)
	if err != nil {
		return nil, err
	}

	var rootCmd = &cobra.Command{
		Use:   "worktree",
		Short: "Git worktrees, dawg",
		Long:  "commands to create and manage git worktrees",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root worktree command")
		},
	}

	rootCmd.AddCommand(pr.New(rest))
	return rootCmd, nil
}
