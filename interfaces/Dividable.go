package interfaces

// Dividable is a generic interface that requires an Divide method
type Dividable[T any] interface {
	Divide(others ...T) T
}