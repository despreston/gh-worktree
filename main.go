package main

import (
	"fmt"
	"os"

	"github.com/despreston/gh-worktree/internal/worktree"
)

func main() {
	cmd, err := worktree.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
