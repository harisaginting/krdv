package dao

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type WatchlistMovie struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	MovieId   int        `json:"movie_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedBy string     `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedBy string     `json:"deleted_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (WatchlistMovie) TableName() string {
	return "watchlist_movie"
}

func MigrateWatchlistMovie(db *gorm.DB) {
	if !db.Migrator().HasTable(&WatchlistMovie{}) {
		log.Println("migrate table watchlist movie")
		db.AutoMigrate(&WatchlistMovie{})
	}
}
