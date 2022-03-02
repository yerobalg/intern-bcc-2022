package role

type Roles struct {
	ID 	 uint64   `json:"id" gorm:"primary_key"`
	Nama string `json:"name" gorm:"type:varchar(100);not null"`
}
