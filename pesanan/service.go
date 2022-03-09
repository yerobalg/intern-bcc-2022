package pesanan

type PesananService struct {
	repo PesananRepository
}

func NewPesananService(pesananRepo PesananRepository) PesananService {
	return PesananService{repo: pesananRepo}
}

func (s *PesananService) AddPesanan(
	pesanan *Pesanan,
	idKeranjang []uint64,
) error {
	err := s.repo.AddPesanan(pesanan)
	if err != nil {
		return err
	}
	return s.repo.AddKeranjangPesanan(uint64(pesanan.ID), idKeranjang)
}

func (s *PesananService) GetPesanan(idPesanan uint64) (Pesanan, error) {
	return s.repo.GetPesanan(idPesanan)
}

func (s *PesananService) DeletePesanan(pesanan *Pesanan) error {
	err := s.repo.DeletePesanan(pesanan)
	if err != nil {
		return err
	}

	return s.repo.DeleteKeranjangPesanan(uint64(pesanan.ID))
}
