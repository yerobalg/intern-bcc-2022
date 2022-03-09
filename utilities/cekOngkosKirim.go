package utilities

import (
	"fmt"
	"github.com/Bhinneka/go-rajaongkir"
	"os"
	"time"
)

type OngkirFormat struct {
	Kode    string    `json:"kode"`
	Nama    string    `json:"nama"`
	Layanan []Layanan `json:"layanan"`
}

type Layanan struct {
	Kode          string `json:"kode"`
	Nama          string `json:"nama"`
	Harga         uint64 `json:"harga"`
	EstimasiKirim string `json:"estimasiKirim"`
}

func CekOngkosKirim(idKabUser string, idKabSeller string, berat int) []OngkirFormat {
	raja := ro.New(os.Getenv("RAJA_ONGKIR_API_KEY"), 10*time.Second)

	q := ro.QueryRequest{
		Origin:      idKabSeller,
		Destination: idKabUser,
		Weight:      berat,
		Courier:     "jne",
	}

	result := raja.GetCost(q)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	cost, ok := result.Result.(ro.Cost)
	if !ok {
		fmt.Println("Result is not Cost")
	}

	var ongkir []OngkirFormat

	for _, res := range cost.Providers {
		var layanan []Layanan

		for _, prov := range res.Costs {
			layanan = append(layanan, Layanan{
				Kode:          prov.Service,
				Nama:          prov.Description,
				Harga:         uint64(prov.Cost[0].Value),
				EstimasiKirim: prov.Cost[0].EstimatedDay,
			})
		}

		ongkir = append(ongkir, OngkirFormat{
			Nama:    res.Name,
			Kode:    res.Code,
			Layanan: layanan,
		})
	}

	return ongkir
}
