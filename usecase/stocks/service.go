package stocks

import repository "github.com/yoshihiro-shu/financial-bot/repository/postgresql"

type Service struct {
	Repository *repository.Queries
}

func NewService(repo *repository.Queries) UseCase {
	return &Service{
		Repository: repo,
	}
}
