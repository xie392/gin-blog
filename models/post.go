package models

import "gorm.io/gorm"

type Post struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	AuthorID   int    `json:"author_id" binding:"required"`
	CreatedAt  string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  string `json:"updated_at" gorm:"autoUpdateTime"`
	CoverImage string `json:"cover_image"`
}

func AutoMigratePostTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&Post{}); err != nil {
		return err
	}
	return nil
}
