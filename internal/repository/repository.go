package repository

type Repository struct {
}

func NewRpository() *Repository {

	return &Repository{}
}

func (repo *Repository) GetUserID() int {
	return 1
}

func (repo *Repository) SelectUserID() (int, error) {
	return 15, nil
}
