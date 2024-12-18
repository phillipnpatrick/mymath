package datastructures

// Define a generic Set type
type Set[T comparable] struct {
	data map[T]struct{} // Use struct{} as the value type for memory efficiency
}

// Create a new Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

// Add an element to the set
func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

// Remove an element from the set
func (s *Set[T]) Remove(value T) {
	delete(s.data, value)
}

// Check if the set contains an element
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

// Get the size of the set
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Clear the set
func (s *Set[T]) Clear() {
	for k := range s.data {
		delete(s.data, k)
	}
}

// Get all elements in the set as a slice
func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.data))
	for key := range s.data {
		elements = append(elements, key)
	}
	return elements
}

// // Main function to demonstrate the usage
// func main() {
// 	// Create a set for integers
// 	intSet := NewSet[int]()
// 	intSet.Add(1)
// 	intSet.Add(2)
// 	intSet.Add(3)

// 	fmt.Println("Integer Set Elements:", intSet.Elements()) // [1 2 3]
// 	fmt.Println("Integer Set Size:", intSet.Size())         // 3
// 	fmt.Println("Contains 2?", intSet.Contains(2))          // true

// 	intSet.Remove(2)
// 	fmt.Println("Contains 2 after removal?", intSet.Contains(2)) // false

// 	// Create a set for strings
// 	stringSet := NewSet[string]()
// 	stringSet.Add("apple")
// 	stringSet.Add("banana")
// 	stringSet.Add("cherry")

// 	fmt.Println("String Set Elements:", stringSet.Elements()) // [apple banana cherry]
// 	fmt.Println("String Set Size:", stringSet.Size())         // 3
// 	stringSet.Clear()
// 	fmt.Println("String Set Size after clear:", stringSet.Size()) // 0
// }