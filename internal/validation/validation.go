package validation

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
