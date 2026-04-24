package basics

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:64;not null"`
	Email     string    `gorm:"size:128;uniqueIndex;not null"`
	Age       uint8     `gorm:"not null"`
	Status    string    `gorm:"size:16;default:active;index"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:128;not null"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null;index"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
	Tags      []Tag     `gorm:"many2many:post_tags;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Comment struct {
	ID        uint           `gorm:"primaryKey"`
	Content   string         `gorm:"type:text;not null"`
	PostID    uint           `gorm:"not null;index"`
	UserID    uint           `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
}

type Tag struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:64;not null;uniqueIndex"`
	Posts     []Post    `gorm:"many2many:post_tags;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
