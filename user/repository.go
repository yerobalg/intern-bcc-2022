package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) UserRepository {
	return UserRepository{Conn}
}

func (r *UserRepository) Register(user *Users) error {
	return r.Conn.Create(&user).Error
}

func (r *UserRepository) Login(
	userLogin *UserLoginInput,
) (*Users, error) {
	user := &Users{}

	result := r.Conn.
		Where("email = ?", userLogin.UsernameOrEmail).
		Or("username = ?", userLogin.UsernameOrEmail).
		First(&user)
	return user, result.Error
}