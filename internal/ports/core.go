package ports

// interface that our core adapter will implement
type ArithmeticPort interface {
	Addition(a, b int32) (int32, error)
	Subtraction(a, b int32) (int32, error)
	Multiplication(a, b int32) (int32, error)
	Division(a, b int32) (int32, error)
}
