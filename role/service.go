package role

type RoleService struct {
	repo RoleRepository
}

func NewRoleService(roleRepo RoleRepository) RoleService {
	return RoleService{repo: roleRepo}
}

func (r RoleService) GetRoleById(id uint64) (*Roles, error) {
	return r.repo.GetRoleById(id)
}