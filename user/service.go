package user

type UserService struct {
	repo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return UserService{repo: userRepo}
}

func (s UserService) Register(model *Users) error {
	return s.repo.Register(model)
}

func (s UserService) Login(
	model *UserLoginInput,
) (*Users, error) {
	return s.repo.Login(model)
}

func (s UserService)GetByID(id uint64) (*Users, error) {
	return s.repo.GetByID(id)
}


