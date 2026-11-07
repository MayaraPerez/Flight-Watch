package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/MayaraPerez/Flight-Watch/internal/pricing"
	"github.com/MayaraPerez/Flight-Watch/internal/terminal"
	"github.com/MayaraPerez/Flight-Watch/internal/validation"
)

func main() {
	dados, err := terminal.CapturarDados()
	if err != nil {
		fmt.Println("Erro ao ler dados do terminal:", err)
		return
	}

	//Validação
	inputValor, err := validation.ValidarValor(dados.Valor)
	if err != nil {
		fmt.Println("Não é um valor válido:", err)
		return
	}

	ida, volta, err := validation.ValidarData(dados.DataIda, dados.DataVolta)
	if err != nil {
		fmt.Println("Erro na validação das datas:", err)
		return
	}

	terminal.ExibirSucesso(dados.Destino, ida, volta, inputValor)

	fmt.Println("----------------------------------------")
	fmt.Println("Verificando Oportunidades banco de dados...")
	time.Sleep(8 * time.Second)

	fmt.Println("")

	preco, err := pricing.AnalisePrecos(inputValor)
	if err != nil {
		fmt.Println("Preços acima do limite")
		return
	} else {
		fmt.Println("Temos uma Oportunidade")
	}

	var input int64
	fmt.Println("Digite 1 para ver: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println("Temos essas oferta(s): ", preco)
	case 2:
		fmt.Println("Proposta enviada por email")
	default:
		fmt.Println("Numero não foi legivel")
	}

	fmt.Println("----------------------------------------")

	var escolha string

	fmt.Print("Você interesse em alguma delas?\nSe sim entre com preço: ")
	fmt.Scan(&escolha)

	converEscolha, err := strconv.ParseFloat(escolha, 64)
	if err != nil {
		fmt.Println("Erro ao converter:", err)
		return
	}

	if pricing.ExistePreco(converEscolha, preco) {
		fmt.Println("Preço escolhido:", converEscolha)
	} else {
		fmt.Println("Erro: Esse preço não estava na lista de ofertas válidas para você.")
	}

}
