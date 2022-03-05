package daerah

type Provinsi struct {
	IDProv string `json:"idProvinsi" gorm:"primary_key;type:varchar(2);column:id_prov;not null"`
	Nama   string `json:"nama" gorm:"type:varchar(35);column:nama;not null"`
}

func (Provinsi) TableName() string {
	return "provinsi"
}

type Kabupaten struct {
	IDKab    string   `gorm:"primary_key;type:varchar(3);column:id_kab;not null"`
	Nama     string   `gorm:"type:varchar(45);column:nama;not null"`
	Tipe     string   `gorm:"type:varchar(20);column:tipe;not null"`
	KodePos  string   `gorm:"type:varchar(5);column:kode_pos;not null"`
	IDProv   string   `gorm:"type:varchar(2);column:id_prov;not null"`
	Provinsi Provinsi `gorm:"foreignKey:IDProv;references:IDProv"`
}

func (Kabupaten) TableName() string {
	return "kabupaten"
}
