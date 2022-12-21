package auth

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/harisaginting/krdv/common/cache"
	"github.com/harisaginting/krdv/common/utils/helper"
)

type Service struct {
	repo Repository
}

func ProviderService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (service *Service) Register(ctx context.Context, p PayloadUserRegister) (err error, res ResponseUserRegister) {
	check, _ := service.repo.FindByUsername(ctx, p.Username)
	if check.ID != 0 {
		err = errors.New("username already used")
		return
	}

	p.Password, err = helper.HashPassword([]byte(p.Password))
	if err != nil {
		return
	}

	err = service.repo.Register(ctx, p)
	if err != nil {
		return
	}

	token := uuid.NewString()
	cacheKey := cache.CreateCacheKey("auth:" + token)
	cacheData, err := json.Marshal(p.Username)
	if err != nil {
		return
	}
	err = cache.SetKeyWithExpired(cacheKey, cacheData, "120m")
	if err != nil {
		return
	}
	res.Token = token
	return
}

func (service *Service) Login(ctx context.Context, p PayloadUserLogin) (err error, res ResponseUserLogin) {
	check, _ := service.repo.FindByUsername(ctx, p.Username)
	if check.ID == 0 {
		err = errors.New("invalid Username")
		return
	}

	valid := helper.ComparePasswords(p.Password, check.Password)
	if !valid {
		err = errors.New("invalid Password")
		return
	}

	helper.AdjustStructToStruct(check, &res)
	token := uuid.NewString()
	cacheKey := cache.CreateCacheKey("auth:" + token)
	cacheData, err := json.Marshal(p.Username)
	if err != nil {
		return
	}
	err = cache.SetKeyWithExpired(cacheKey, cacheData, "120m")
	if err != nil {
		return
	}
	res.Token = token
	return
}

func (service *Service) GetByUsername(ctx context.Context, p string) (err error, res ResponseMe) {
	user, err := service.repo.FindByUsername(ctx, p)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errors.New("username not found")
		return
	}
	helper.AdjustStructToStruct(user, &res)
	return
}

func (service *Service) List(ctx context.Context, res *ResponseList) {
	users := service.repo.FindAll(ctx)
	res.Items = users
	res.Total = len(users)
	return
}
