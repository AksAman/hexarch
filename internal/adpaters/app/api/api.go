package api

import "github.com/AksAman/hexarch/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
}

// constructor
func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apia *Adapter) GetAddition(a int32, b int32) (int32, error) {
	res, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (apia *Adapter) GetSubtraction(a int32, b int32) (int32, error) {
	res, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (apia *Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	res, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (apia *Adapter) GetDivision(a int32, b int32) (int32, error) {
	res, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	return res, nil
}
