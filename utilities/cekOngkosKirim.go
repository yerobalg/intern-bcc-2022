package utilities

import (
	"fmt"
	"github.com/Bhinneka/go-rajaongkir"
	"os"
	"time"
)

func CekOngkosKirim(idKabUser string, idKabSeller string, berat int) ro.Cost {
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

	return cost
}
