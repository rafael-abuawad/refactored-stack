package models

import (
	"github.com/rafael-abuawad/refactored-stack/config/database"
	"gorm.io/gorm"
)

// User struct defines the user
type User struct {
	gorm.Model
	Name       string
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	ResetToken string
}

// CreateUser create a user entry in the user's table
func CreateUser(user *User) *gorm.DB {
	return database.DB.Create(user)
}

// FindUser searches the user's table with the condition given
func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&User{}).Take(dest, conds...)
}

// FindUserByEmail searches the user's table with the email given
func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUser(dest, "email = ?", email)
}

// FindUserByToken searches the user's table with the reset token given
func FindUserByToken(dest interface{}, token string) *gorm.DB {
	return FindUser(dest, "reset_token = ?", token)
}

// DeleteUser deletes a todo from todos' table with the given todo and user identifier
func DeleteUser(userID interface{}) *gorm.DB {
	return database.DB.Unscoped().Delete(&User{}, "id = ?", userID)
}

// UpdateUser allows to update the todo with the given todoID and userID
func UpdateUser(userID interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&User{}).Where("id = ?", userID).Updates(data)
}
