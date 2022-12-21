package dao

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Watchlist struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	Note      string     `json:"note" gorm:"type:text"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedBy string     `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedBy string     `json:"deleted_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (Watchlist) TableName() string {
	return "watchlist"
}

func MigrateWatchlist(db *gorm.DB) {
	if !db.Migrator().HasTable(&Watchlist{}) {
		log.Println("migrate table watchlist")
		db.AutoMigrate(&Watchlist{})
	}
}
