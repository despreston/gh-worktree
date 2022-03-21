package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh"
	"github.com/despreston/gh-worktree/internal/worktree"
)

func main() {
	rest, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd := worktree.New(rest)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
