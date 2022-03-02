package alamat

type AlamatService struct {
	repo AlamatRepository
}

func NewAlamatService(userRepo AlamatRepository) AlamatService {
	return AlamatService{repo: userRepo}
}

func (s AlamatService) Save(model *Alamat) error {
	return s.repo.Save(model)
}
