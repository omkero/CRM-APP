package services

import (
	"crm_system/internal/repository"
)

type Services struct {
	Repo *repository.Repository
}

func NewService(Repo *repository.Repository) *Services {
	return &Services{
		Repo: Repo,
	}
}

func (Service *Services) EntryPoint() int {
	value, err := Service.Repo.SelectUserID()
	if err != nil {
		return 0
	}
	return value
}
