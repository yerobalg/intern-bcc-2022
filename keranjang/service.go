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

func (s *KeranjangService) GetKeranjangUser(idUser uint64) (Keranjang, error) {
	return s.repo.GetKeranjangUser(idUser)
}
