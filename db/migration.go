package db

import (
	"log"

	"github.com/harisaginting/krdv/db/dao"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	log.Println("migration db", db.Migrator().CurrentDatabase())
	dao.MigrateUser(db)
	dao.MigrateMovie(db)
	dao.MigrateWatchlist(db)
	dao.MigrateWatchlistMovie(db)
	log.Println("migration success")
}
