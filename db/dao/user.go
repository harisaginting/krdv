package dao

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Username  string     `json:"username"  gorm:"unique_index"`
	Fullname  string     `json:"fullname"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedBy string     `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedBy string     `json:"deleted_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func MigrateUser(db *gorm.DB) {
	if !db.Migrator().HasTable(&User{}) {
		log.Println("migrate table user")
		db.AutoMigrate(&User{})
	}
}
