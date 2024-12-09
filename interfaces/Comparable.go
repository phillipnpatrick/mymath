package interfaces

// Comparable is a generic interface that requires:
// Compare
// Equals
// GreaterThan
// GreaterThanOrEqualTo
// LessThan
// LessThanOrEqualTo
type Comparable[T any] interface {
	// Compare returns -1 if less than, 0 if equal, 1 if greater than
	Compare(other T) int

	// Equals determines if items have the same value
	Equals(other T) bool

	GreaterThan(other T) bool
	
	GreaterThanOrEqualTo(other T) bool
	
	LessThan(other T) bool
	
	LessThanOrEqualTo(other T) bool
}
