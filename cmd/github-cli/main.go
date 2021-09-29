package main

import (
	ghRepos "github.com/mstefa/github-cli-tool/internal"
	"github.com/mstefa/github-cli-tool/internal/cli"
	"github.com/mstefa/github-cli-tool/internal/csv"
	"github.com/mstefa/github-cli-tool/internal/repositories/github"
	"github.com/spf13/cobra"
)

func main() {

	var repository ghRepos.ReposRepository
	var csvRepository ghRepos.ReposCsv

	repository = github.NewGitHubUserRepository()
	csvRepository = csv.NewRepository()

	rootCmd := &cobra.Command{Use: "github-cil"}
	rootCmd.AddCommand(cli.InitCmd(repository, csvRepository))
	rootCmd.Execute()

}
