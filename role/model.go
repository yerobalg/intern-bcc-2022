package role

type Roles struct {
	ID 	 uint64   `json:"id" gorm:"primary_key;auto_increment:true"`
	Nama string `json:"name" gorm:"type:varchar(100);not null"`
}
