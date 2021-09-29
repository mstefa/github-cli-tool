package ghRepos

type GhRepo struct {
	RepoID      int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Fork        bool   `json:"fork"`
}

type ReposRepository interface {
	GetReposByUserName() ([]GhRepo, error)
}

type ReposCsv interface {
	SaveRepos([]GhRepo)
}

func NewGhRepo(repoID int64, name, description string, fork bool) (r GhRepo) {
	r = GhRepo{
		RepoID:      repoID,
		Name:        name,
		Description: description,
		Fork:        fork,
	}
	return
}
