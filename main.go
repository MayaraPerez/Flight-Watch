package main

import (
	"fmt"
	"time"
)

func ValidarValor(v float64) (float64, error) {
	if v <= 10 {
		return 0, fmt.Errorf(":erro só aceita valores acima de $10")
	}
	return v, nil
}

func ValidarData(ida, volta string) (time.Time, time.Time, error) {
	layout := "02/01/2006"
	hoje := time.Now().Truncate(24 * time.Hour)

	dataIda, err := time.Parse(layout, ida)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("Use o formato DD/MM/AAAA")
	}

	dataVolta, err := time.Parse(layout, volta)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("Use o formato DD/MM/AAAA")
	}
	// Agora a comparação compara 06/11 com 05/11
	if dataIda.Before(hoje) {
		return time.Time{}, time.Time{}, fmt.Errorf("Erro: A data de ida não pode ser anterior a hoje!")
	}

	// Verifica se a data de ida é igual a data de volta
	if dataIda.Equal(dataVolta) {
		return time.Time{}, time.Time{}, fmt.Errorf("Erro: A data de ida não pode ser igual à data de volta!")
	}

	// validação para evitar volta antes da ida!
	if dataVolta.Before(dataIda) {
		return time.Time{}, time.Time{}, fmt.Errorf("Erro: A data de volta não pode ser antes da data de ida!")
	}

	return dataIda, dataVolta, nil
}

func AnalisePrecos(valorUsuario float64) (float64, error) {
	listaPrecos := []float64{
		4200,
		3800,
		3250,
		2900,
		1000,
	}

	for _, precoDisponivel := range listaPrecos {
		if precoDisponivel <= valorUsuario {
			return precoDisponivel, nil
		}
	}
	return 0, fmt.Errorf("nenhum preço abaixo do limite encontrado")
}

func main() {
	var destino string
	var valor float64
	var dataIda, dataVolta string

	fmt.Print("Entre com Destino: ")
	fmt.Scan(&destino)

	fmt.Print("Entre com Valor: ")
	fmt.Scan(&valor)

	valor, err := ValidarValor(valor)
	if err != nil {
		fmt.Println("não é um valor válido", err)
		return
	}

	fmt.Print("Entre com data de ida (DD/MM/AAAA): ")
	fmt.Scan(&dataIda)

	fmt.Print("Entre com data de volta (DD/MM/AAAA): ")
	fmt.Scan(&dataVolta)

	ida, volta, err := ValidarData(dataIda, dataVolta)
	if err != nil {
		fmt.Println("Erro na validação das datas:", err)
		return
	}

	fmt.Printf("\nSucesso! Dados cadastrados %s .\nIda: %s | Volta: %s\nValor proposto: $%.2f\n",
		destino, ida.Format("02/01/2006"), volta.Format("02/01/2006"), valor)

	fmt.Println("----------------------------------------")
	fmt.Println("Verificando Oportunidades banco de dados...")
	time.Sleep(8 * time.Second)

	fmt.Println("")

	preco, err := AnalisePrecos(valor)
	if err != nil {
		fmt.Println("Preços acima do limite")
		return
	} else {
		fmt.Println("Temos uma Oportunidade")
	}

	var input int64
	fmt.Println("Digite 1 caso tenha interesse ou 2 caso não tem interesse: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println("Achamos uma viagem no valor", preco)
	case 2:
		fmt.Println("Proposta enviada por email")
	default:
		fmt.Println("Numero não foi legivel")
	}

	fmt.Println("----------------------------------------")

}
