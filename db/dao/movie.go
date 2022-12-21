package dao

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID            int        `json:"id" gorm:"primaryKey"`
	IDExternal    int        `json:"id_external"`
	Featured      int        `json:"featured"`
	Description   string     `json:"description"`
	Revenue       string     `json:"revenue"`
	Public        int        `json:"public"`
	Name          string     `json:"name"`
	SortBy        int        `json:"sort_by"`
	BackdropPath  string     `json:"backdrop_path,omitempty"`
	Runtime       int        `json:"runtime"`
	AverageRating float64    `json:"average_rating"`
	Iso31661      string     `json:"iso_3166_1"`
	Adult         int        `json:"adult"`
	PosterPath    string     `json:"poster_path,omitempty"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedBy     string     `json:"updated_by"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedBy     string     `json:"deleted_by,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
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
