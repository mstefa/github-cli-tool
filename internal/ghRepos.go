package ghRepos

type GhRepo struct {
	RepoID      int    `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Fork        bool   `json: "fork"`
}

type ReposRepository interface {
	GetReposByUserName() ([]GhRepo, error)
}

func NewGhRepo(repoID int, name, description string, fork bool) (r GhRepo) {
	r = GhRepo{
		RepoID:      repoID,
		Name:        name,
		Description: description,
		Fork:        fork,
	}
	return
}
