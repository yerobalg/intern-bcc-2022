package role

import (
	"gorm.io/gorm"
)

type RoleRepository struct {
	Conn *gorm.DB
}

func NewRoleRepository(Conn *gorm.DB) RoleRepository {
	return RoleRepository{Conn}
}

func (r *RoleRepository) GetRoleById(id uint64) (*Roles, error) {
	role := &Roles{}

	result := r.Conn.
		Where("id = ?", id).
		First(&role)
	return role, result.Error
}