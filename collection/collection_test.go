package collection

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestMap(t *testing.T) {

	t.Run("test use sum() function", func(t *testing.T) {
		// Example list of doubles
		numbers := []float64{1.5, 2.0, 3.5, 4.0}

		// Use utility.Map to double each number in the list
		doubledNumbers := Map(numbers, func(item float64) float64 {
			return item * 2
		})

		// Use utility.Sum to get the summation of the doubled numbers
		sum := Sum(doubledNumbers)

		// Assert the expected values
		assert.Equal(t, []float64{3.0, 4.0, 7.0, 8.0}, doubledNumbers, "Doubled numbers should match the expected list")
		assert.Equal(t, 22.0, sum, "Summation of doubled numbers should be 22.0")
	})

	t.Run("map with nil list", func(t *testing.T) {
		source := []int(nil)

		mappingFunc := func(item int) string {
			return fmt.Sprintf("string_%v", item) // Convert each integer to string with prefix
		}

		result := Map(source, mappingFunc)

		expected := []string{}
		assert.Equal(t, expected, result)
	})

	t.Run("map integers to strings", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		mappingFunc := func(item int) string {
			return fmt.Sprintf("string_%v", item) // Convert each integer to string with prefix
		}

		result := Map(source, mappingFunc)

		expected := []string{"string_1", "string_2", "string_3", "string_4", "string_5"}
		assert.Equal(t, expected, result)
	})

	t.Run("map empty list", func(t *testing.T) {
		source := []int{}
		mappingFunc := func(item int) string {
			return fmt.Sprintf("string_%v", item) // Convert each integer to string with prefix
		}

		result := Map(source, mappingFunc)

		expected := []string{}
		assert.Equal(t, expected, result)
	})
}

func TestFilterMap(t *testing.T) {
	tests := []struct {
		name          string
		source        map[int]string // Example uses int keys and string values for demonstration
		filteringFunc func(int, string) bool
		want          map[int]string
	}{
		{
			name:   "filter out odd keys",
			source: map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
			filteringFunc: func(key int, value string) bool {
				return key%2 == 0 // Keep if key is even
			},
			want: map[int]string{2: "b", 4: "d"},
		},
		{
			name:   "filter out values with length > 1",
			source: map[int]string{1: "a", 2: "bb", 3: "ccc", 4: "dddd"},
			filteringFunc: func(key int, value string) bool {
				return len(value) <= 1 // Keep if value's length is 1 or less
			},
			want: map[int]string{1: "a"},
		},
		{
			name:          "empty map",
			source:        map[int]string{},
			filteringFunc: func(key int, value string) bool { return true },
			want:          map[int]string{},
		},
		{
			name:   "filter everything",
			source: map[int]string{1: "a", 2: "b"},
			filteringFunc: func(key int, value string) bool {
				return false // Filter out everything
			},
			want: map[int]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterMap(tt.source, tt.filteringFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHigherOrderFunction_Reduce(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		reduceFunc := func(acc, value int) int {
			return acc + value
		}

		result := Reduce[int](source, reduceFunc, 0)

		expected := 15
		assert.Equal(t, expected, result)
	})

	t.Run("Success_String", func(t *testing.T) {
		source := []string{"D", "a", "r", "k", " ", "m", "a", "g", "i", "c"}

		reduceFunc := func(acc, value string) string {
			return acc + value
		}

		result := Reduce[string](source, reduceFunc, "")

		expected := "Dark magic"
		assert.Equal(t, expected, result)
	})

	t.Run("Success_Empty_List_Int", func(t *testing.T) {
		source := []int{}

		reduceFunc := func(acc, value int) int {
			return acc + value
		}

		result := Reduce[int](source, reduceFunc, 0)

		expected := 0
		assert.Equal(t, expected, result)
	})

	t.Run("Success_Person", func(t *testing.T) {
		type person struct {
			Name string
			Age  int
		}

		source := []person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		// reduceFunc to accumulate total age and count persons
		reduceFunc := func(acc person, p person) person {
			acc.Age = acc.Age + p.Age
			return acc
		}

		initialAccumulator := person{
			Name: "",
			Age:  0,
		}
		result := Reduce[person](source, reduceFunc, initialAccumulator)

		expected := person{Name: "", Age: 90} // Total age is 90, with 3 persons
		assert.Equal(t, expected, result)
	})

}

func TestHigherOrderFunction_FlatMap(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := [][]int{
			{1, 2, 3},
			{4, 5},
			{6, 7, 8},
		}

		result := FlatMap(source)

		expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
		assert.Equal(t, expected, result)
	})

	t.Run("Success_String", func(t *testing.T) {

		source := [][]string{
			{"a", "b", "c"},
			{"d", "e"},
			{"f", "g", "h"},
		}

		result := FlatMap(source)

		expected := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
		assert.Equal(t, expected, result)
	})
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		slice    interface{}
		expected interface{}
	}{
		{
			name:     "ints",
			slice:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "floats",
			slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: 16.5,
		},
		{
			name:     "empty int slice",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "empty float slice",
			slice:    []float64{},
			expected: 0.0,
		},
		{
			name:     "no elements",
			slice:    []int{},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var sum interface{}
			switch tc.slice.(type) {
			case []int:
				sum = Sum(tc.slice.([]int))
			case []float64:
				sum = Sum(tc.slice.([]float64))
			default:
				t.Fatalf("Unsupported type in test case")
			}

			if !reflect.DeepEqual(sum, tc.expected) {
				t.Errorf("Sum(%v) = %v, want %v", tc.slice, sum, tc.expected)
			}
		})
	}
}

func TestCloneMap(t *testing.T) {
	tests := []struct {
		name   string
		source map[string]int // This test uses string keys and int values for simplicity
		want   map[string]int
	}{
		{
			name:   "non-empty map",
			source: map[string]int{"a": 1, "b": 2},
			want:   map[string]int{"a": 1, "b": 2},
		},
		{
			name:   "empty map",
			source: map[string]int{},
			want:   map[string]int{},
		},
		{
			name:   "nil map",
			source: nil,
			want:   map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CloneMap(tt.source)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestSort_Ints(t *testing.T) {
	intSlice := []int{5, 2, 8, 1, 9}
	expected := []int{1, 2, 5, 8, 9}

	sorted := Sort(intSlice, func(i, j int) bool { return intSlice[i] < intSlice[j] })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for int slice. Got: %v, Expected: %v", intSlice, expected)
	}
}

func TestSort_Strings(t *testing.T) {
	stringSlice := []string{"c", "a", "b"}
	expected := []string{"a", "b", "c"}

	sorted := Sort(stringSlice, func(i, j int) bool { return stringSlice[i] < stringSlice[j] })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string slice. Got: %v, Expected: %v", stringSlice, expected)
	}
}

func TestSort_CustomType(t *testing.T) {
	// Define a custom type for testing
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}

	sorted := Sort(people, func(i, j int) bool { return people[i].Age < people[j].Age })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for custom type slice. Got: %v, Expected: %v", people, expected)
	}
}

func TestSort_StringsByReverseOrder(t *testing.T) {
	stringSlice := []string{"apple", "banana", "cherry"}
	expected := []string{"cherry", "banana", "apple"}

	sorted := Sort(stringSlice, func(i, j int) bool { return stringSlice[i] > stringSlice[j] })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string sort by reverse order. Got: %v, Expected: %v", stringSlice, expected)
	}
}

func TestSort_StringsCaseInsensitive(t *testing.T) {
	stringSlice := []string{"banana", "Apple", "CHERRY"}
	expected := []string{"Apple", "banana", "CHERRY"}

	sorted := Sort(stringSlice, func(i, j int) bool {
		return strings.ToLower(stringSlice[i]) < strings.ToLower(stringSlice[j])
	})

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string sort case insensitive. Got: %v, Expected: %v", stringSlice, expected)
	}
}

func TestSort_StringsByLength(t *testing.T) {
	stringSlice := []string{"ccccc", "aaa", "bbbb"}
	expected := []string{"aaa", "bbbb", "ccccc"}

	sorted := Sort(stringSlice, func(i, j int) bool { return len(stringSlice[i]) < len(stringSlice[j]) })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string sort by length. Got: %v, Expected: %v", stringSlice, expected)
	}
}
func TestSort_StringsByLength_reversed(t *testing.T) {
	stringSlice := []string{"ccccc", "aaa", "bbbb"}
	expected := []string{"ccccc", "bbbb", "aaa"}

	sorted := Sort(stringSlice, func(i, j int) bool { return len(stringSlice[i]) > len(stringSlice[j]) })

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Sort failed for string sort by length. Got: %v, Expected: %v", stringSlice, expected)
	}
}

// TestDistinct tests the Distinct function for various slice types.
func TestDistinct(t *testing.T) {
	tests := []struct {
		name     string
		slice    interface{}
		expected interface{}
	}{
		{
			name:     "ints",
			slice:    []int{1, 2, 3, 2, 4, 5, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "strings",
			slice:    []string{"apple", "banana", "apple", "orange", "banana"},
			expected: []string{"apple", "banana", "orange"},
		},
		{
			name:     "empty int slice",
			slice:    []int{},
			expected: []int{},
		},
		{
			name:     "empty string slice",
			slice:    []string{},
			expected: []string{},
		},
		{
			name:     "no duplicates",
			slice:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var unique interface{}
			switch tc.slice.(type) {
			case []int:
				unique = Distinct(tc.slice.([]int))
			case []string:
				unique = Distinct(tc.slice.([]string))
			default:
				t.Fatalf("Unsupported type in test case")
			}

			if !reflect.DeepEqual(unique, tc.expected) {
				t.Errorf("Distinct(%v) = %v, want %v", tc.slice, unique, tc.expected)
			}
		})
	}
}

func TestDistinctFunc(t *testing.T) {
	tests := []struct {
		name     string
		slice    interface{}
		fn       interface{}
		expected interface{}
	}{
		{
			name:  "ints",
			slice: []int{1, 2, 3, 2, 4, 5, 4, 6},
			fn: func(i, j int) bool {
				return i == j
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "strings",
			slice: []string{"apple", "banana", "apple", "orange", "banana"},
			fn: func(i, j string) bool {
				return i == j
			},
			expected: []string{"apple", "banana", "orange"},
		},
		{
			name:  "empty int slice",
			slice: []int{},
			fn: func(i, j int) bool {
				return i == j
			},
			expected: []int{},
		},
		{
			name:  "empty string slice",
			slice: []string{},
			fn: func(i, j string) bool {
				return i == j
			},
			expected: []string{},
		},
		{
			name:  "no duplicates",
			slice: []int{1, 2, 3},
			fn: func(i, j int) bool {
				return i == j
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var unique interface{}
			switch tc.slice.(type) {
			case []int:
				unique = DistinctFunc(tc.slice.([]int), tc.fn.(func(int, int) bool))
			case []string:
				unique = DistinctFunc(tc.slice.([]string), tc.fn.(func(string, string) bool))
			default:
				t.Fatalf("Unsupported type in test case")
			}

			if !reflect.DeepEqual(unique, tc.expected) {
				t.Errorf("DistinctFunc(%v) = %v, want %v", tc.slice, unique, tc.expected)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	t.Run("filter > 3", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		mappingFunc := func(data int) bool {
			return data > 3 // Convert each integer to string with prefix
		}

		result := Filter(source, mappingFunc)

		expected := []int{4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("filter empty list", func(t *testing.T) {
		source := []int{}
		mappingFunc := func(data int) bool {
			return data > 3 // Convert each integer to string with prefix
		}

		result := Filter(source, mappingFunc)

		expected := []int{}
		assert.Equal(t, expected, result)
	})
}

func TestForEach(t *testing.T) {
	t.Run("print integers", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		forEachFunc := func(item int) {
			fmt.Println(item)
		}

		ForEach(source, forEachFunc)
	})

	t.Run("change value each item", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		forEachFunc := func(item int) {
			item = item * 2
		}

		ForEach(source, forEachFunc)

		expected := []int{1, 2, 3, 4, 5}
		assert.Equal(t, expected, source)
	})

	t.Run("change value object", func(t *testing.T) {

		type TempStruct struct {
			Name  string
			Value int
		}
		value1 := TempStruct{
			Name:  "value1",
			Value: 1,
		}
		value2 := TempStruct{
			Name:  "value2",
			Value: 2,
		}

		source := []TempStruct{value1, value2}
		forEachFunc := func(item TempStruct) {
			item.Value = item.Value * 2
		}

		ForEach(source, forEachFunc)

		expected := []TempStruct{value1, value2}
		assert.Equal(t, expected, source)
	})

	t.Run("change value object pointer", func(t *testing.T) {

		type TempStruct struct {
			Name  string
			Value int
		}
		value1 := &TempStruct{
			Name:  "value1",
			Value: 1,
		}
		value2 := &TempStruct{
			Name:  "value2",
			Value: 2,
		}

		source := []*TempStruct{value1, value2}
		forEachFunc := func(item *TempStruct) {
			item.Value = item.Value * 2
		}

		ForEach(source, forEachFunc)

		expected := []*TempStruct{value1, value2}
		assert.Equal(t, expected, source)
	})
}

func TestForEachWithError(t *testing.T) {
	t.Run("print integers", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		forEachFunc := func(item int) error {

			fmt.Println(item)
			return nil
		}

		err := ForEachWithError(source, forEachFunc)
		assert.NoError(t, err)
	})
	t.Run("print integers error", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}
		forEachFunc := func(item int) error {
			if item == 3 {
				return errors.New("error")
			}
			fmt.Println(item)
			return nil
		}

		err := ForEachWithError(source, forEachFunc)
		assert.Error(t, err)
	})
}

func TestCloneStringList(t *testing.T) {
	tests := []struct {
		name   string
		source []string
		want   []string
	}{
		{
			name:   "empty list",
			source: []string{},
			want:   []string{},
		},
		{
			name:   "single element",
			source: []string{"element"},
			want:   []string{"element"},
		},
		{
			name:   "multiple elements",
			source: []string{"hello", "world"},
			want:   []string{"hello", "world"},
		},
		{
			name:   "nil list",
			source: nil,
			want:   []string{},
		},
		{
			name:   "empty list",
			source: []string{},
			want:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CloneList(tt.source)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMapReturnWithError(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		mappingFunc := func(data int) (int, error) {
			return data * 2, nil
		}

		result, err := MapReturnWithError(source, mappingFunc)
		assert.NoError(t, err)

		expected := []int{2, 4, 6, 8, 10}
		assert.Equal(t, expected, result)
	})

	t.Run("some_element_has_Error", func(t *testing.T) {

		source := []int{1, 2, 3, 4, 5}

		mappingFunc := func(data int) (int, error) {
			if data == 3 {
				return 0, errors.New("fake error for 3")
			}
			return data * 2, nil
		}

		result, err := MapReturnWithError(source, mappingFunc)
		assert.Error(t, err)
		assert.Equal(t, "error mapping at index:'2', error: fake error for 3", err.Error())

		assert.Nil(t, result)
	})

}

func TestHigherOrderFunction_Sort(t *testing.T) {
	t.Run("Success_Int", func(t *testing.T) {

		source := []int{5, 2, 8, 1, 9}

		sortFunc := func(i, j int) bool {
			return source[i] < source[j]
		}

		sorted := Sort(source, sortFunc)

		expected := []int{1, 2, 5, 8, 9}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_Int_reverse", func(t *testing.T) {
		source := []int{5, 2, 8, 1, 9}

		sortFunc := func(i, j int) bool {
			return source[i] > source[j]
		}

		sorted := Sort(source, sortFunc)

		expected := []int{9, 8, 5, 2, 1}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_String", func(t *testing.T) {

		source := []string{"c", "a", "b"}

		sortFunc := func(i, j int) bool {
			return source[i] < source[j]
		}

		sorted := Sort(source, sortFunc)

		expected := []string{"a", "b", "c"}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_String_reverse", func(t *testing.T) {

		source := []string{"c", "a", "b"}

		sortFunc := func(i, j int) bool {
			return source[i] > source[j]
		}

		sorted := Sort(source, sortFunc)

		expected := []string{"c", "b", "a"}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_CustomType", func(t *testing.T) {

		type Person struct {
			Name string
			Age  int
		}

		source := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		sortFunc := func(i, j int) bool {
			return source[i].Age < source[j].Age
		}

		sorted := Sort(source, sortFunc)

		expected := []Person{
			{"Bob", 25},
			{"Alice", 30},
			{"Charlie", 35},
		}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_CustomType_reverse", func(t *testing.T) {

		type Person struct {
			Name string
			Age  int
		}

		source := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		sortFunc := func(i, j int) bool {
			return source[i].Age > source[j].Age
		}

		sorted := Sort(source, sortFunc)

		expected := []Person{
			{"Charlie", 35},
			{"Alice", 30},
			{"Bob", 25},
		}
		assert.Equal(t, expected, sorted)
	})

	t.Run("Success_sort_2_layers_of_customerType_sort_customerCode_and_SalesOrderNumber", func(t *testing.T) {

		type SalesOrder struct {
			CustomerCode     string
			SalesOrderNumber string
			Amount           float64
		}

		source := []SalesOrder{
			{"C2", "S2", 200},
			{"C1", "S3", 300},
			{"C2", "S4", 400},
			{"C1", "S1", 100},
		}

		sortFunc := func(i, j int) bool {
			if source[i].CustomerCode == source[j].CustomerCode {
				return source[i].SalesOrderNumber < source[j].SalesOrderNumber
			}
			return source[i].CustomerCode < source[j].CustomerCode
		}

		sorted := Sort(source, sortFunc)

		expected := []SalesOrder{
			{"C1", "S1", 100},
			{"C1", "S3", 300},
			{"C2", "S2", 200},
			{"C2", "S4", 400},
		}
		assert.Equal(t, expected, sorted)
	})
}

func TestExists(t *testing.T) {
	testCases := []struct {
		name      string
		input     []int
		condition func(int) bool
		expected  bool
	}{
		{
			name:      "Any element greater than 10",
			input:     []int{1, 2, 3, 4, 11},
			condition: func(n int) bool { return n > 10 },
			expected:  true,
		},
		{
			name:      "No element greater than 10",
			input:     []int{1, 2, 3, 4},
			condition: func(n int) bool { return n > 10 },
			expected:  false,
		},
		{
			name:      "Empty slice",
			input:     []int{},
			condition: func(n int) bool { return n > 10 },
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Exists(tc.input, tc.condition)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// Test for Max function
func TestMax(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		found    bool
	}{
		{[]int{1, 2, 3, 4, 5}, 5, true},    // Typical case
		{[]int{-10, -3, -6, -1}, -1, true}, // Negative numbers
		{[]int{5}, 5, true},                // Single element
		{[]int{}, 0, false},                // Empty slice
	}

	for _, test := range tests {
		result, found := Max(test.input)
		if result != test.expected || found != test.found {
			t.Errorf("Max(%v) = (%v, %v); expected (%v, %v)",
				test.input, result, found, test.expected, test.found)
		}
	}
}

// Test for Min function
func TestMin(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		found    bool
	}{
		{[]int{1, 2, 3, 4, 5}, 1, true},        // Typical case
		{[]int{-10, -3, -6, -1}, -10, true},    // Negative numbers
		{[]int{5}, 5, true},                    // Single element
		{[]int{}, 0, false},                    // Empty slice
		{[]int{5, 3, 7, 1, 8}, 1, true},        // Min value at the end
		{[]int{1, 5, 3, 7, 8}, 1, true},        // Min value at the beginning
		{[]int{5, 3, 1, 7, 8}, 1, true},        // Min value in the middle
		{[]int{5, 3, 1, 7, 1}, 1, true},        // Multiple occurrences of the min value
		{[]int{1, 1, 1, 1}, 1, true},           // All values are the same
		{[]int{100, 90, 80, 70, 60}, 60, true}, // Descending order
		{[]int{60, 70, 80, 90, 100}, 60, true}, // Ascending order
	}

	for _, test := range tests {
		result, found := Min(test.input)
		if result != test.expected || found != test.found {
			t.Errorf("Min(%v) = (%v, %v); expected (%v, %v)",
				test.input, result, found, test.expected, test.found)
		}
	}
}

// Test with float64 for Max function
func TestMaxFloat(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
		found    bool
	}{
		{[]float64{1.1, 2.2, 3.3, 4.4}, 4.4, true}, // Typical case
		{[]float64{-10.1, -3.5, -6.7}, -3.5, true}, // Negative floats
		{[]float64{5.5}, 5.5, true},                // Single element
		{[]float64{}, 0, false},                    // Empty slice
	}

	for _, test := range tests {
		result, found := Max(test.input)
		if result != test.expected || found != test.found {
			t.Errorf("Max(%v) = (%v, %v); expected (%v, %v)",
				test.input, result, found, test.expected, test.found)
		}
	}
}

// Test with float64 for Min function
func TestMinFloat(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
		found    bool
	}{
		{[]float64{1.1, 2.2, 3.3, 4.4}, 1.1, true},  // Typical case
		{[]float64{-10.1, -3.5, -6.7}, -10.1, true}, // Negative floats
		{[]float64{5.5}, 5.5, true},                 // Single element
		{[]float64{}, 0, false},                     // Empty slice
	}

	for _, test := range tests {
		result, found := Min(test.input)
		if result != test.expected || found != test.found {
			t.Errorf("Min(%v) = (%v, %v); expected (%v, %v)",
				test.input, result, found, test.expected, test.found)
		}
	}
}

func TestPartition(t *testing.T) {
	// Test cases for Partition
	t.Run("PartitionEvenOdd", func(t *testing.T) {
		ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evens, odds := Partition(ints, func(v int) bool {
			return v%2 == 0
		})
		expectedEvens := []int{2, 4, 6, 8, 10}
		expectedOdds := []int{1, 3, 5, 7, 9}

		assert.Equal(t, expectedEvens, evens)
		assert.Equal(t, expectedOdds, odds)
	})

	t.Run("PartitionStringsByLength", func(t *testing.T) {
		strings := []string{"Go", "is", "awesome", "!", "I", "love", "coding"}
		longWords, shortWords := Partition(strings, func(v string) bool {
			return len(v) > 2
		})
		expectedLong := []string{"awesome", "love", "coding"}
		expectedShort := []string{"Go", "is", "!", "I"}

		assert.Equal(t, expectedLong, longWords)
		assert.Equal(t, expectedShort, shortWords)
	})

	t.Run("PartitionEmptySlice", func(t *testing.T) {
		empty := []int{}
		trues, falses := Partition(empty, func(v int) bool {
			return v > 0
		})

		assert.Equal(t, []int{}, trues)
		assert.Equal(t, []int{}, falses)
	})

	t.Run("PartitionAllTrue", func(t *testing.T) {
		ints := []int{1, 2, 3, 4, 5}
		allTrue, noneFalse := Partition(ints, func(v int) bool {
			return true
		})

		assert.Equal(t, ints, allTrue)
		assert.Equal(t, []int{}, noneFalse)
	})

	t.Run("PartitionAllFalse", func(t *testing.T) {
		ints := []int{1, 2, 3, 4, 5}
		noneTrue, allFalse := Partition(ints, func(v int) bool {
			return false
		})

		assert.Equal(t, []int{}, noneTrue)
		assert.Equal(t, ints, allFalse)
	})
}

func TestMaxBy(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}
	t.Run("MaxByAge", func(t *testing.T) {
		people := []person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		oldest, found := MaxBy(people, func(p person) int {
			return p.Age
		})

		assert.True(t, found)
		assert.Equal(t, "Charlie", oldest.Name)
		assert.Equal(t, 35, oldest.Age)
	})

	t.Run("MaxByAgeOnePerson", func(t *testing.T) {
		people := []person{
			{"Alice", 30},
		}

		oldest, found := MaxBy(people, func(p person) int {
			return p.Age
		})

		assert.True(t, found)
		assert.Equal(t, "Alice", oldest.Name)
		assert.Equal(t, 30, oldest.Age)
	})

	t.Run("MaxByEmptySlice", func(t *testing.T) {
		people := []person{}

		_, found := MaxBy(people, func(p person) int {
			return p.Age
		})

		assert.False(t, found)
	})

	t.Run("MaxByNegativeValues", func(t *testing.T) {
		numbers := []int{-10, -5, -20, -1}

		max, found := MaxBy(numbers, func(n int) int {
			return n
		})

		assert.True(t, found)
		assert.Equal(t, -1, max)
	})
}

func TestMinBy(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}
	t.Run("MinByAge", func(t *testing.T) {
		people := []person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		youngest, found := MinBy(people, func(p person) int {
			return p.Age
		})

		assert.True(t, found)
		assert.Equal(t, "Bob", youngest.Name)
		assert.Equal(t, 25, youngest.Age)
	})

	t.Run("MinByAgeOnePerson", func(t *testing.T) {
		people := []person{
			{"Alice", 30},
		}

		youngest, found := MinBy(people, func(p person) int {
			return p.Age
		})

		assert.True(t, found)
		assert.Equal(t, "Alice", youngest.Name)
		assert.Equal(t, 30, youngest.Age)
	})

	t.Run("MinByEmptySlice", func(t *testing.T) {
		people := []person{}

		_, found := MinBy(people, func(p person) int {
			return p.Age
		})

		assert.False(t, found)
	})

	t.Run("MinByNegativeValues", func(t *testing.T) {
		numbers := []int{-10, -5, -20, -1}

		min, found := MinBy(numbers, func(n int) int {
			return n
		})

		assert.True(t, found)
		assert.Equal(t, -20, min)
	})
}

func TestCount(t *testing.T) {
	t.Run("CountEvenNumbers", func(t *testing.T) {
		// Given a slice of numbers
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		// When counting even numbers
		evenCount := Count(numbers, func(n int) bool {
			return n%2 == 0
		})

		// Then the count should be 5
		assert.Equal(t, 5, evenCount)
	})

	t.Run("CountOddNumbers", func(t *testing.T) {
		// Given a slice of numbers
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		// When counting odd numbers
		oddCount := Count(numbers, func(n int) bool {
			return n%2 != 0
		})

		// Then the count should be 5
		assert.Equal(t, 5, oddCount)
	})

	t.Run("CountStringsWithLengthGreaterThan3", func(t *testing.T) {
		// Given a slice of strings
		strings := []string{"Go", "is", "awesome", "I", "love", "coding"}

		// When counting strings with length greater than 3
		longStringCount := Count(strings, func(s string) bool {
			return len(s) > 3
		})

		// Then the count should be 3
		assert.Equal(t, 3, longStringCount)
	})

	t.Run("CountEmptySlice", func(t *testing.T) {
		// Given an empty slice of integers
		empty := []int{}

		// When counting elements in an empty slice
		count := Count(empty, func(n int) bool {
			return n > 0
		})

		// Then the count should be 0
		assert.Equal(t, 0, count)
	})

	t.Run("CountPeopleOlderThan30", func(t *testing.T) {
		// Given a slice of Person structs
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
			{"Dave", 40},
		}

		// When counting people older than 30
		olderThan30Count := Count(people, func(p Person) bool {
			return p.Age > 30
		})

		// Then the count should be 2
		assert.Equal(t, 2, olderThan30Count)
	})
}

func TestCurry(t *testing.T) {
	// Test with an addition function
	t.Run("IntegerAddition", func(t *testing.T) {
		add := func(a, b int) int {
			return a + b
		}
		curriedAdd := Curry(add)

		add5 := curriedAdd(5)

		result := add5(3)
		expected := 8

		if result != expected {
			t.Errorf("curriedAdd(5)(3) = %d; want %d", result, expected)
		}
	})

	// Test with a string concatenation function
	t.Run("StringConcatenation", func(t *testing.T) {
		concat := func(a, b string) string {
			return a + b
		}
		curriedConcat := Curry(concat)

		helloConcat := curriedConcat("Hello, ")

		resultStr := helloConcat("World!")
		expectedStr := "Hello, World!"

		if resultStr != expectedStr {
			t.Errorf(`curriedConcat("Hello, ")("World!") = %s; want %s`, resultStr, expectedStr)
		}
	})

	// Test with a multiplication function
	t.Run("FloatMultiplication", func(t *testing.T) {
		multiply := func(a, b float64) float64 {
			return a * b
		}
		curriedMultiply := Curry(multiply)

		multiplyBy10 := curriedMultiply(10)

		resultMul := multiplyBy10(3.5)
		expectedMul := 35.0

		if resultMul != expectedMul {
			t.Errorf("curriedMultiply(10)(3.5) = %f; want %f", resultMul, expectedMul)
		}
	})

	// Multiple curried functions test case
	t.Run("MultipleCurriedFunctions", func(t *testing.T) {
		// First curried function: addition
		add := func(a, b int) int {
			return a + b
		}
		curriedAdd := Curry(add)

		// Second curried function: multiplication
		multiply := func(a, b int) int {
			return a * b
		}
		curriedMultiply := Curry(multiply)

		// Apply the curried addition and multiplication
		add5 := curriedAdd(5)
		multiplyBy2 := curriedMultiply(2)

		// Combine them by adding first, then multiplying the result
		result := multiplyBy2(add5(3)) // (5 + 3) * 2 = 16
		expected := 16

		if result != expected {
			t.Errorf("multiplyBy2(add5(3)) = %d; want %d", result, expected)
		}

		// Another case, add 7, then multiply by 4
		add7 := curriedAdd(7)
		multiplyBy4 := curriedMultiply(4)

		result2 := multiplyBy4(add7(2)) // (7 + 2) * 4 = 36
		expected2 := 36

		if result2 != expected2 {
			t.Errorf("multiplyBy4(add7(2)) = %d; want %d", result2, expected2)
		}
	})

	// Additional test case for chaining curried functions with different types
	t.Run("ChainingDifferentTypes", func(t *testing.T) {
		// Curried string concatenation function
		concat := func(a, b string) string {
			return a + b
		}
		curriedConcat := Curry(concat)

		// Curried string length function
		stringLength := func(a string, b string) int {
			return len(a + b)
		}
		curriedStringLength := Curry(stringLength)

		// Chain the functions: concatenate first, then get the length
		helloConcat := curriedConcat("Hello, ")
		length := curriedStringLength(helloConcat("World!"))

		resultLength := length("!")
		expectedLength := 14 // "Hello, World!!" has 14 characters

		if resultLength != expectedLength {
			t.Errorf(`curriedStringLength(helloConcat("World!"))("!") = %d; want %d`, resultLength, expectedLength)
		}
	})

}

func TestCompose(t *testing.T) {
	// Integer functions for composition
	multiplyBy2 := func(x int) int {
		return x * 2
	}

	add3 := func(x int) int {
		return x + 3
	}

	// String functions for composition
	upper := func(s string) string {
		return strings.ToUpper(s)
	}

	addExclamation := func(s string) string {
		return s + "!"
	}

	// Run sub-tests
	t.Run("IntegerCompose", func(t *testing.T) {
		composed := Compose(multiplyBy2, add3)

		t.Run("TestWith5", func(t *testing.T) {
			result := composed(5)
			expected := 16 // (5 + 3) * 2 = 16
			if result != expected {
				t.Errorf("composed(5) = %d; want %d", result, expected)
			}
		})

		t.Run("TestWith7", func(t *testing.T) {
			result := composed(7)
			expected := 20 // (7 + 3) * 2 = 20
			if result != expected {
				t.Errorf("composed(7) = %d; want %d", result, expected)
			}
		})
	})

	t.Run("StringCompose", func(t *testing.T) {
		composedString := Compose(addExclamation, upper)

		t.Run("TestWithHello", func(t *testing.T) {
			resultStr := composedString("hello")
			expectedStr := "HELLO!"
			if resultStr != expectedStr {
				t.Errorf(`composedString("hello") = %s; want %s`, resultStr, expectedStr)
			}
		})

		t.Run("TestWithGo", func(t *testing.T) {
			resultStr := composedString("go")
			expectedStr := "GO!"
			if resultStr != expectedStr {
				t.Errorf(`composedString("go") = %s; want %s`, resultStr, expectedStr)
			}
		})
	})

	t.Run("EdgeCases", func(t *testing.T) {
		t.Run("EmptyString", func(t *testing.T) {
			composedString := Compose(addExclamation, upper)
			resultStr := composedString("")
			expectedStr := "!"
			if resultStr != expectedStr {
				t.Errorf(`composedString("") = %s; want %s`, resultStr, expectedStr)
			}
		})

		t.Run("NegativeNumber", func(t *testing.T) {
			composed := Compose(multiplyBy2, add3)
			result := composed(-3) // (-3 + 3) * 2 = 0
			expected := 0
			if result != expected {
				t.Errorf("composed(-3) = %d; want %d", result, expected)
			}
		})
	})
}
