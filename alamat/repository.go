package alamat

import (
	"gorm.io/gorm"
)

type AlamatRepository struct {
	Conn *gorm.DB
}

func NewAlamatRepository(Conn *gorm.DB) AlamatRepository {
	return AlamatRepository{Conn}
}

func (r *AlamatRepository) Save(alamat *Alamat) error {
	return r.Conn.Create(&alamat).Error
}

func (r *AlamatRepository) Update(alamat *Alamat) error {
	return r.Conn.Save(&alamat).Error
}

func (r *AlamatRepository) GetById(id uint64) (alamat *Alamat, err error) {
	alamatObj := &Alamat{}

	result := r.Conn.
		Preload("Kabupaten").
		Preload("Kabupaten.Provinsi").
		Where("id = ?", id).
		First(&alamatObj)
	return alamatObj, result.Error
}

func (r *AlamatRepository) Delete(alamat *Alamat) (err error) {
	return r.Conn.Delete(alamat).Error
}

func (r *AlamatRepository) GetAllUserAddress(idUser uint64) ([]Alamat, error) {
	alamatList := []Alamat{}
	result := r.Conn.
		Preload("Kabupaten").
		Preload("Kabupaten.Provinsi").
		Where("id_user = ?", idUser).
		Find(&alamatList)
	return alamatList, result.Error
}

func (r *AlamatRepository) GetAdminAddress() (Alamat, error) {
	alamat := Alamat{}
	result := r.Conn.
		Preload("Kabupaten").
		Preload("Kabupaten.Provinsi").
		Where("is_user = ?", false).
		First(&alamat)
	return alamat, result.Error
}
