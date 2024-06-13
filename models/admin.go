package models

import "gorm.io/gorm"

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

func AutoMigrateAdminTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&Admin{}); err != nil {
		return err
	}
	return nil
}
