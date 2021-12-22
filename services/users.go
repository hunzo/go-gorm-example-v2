package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserUUID uuid.UUID `gorm:"uniqueIndex"`
	UserName string    `json:"username" gorm:"uniqueIndex"`
	Password string    `json:"password"`
	Roles    int       `gorm:"foreignKey:RoleId"`
}
type Roles struct {
	Name   string `json:"role_name" gorm:"uniqueIndex"`
	RoleId int    `json:"role_id" gorm:"primaryKey;autoIncrement"`
}

type UserRepo interface {
	GetUsers() ([]Users, error)
	AddUser(Users) (*Users, error)
	AddRole(string) error
}
