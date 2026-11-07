package terminal

import (
	"fmt"
	"time"

	"github.com/MayaraPerez/Flight-Watch/internal/model"
)

func CapturarDados() (model.DadosEntrada, error) {

	var dados model.DadosEntrada

	fmt.Print("Entre com Destino: ")
	fmt.Scan(&dados.Destino)

	fmt.Print("Entre com Valor: ")
	fmt.Scan(&dados.Valor)

	fmt.Print("Entre com data de ida (DD/MM/AAAA): ")
	fmt.Scan(&dados.DataIda)

	fmt.Print("Entre com data de volta (DD/MM/AAAA): ")
	fmt.Scan(&dados.DataVolta)

	return dados, nil

}

func ExibirSucesso(destino string, ida, volta time.Time, valor float64) {
	fmt.Printf("\nSucesso! Dados cadastrados %s .\nIda: %s | Volta: %s\nValor proposto: $%.2f\n",
		destino, ida.Format("02/01/2006"), volta.Format("02/01/2006"), valor)
}
