package interfaces

type Operable[T any] interface {
	Add(others ...T) T
	Subtract(others ...T) T
	Multiply(others ...T) T
	Divide(others ...T) T
}