package dao

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Watchlist struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	UserId      int        `json:"user_id"`
	Name        string     `json:"name" gorm:"type:varchar(255)"`
	Note        string     `json:"note" gorm:"type:text"`
	IsFavourite bool       `json:"is_favourite" gorm:"default:false"`
	CreatedBy   string     `json:"created_by,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedBy   string     `json:"updated_by,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedBy   string     `json:"deleted_by,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
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
