package dao

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID           int        `json:"id,omitempty" gorm:"primaryKey"`
	IDExternal   int        `json:"id_external"`
	Title        string     `json:"title" gorm:"type:varchar(255)"`
	BackdropPath string     `json:"backdrop_path,omitempty" gorm:"type:varchar(255)"`
	Adult        int        `json:"adult"`
	Overview     string     `json:"overview"`
	CreatedBy    string     `json:"created_by,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedBy    string     `json:"updated_by,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedBy    string     `json:"deleted_by,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

func (Movie) TableName() string {
	return "movie"
}

func MigrateMovie(db *gorm.DB) {
	if !db.Migrator().HasTable(&Movie{}) {
		log.Println("migrate table movie")
		db.AutoMigrate(&Movie{})
	}
}
