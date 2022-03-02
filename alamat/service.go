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

func (s AlamatService) Update(model *Alamat) error {
	return s.repo.Update(model)
}

func (s AlamatService) GetById(id uint64) (alamat *Alamat, err error) {
	return s.repo.GetById(id)
}

func (s AlamatService) Delete(alamat *Alamat) error {
	return s.repo.Delete(alamat)
}
