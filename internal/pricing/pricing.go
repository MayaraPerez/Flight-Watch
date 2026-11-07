package pricing

import "fmt"

func AnalisePrecos(inputUsuario float64) ([]float64, error) {
	listaPrecos := []float64{
		4200,
		3800,
		3250,
		2900,
		1000,
	}

	var filtrarPrecos []float64

	for _, preco := range listaPrecos {
		if preco <= inputUsuario {
			filtrarPrecos = append(filtrarPrecos, preco)
		}
	}

	if len(filtrarPrecos) == 0 {
		return nil, fmt.Errorf("nenhum preço abaixo do limite encontrado")
	}

	return filtrarPrecos, nil
}

func ExistePreco(inputEscolhido float64, listaFiltrada []float64) bool {
	for _, preco := range listaFiltrada {
		if preco == inputEscolhido {
			return true
		}
	}
	return false

}
