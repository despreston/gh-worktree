package workspace

import (
	"fmt"

	ghapi "github.com/cli/go-gh/pkg/api"
	"github.com/despreston/gh-workspace/internal/workspace/commands/pr"
	"github.com/spf13/cobra"
)

func New(restClient ghapi.RESTClient) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "workspace",
		Short: "Git workspaces, dawg",
		Long:  "commands to create and manage git workspaces",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root workspace command")
		},
	}

	rootCmd.AddCommand(pr.New(restClient))
	return rootCmd
}
