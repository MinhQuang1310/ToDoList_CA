// entities/user.go
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID        int64     `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
