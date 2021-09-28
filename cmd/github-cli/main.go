package main

import (
	ghRepos "github.com/mstefa/github-cli-tool/internal"
	"github.com/mstefa/github-cli-tool/internal/cli"
	"github.com/mstefa/github-cli-tool/internal/repositories/github"
	"github.com/spf13/cobra"
)

func main() {

	var repository ghRepos.ReposRepository

	repository = github.NewGitHubUserRepository()

	rootCmd := &cobra.Command{Use: "github-cil"}
	rootCmd.AddCommand(cli.InitCmd(repository))
	rootCmd.Execute()

}
