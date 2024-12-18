package datastructures

// Finds keys that are in map1 but not map2.
func Difference[K comparable, V any](map1, map2 map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range map1 {
		if _, exists := map2[k]; !exists {
			result[k] = v
		}
	}
	return result
}

// Filters the map based on a condition provided as a function.
func Filter[K comparable, V any](m map[K]V, condition func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if condition(k, v) {
			result[k] = v
		}
	}
	return result
}

// Intersection function for generic maps
func Intersection[K comparable, V any](map1, map2 map[K]V) map[K]V {
	result := make(map[K]V)

	// Iterate over the first map and check if the key exists in the second map
	for k := range map1 {
		if v2, exists := map2[k]; exists {
			// Add the key to the result if it's in both maps
			result[k] = v2
		}
	}

	return result
}

// Inverts a map by swapping keys and values. Useful for reversing lookups.
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K)
	for k, v := range m {
		result[v] = k
	}
	return result
}

// Extracts all keys from a map as a slice.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Applies a transformation function to all values in the map.
func MapApply[K comparable, V any](m map[K]V, apply func(K, V) V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = apply(k, v)
	}
	return result
}

// Checks if two maps are equal (keys and values).
func MapsEqual[K comparable, V comparable](map1, map2 map[K]V) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if v2, exists := map2[k]; !exists || v != v2 {
			return false
		}
	}
	return true
}

// Merges two maps, optionally handling conflicts with a custom resolver.
func Merge[K comparable, V any](map1, map2 map[K]V, resolver func(V, V) V) map[K]V {
	result := make(map[K]V)
	for k, v := range map1 {
		result[k] = v
	}
	for k, v := range map2 {
		if existing, exists := result[k]; exists {
			result[k] = resolver(existing, v)
		} else {
			result[k] = v
		}
	}
	return result
}

// Finds keys that are in one map or the other but not both.
func SymmetricDifference[K comparable, V any](map1, map2 map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range map1 {
		if _, exists := map2[k]; !exists {
			result[k] = v
		}
	}
	for k, v := range map2 {
		if _, exists := map1[k]; !exists {
			result[k] = v
		}
	}
	return result
}

// Union function for generic maps
func Union[K comparable, V any](map1, map2 map[K]V) map[K]V {
	result := make(map[K]V)

	// Add all elements from the first map
	for k, v := range map1 {
		result[k] = v
	}

	// Add all elements from the second map (overwriting duplicates)
	for k, v := range map2 {
		result[k] = v
	}

	return result
}

// Extracts all values from a map as a slice.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
