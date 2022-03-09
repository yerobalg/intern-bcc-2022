package config

import (
	"clean-arch-2/alamat"
	"clean-arch-2/user"
	"clean-arch-2/kategori"
	"clean-arch-2/produk"
	"clean-arch-2/keranjang"
	"clean-arch-2/pesanan"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDB() (*gorm.DB, error) {
	Init(".database")
	db, error := gorm.Open(
		postgres.Open(fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DBNAME"),
		)),
		&gorm.Config{},
	)

	if error != nil {
		return nil, error
	}
	fmt.Println("Successfully connected to database!")

	db.AutoMigrate(
		// &models.Roles{},
		&user.Users{},
		&alamat.Alamat{},
		&kategori.Kategori{},	
		&produk.Produk{},
		&produk.Kategori_Produk{},
		&produk.Gambar_Produk{},
		&keranjang.Keranjang{},
		&pesanan.Pesanan{},
		&pesanan.Keranjang_Pesanan{},
	)
	return db, error
}
