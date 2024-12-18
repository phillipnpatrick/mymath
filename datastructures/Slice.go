package datastructures

import "fmt"

// SliceContains checks if a slice contains a specific element
func SliceContains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func SliceRemoveAtIndex[T any](slice []T, index int) ([]T, error) {
	if index < 0 || index >= len(slice) {
		return nil, fmt.Errorf("index out of range")
	}
	return append(slice[:index], slice[index+1:]...), nil
}
