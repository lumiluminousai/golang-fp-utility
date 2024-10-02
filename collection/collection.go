package collection

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

// Package utility provides utility functions for functional programming in Go.
//
// This file is part of golang-fp-utility.
//
// golang-fp-utility is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// golang-fp-utility is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with golang-fp-utility. If not, see <https://www.gnu.org/licenses/lgpl-3.0.txt>.

// Map applies a transformation function to each item in the list and returns a new list.
func Map[T1 any, T2 any](source []T1, transform func(item T1) T2) []T2 {
	result := []T2{}
	for _, item := range source {
		result = append(result, transform(item))
	}
	return result
}

// FilterMap filters a hashmap based on a provided function.
func FilterMap[K comparable, V any](source map[K]V, filteringFunc func(key K, value V) bool) map[K]V {
	result := make(map[K]V)
	for key, value := range source {
		if filteringFunc(key, value) {
			result[key] = value
		}
	}
	return result
}

// FlatMap flattens a list of lists into a single list.
func FlatMap[T1 any](source [][]T1) []T1 {
	result := []T1{}
	for _, item := range source {
		result = append(result, item...)
	}
	return result
}

// Reduce reduces a list to a single value using the provided function.
func Reduce[T any](source []T, reduceFunc func(acc T, item T) T, initialValue T) T {
	acc := initialValue
	for _, item := range source {
		acc = reduceFunc(acc, item)
	}
	return acc
}

// Summable includes all types that can be summed, such as integers and floats.
type Summable interface {
	int | int32 | int64 | float32 | float64
}

// Sum returns the sum of elements in a slice of summable types.
func Sum[T Summable](list []T) T {
	var total T
	for _, v := range list {
		total += v
	}
	return total
}

// CloneMap creates a shallow copy of the given map.
func CloneMap[K comparable, V any](source map[K]V) map[K]V {
	clone := make(map[K]V, len(source))
	for key, value := range source {
		clone[key] = value
	}
	return clone
}

// CloneList creates a shallow copy of the given list.
func CloneList[T any](source []T) []T {
	clone := make([]T, len(source))
	copy(clone, source)
	return clone
}

// Sort sorts a slice using a custom less function.
func Sort[T any](list []T, less func(i, j int) bool) []T {
	sort.Slice(list, less)
	return list
}

// Distinct returns a slice containing only unique elements.
func Distinct[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	unique := []T{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

// DistinctFunc returns a slice containing unique elements using a custom comparison function.
func DistinctFunc[T comparable](slice []T, compareFunc func(a, b T) bool) []T {
	seen := make(map[T]bool)
	unique := []T{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

// ForEach executes a function for each item in the list.
func ForEach[T any](source []T, action func(item T)) {
	for _, item := range source {
		action(item)
	}
}

// ForEachWithError executes a function for each item and handles errors.
func ForEachWithError[T any](source []T, action func(item T) error) error {
	for _, item := range source {
		if err := action(item); err != nil {
			return err
		}
	}
	return nil
}

// MapReturnWithError applies a transformation function to each item and handles errors.
func MapReturnWithError[T1 any, T2 any](source []T1, mappingFunc func(item T1) (T2, error)) ([]T2, error) {
	result := []T2{}

	for idx, item := range source {
		res, err := mappingFunc(item)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error mapping at index:'%v', error", idx))
		}
		result = append(result, res)
	}
	return result, nil
}

// Filter returns a filtered list based on the provided function.
func Filter[T any](source []T, filterFunc func(item T) bool) []T {
	result := []T{}
	for _, item := range source {
		if filterFunc(item) {
			result = append(result, item)
		}
	}
	return result
}

// Exists checks if any element in the collection satisfies the condition.
// T is a generic type parameter that can represent any type.
func Exists[T any](collection []T, condition func(T) bool) bool {
	for _, item := range collection {
		if condition(item) {
			return true
		}
	}
	return false
}

// Generic function to find the highest value
func Max[T constraints.Ordered](slice []T) (max T, found bool) {
	if len(slice) == 0 {
		return max, false // If empty, return default value and found = false
	}

	max = slice[0] // Set the first value as initial max
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max, true // Return max and found = true
}

// Generic function to find the lowest value
func Min[T constraints.Ordered](slice []T) (min T, found bool) {
	if len(slice) == 0 {
		return min, false // If empty, return default value and found = false
	}

	min = slice[0] // Set the first value as initial min
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min, true // Return min and found = true
}
