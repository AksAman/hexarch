package arithmetic

import "errors"

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith *Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith *Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith *Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith *Adapter) Division(a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
