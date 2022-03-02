package alamat

type AlamatService struct {
	repo AlamatRepository
}

func NewUserService(userRepo AlamatRepository) AlamatService {
	return AlamatService{repo: userRepo}
}

func (s AlamatService) Save(model *Alamat) error {
	return s.repo.Save(model)
}
