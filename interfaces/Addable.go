package interfaces

// Addable is a generic interface that requires an Add method
type Addable[T any] interface {
	Add(others ...T) T
}