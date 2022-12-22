package auth

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

func (repo *Repository) FindAll(ctx context.Context) (users []User) {
	var user dao.User
	qx := repo.db
	qx.Find(&user)
	if qx.Error != nil {
		log.Error(ctx, qx.Error, "FindAllByCustomer: ")
	}
	log.Info(ctx, "Repo : ", user)
	helper.AdjustStructToStruct(user, &users)
	return
}

func (repo *Repository) FindByUsername(ctx context.Context, username string) (user dao.User, err error) {
	log.Info(ctx, "repo u : "+username)
	qx := repo.db.Where("username = ?", username).First(&user)
	err = qx.Error
	return
}

func (repo *Repository) Register(ctx context.Context, p PayloadUserRegister) (err error) {
	var table dao.User
	qx := repo.db
	tx := qx.Debug().Begin()

	helper.AdjustStructToStruct(p, &table)
	user := tx.Save(&table)
	err = user.Error
	if err != nil {
		log.Error(ctx, err)
		tx.Rollback()
		return
	}

	favourite := dao.Watchlist{
		Name:        "Favourite",
		UserId:      table.ID,
		IsFavourite: true,
	}

	wl := tx.Save(&favourite)
	err = wl.Error
	if err != nil {
		log.Error(ctx, err)
		tx.Rollback()
		return
	}
	tx.Commit()

	return
}
