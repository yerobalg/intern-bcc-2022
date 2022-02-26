package services

import (
	"clean-arch-2/app/models"
	"clean-arch-2/app/repositories"
)

type DaerahService struct {
	repo repositories.DaerahRepository
}

func NewDaerahService(daerahRepo repositories.DaerahRepository) DaerahService {
	return DaerahService{repo: daerahRepo}
}

func (r *DaerahService) GetDaerah(
	daerah *[]models.OutputDaerah,
	id string,
	kolomAsal string,
	kolomTujuan string,
	tabelTujuan string,
) error {
	return r.repo.GetDaerah(daerah, id, kolomAsal, kolomTujuan, tabelTujuan)
}
