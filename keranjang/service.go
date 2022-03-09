package keranjang

type KeranjangService struct {
	repo KeranjangRepository
}

func NewKeranjangService(keranjangRepo KeranjangRepository) KeranjangService {
	return KeranjangService{repo: keranjangRepo}
}

func (s *KeranjangService) AddKeranjang(keranjang *Keranjang) error {
	return s.repo.AddKeranjang(keranjang)
}

func (s *KeranjangService) GetKeranjangUser(idUser uint64) ([]Keranjang, error) {
	return s.repo.GetKeranjangUser(idUser)
}

func (s *KeranjangService) GetKeranjangBatch(idBatch []uint64) ([]Keranjang, error) {
	return s.repo.GetKeranjangBatch(idBatch)
}

func (s *KeranjangService) GetSemuaMetodeBayar() ([]Metode_Pembayaran, error) {
	return s.repo.GetSemuaMetodeBayar()
}

func (s *KeranjangService) GetMetodeByID(id uint64) (Metode_Pembayaran, error) {
	return s.repo.GetMetodeByID(id)
}
