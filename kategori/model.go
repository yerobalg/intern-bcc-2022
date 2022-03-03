package kategori

type Kategori struct {
	ID   uint64 `json:"id" gorm:"primary_key;autoIncrement:true;not null;type:bigint"`
	Nama string `json:"nama" gorm:"type:varchar(100);not null"`
}

type KategoriInput struct {
	Nama string `json:"nama" binding:"required"`
}

func (Kategori) TableName() string {
	return "kategori"
}
