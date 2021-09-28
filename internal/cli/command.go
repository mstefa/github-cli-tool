package cli

import (
	"fmt"

	ghRepos "github.com/mstefa/github-cli-tool/internal"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

func InitCmd(repository ghRepos.ReposRepository) *cobra.Command {
	usersCmd := &cobra.Command{
		Use:   "user",
		Short: "Fecht gitHub user data according the parameter it recibe",
		Run:   ghUsersFn(repository),
	}

	return usersCmd
}

func ghUsersFn(repository ghRepos.ReposRepository) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		repos, _ := repository.GetReposByUserName()
		fmt.Println("hola!")
		fmt.Println(repos)
	}
}
