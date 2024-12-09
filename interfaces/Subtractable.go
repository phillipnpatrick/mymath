package interfaces

// Subtractable is a generic interface that requires an Subtract method
type Subtractable[T any] interface {
	Subtract(others ...T) T
}