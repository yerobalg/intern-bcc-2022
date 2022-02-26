package daerah

type DaerahService struct {
	repo DaerahRepository
}

func NewDaerahService(daerahRepo DaerahRepository) DaerahService {
	return DaerahService{repo: daerahRepo}
}

func (r *DaerahService) GetDaerah(
	daerah *[]OutputDaerah,
	id string,
	kolomAsal string,
	kolomTujuan string,
	tabelTujuan string,
) error {
	return r.repo.GetDaerah(daerah, id, kolomAsal, kolomTujuan, tabelTujuan)
}
