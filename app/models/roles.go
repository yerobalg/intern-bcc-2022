package models

type Roles struct {
	ID 	 uint   `json:"id" gorm:"primary_key"`
	Nama string `json:"name" gorm:"type:varchar(100);not null"`
}
