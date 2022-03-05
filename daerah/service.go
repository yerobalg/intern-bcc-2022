package daerah

type DaerahService struct {
	repo DaerahRepository
}

func NewDaerahService(daerahRepo DaerahRepository) DaerahService {
	return DaerahService{repo: daerahRepo}
}

func (s *DaerahService) GetSemuaProvinsi(provinsi *[]Provinsi) error {
	return s.repo.GetSemuaProvinsi(provinsi)
}

func (s *DaerahService) GetKabupaten(idProvinsi string) ([]Kabupaten, error) {
	return s.repo.GetKabupaten(idProvinsi)
}