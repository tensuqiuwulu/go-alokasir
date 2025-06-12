package main

import (
	"gorm.io/gorm"
)

type Users struct{
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsActive bool `gorm:"default:true"`
}

type Migration20250612145010CreateUsersTable struct {}

func (m *Migration20250612145010CreateUsersTable) Up(db *gorm.DB) error {
	// Implement your migration here
	return db.AutoMigrate(&Users{})
}

func (m *Migration20250612145010CreateUsersTable) Down(db *gorm.DB) error {
	// Implement your rollback here
	return db.Migrator().DropTable(&Users{})
}

// Ekspor struct migrasi untuk sistem plugin
var Migration20250612145010CreateUsersTable_Exported = &Migration20250612145010CreateUsersTable{}
