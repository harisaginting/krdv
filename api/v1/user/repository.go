package user

import (
	"context"

	"github.com/harisaginting/krdv/common/log"
	"github.com/harisaginting/krdv/common/utils/helper"
	"github.com/harisaginting/krdv/db/dao"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func ProviderRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (repo *Repository) GetByUsername(ctx context.Context, username string) (user User, err error) {
	var table dao.User
	qx := repo.db
	qx.Where("username = ?", username).First(&table)
	if qx.Error != nil {
		log.Error(ctx, qx.Error, "FindAllByCustomer: ")
		err = qx.Error
		return
	}
	log.Info(ctx, "Repo : ", table)
	helper.AdjustStructToStruct(table, &user)
	return
}

func (repo *Repository) FindAll(ctx context.Context) (users []User) {
	var table dao.User
	qx := repo.db

	qx.Find(&table)
	if qx.Error != nil {
		log.Error(ctx, qx.Error, "FindAllByCustomer: ")
	}
	log.Info(ctx, "Repo : ", table)
	helper.AdjustStructToStruct(table, &users)
	return
}
