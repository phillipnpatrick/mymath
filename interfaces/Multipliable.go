package interfaces

// Multipliable is a generic interface that requires an Multiply method
type Multipliable[T any] interface {
	Multiply(others ...T) T
}