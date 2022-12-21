package user

import (
	"context"
)

type Service struct {
	repo Repository
}

func ProviderService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (service *Service) List(ctx context.Context, res *ResponseList) {
	users := service.repo.FindAll(ctx)
	res.Items = users
	res.Total = len(users)
	return
}
