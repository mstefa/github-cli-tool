package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	ghRepos "github.com/mstefa/github-cli-tool/internal"
)

const (
	apiURL        = "https://api.github.com/users/"
	userName      = "mstefa"
	reposEndpoint = "/repos"
)

type ghRepository struct {
	url string
}

func NewGitHubUserRepository() ghRepos.ReposRepository {
	return &ghRepository{url: apiURL}
}

func (r *ghRepository) GetReposByUserName() (repos []ghRepos.GhRepo, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v%v", r.url, userName, reposEndpoint))
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &repos)
	if err != nil {
		return nil, err
	}
	return
}
