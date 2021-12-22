package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

// Req

func NewUser(db *gorm.DB) UserRepositoryDB {
	db.AutoMigrate(&Users{}, &Roles{})
	return UserRepositoryDB{db: db}
}

func (r UserRepositoryDB) GetUses() ([]Users, error) {
	users := []Users{}
	r.db.Find(&users)
	return users, nil
}

func (r UserRepositoryDB) AddUser(u Users) (*Users, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	user := Users{
		UserName: u.UserName,
		UserUUID: id,
		Password: "test",
		Roles:    1,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepositoryDB) AddRole(roleName string) error {
	role := Roles{
		Name: roleName,
	}
	r.db.Create(&role)
	return nil
}
