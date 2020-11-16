package vo

import (
	"errors"
	"math"
)

var (
	ErrAmountValueInvalid = errors.New("Invalid amount value.")
)

type Amount struct {
	value float64
}

func NewAmount(value float64) (*Amount, error) {
	if value < 0 {
		return nil, ErrAmountValueInvalid
	}

	rounded := math.Round(value*100) / 100

	return &Amount{rounded}, nil
}

func (a Amount) Value() float64 {
	return a.value
}
