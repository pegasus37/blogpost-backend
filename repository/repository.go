package repository

type Repository struct {
	Repo RepositoryInterface
}

func NewRepository(repo RepositoryInterface) *Repository {
	r := new(Repository)
	r.Repo = repo
	return r
}
