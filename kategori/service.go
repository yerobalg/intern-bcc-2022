package kategori

type KategoriService struct {
	repo KategoriRepository
}

func NewKategoriService(kategoriRepo KategoriRepository) KategoriService {
	return KategoriService{repo: kategoriRepo}
}

func (s *KategoriService) Save(model *Kategori) error {
	return s.repo.Save(model)
}

func (s *KategoriService) Update(model *Kategori) error {
	return s.repo.Update(model)
}

func (s *KategoriService) Delete(model *Kategori) error {
	return s.repo.Delete(model)
}

func (s *KategoriService) GetById(id uint64) (Kategori, error) {
	return s.repo.GetById(id)
}

func (s *KategoriService) GetSemuaKategori(isTag bool) ([]Kategori, error) {
	return s.repo.GetSemuaKategori(isTag)
}
