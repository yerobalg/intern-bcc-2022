package keranjang

type KeranjangFormatter struct {
	Seller string `json:"seller"`
	Produk string `json:"produk"`
	JenisPengiriman string `json:"jenis_pengiriman"`
}

type Seller struct {
	ID   uint64 `json:"id"`
	Nama string `json:"nama"`
}

type Produk struct {
	ID    uint64 `json:"id"`
	Nama  string `json:"nama"`
	Harga uint64 `json:"harga"`
	Stok  uint   `json:"stok"`
}

type JenisPengiriman struct {
	Kurir    string `json:"kurir" gorm:"type:varchar(100);not null"`
	Service []Service
}

type Service struct {
	Nama string `json:"nama"`
	Harga uint64 `json:"harga"`
	Estimasi string `json:"estimasi"`
}
