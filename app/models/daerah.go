package models

type OutputDaerah struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

type JenisDaerah struct {
	IDJenis uint   `gorm:"primary_key;autoIncrement:false;column:id_jenis;not null"`
	Nama    string `gorm:"type:varchar(9);column:nama;not null"`
}

func (JenisDaerah) TableName() string {
	return "jenis_daerah"
}

type Provinsi struct {
	IDProv    string      `gorm:"primary_key;type:varchar(2);column:id_prov;not null"`
	Nama      string      `gorm:"type:varchar(35);column:nama;not null"`
	Kabupaten []Kabupaten `gorm:"foreignkey:IDProv; association_foreignkey:IDProv"`
}

func (Provinsi) TableName() string {
	return "provinsi"
}

type Kabupaten struct {
	IDKab       string      `gorm:"primary_key;type:varchar(4);column:id_kab;not null"`
	Nama        string      `gorm:"type:varchar(35);column:nama;not null"`
	IDProv      string      `gorm:"type:varchar(2);column:id_prov;not null"`
	IDJenis     uint        `gorm:"column:id_jenis;not null"`
	Provinsi    Provinsi    `gorm:"foreignKey:IDProv;references:IDProv"`
	JenisDaerah JenisDaerah `gorm:"foreignKey:IDJenis;references:IDJenis"`
	Kecamatan   []Kecamatan `gorm:"foreignkey:IDKab;association_foreignkey:IDKab"`
}

func (Kabupaten) TableName() string {
	return "kabupaten"
}

type Kecamatan struct {
	IDKec     string      `gorm:"primary_key;type:varchar(6);column:id_kec;not null"`
	Nama      string      `gorm:"type:varchar(32);column:nama;not null"`
	IDKab     string      `gorm:"type:varchar(4);column:id_kab;not null"`
	Kabupaten Kabupaten   `gorm:"foreignKey:IDKab;references:IDKab"`
	Kelurahan []Kelurahan `gorm:"foreignkey:IDKec;association_foreignkey:IDKec"`
}

func (Kecamatan) TableName() string {
	return "kecamatan"
}

type Kelurahan struct {
	IDKel       string      `gorm:"primary_key;type:varchar(10);column:id_kel;not null"`
	Nama        string      `gorm:"type:varchar(32);column:nama"`
	IDKec       string      `gorm:"type:varchar(6);column:id_kec"`
	IDJenis     uint        `gorm:"column:id_jenis;not null"`
	Kecamatan   Kecamatan   `gorm:"foreignKey:IDKec;references:IDKec"`
	JenisDaerah JenisDaerah `gorm:"foreignKey:IDJenis;references:IDJenis"`
}

func (Kelurahan) TableName() string {
	return "kelurahan"
}
