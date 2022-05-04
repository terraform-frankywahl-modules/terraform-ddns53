package ddns_test

import (
	"fmt"
)

// FindBy will find an element in the list
// If nothing is found, it will return -1 as the index along
// with an error (this is to differentiate it from its 0 value)
func FindBy[T any](list []T, f func(s T) bool) (int, T, error) {
	for i, s := range list {
		if f(s) == true {
			return i, s, nil
		}
	}
	var x T
	return -1, x, fmt.Errorf("could not select in the list")
}

// Includes lets you know if an element is included in a list
func Includes[T comparable](list []T, v T) bool {
	i, _, _ := FindBy(list, func(x T) bool { return x == v })
	return i >= 0
}

// Map gives a new slice after having applied a function
// to each element of the original slice
func Map[T any, U any](list []T, f func(s T) U) []U {
	newList := make([]U, len(list))

	for i, s := range list {
		newList[i] = f(s)
	}
	return newList
}
