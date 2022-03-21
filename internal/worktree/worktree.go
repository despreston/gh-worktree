package worktree

import (
	"fmt"

	ghapi "github.com/cli/go-gh/pkg/api"
	"github.com/despreston/gh-worktree/internal/worktree/commands/pr"
	"github.com/spf13/cobra"
)

func New(restClient ghapi.RESTClient) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "worktree",
		Short: "Git worktrees, dawg",
		Long:  "commands to create and manage git worktrees",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root worktree command")
		},
	}

	rootCmd.AddCommand(pr.New(restClient))
	return rootCmd
}
