package conditional

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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfThen(t *testing.T) {
	result := IfThen(true, 1, 2)
	assert.Equal(t, 1, result)

	result = IfThen(false, 1, 2)
	assert.Equal(t, 2, result)

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

	resultTempStruct := IfThen(true, value1, value2)
	assert.Equal(t, value1, resultTempStruct)

	resultTempStruct = IfThen(false, value1, value2)
	assert.Equal(t, value2, resultTempStruct)

	resultTempStruct = IfThen(value1.Value < value2.Value, value2, value1)
	assert.Equal(t, value2, resultTempStruct)

	resultInt := IfThen(false, 1, 2)
	assert.Equal(t, 2, resultInt)

	resultString := IfThen(true, "1", "2")
	assert.Equal(t, "1", resultString)

	resultString = IfThen(len("ABC") < 4, "less than length 4.", "more than length 4.")
	assert.Equal(t, "less than length 4.", resultString)

	resultString = IfThen(len("ABD") < 4, "0", "1")
	assert.Equal(t, "0", resultString)

	resultString = IfThen(len("ABCD") < 4, "0", "1")
	assert.Equal(t, "1", resultString)

	tempInfo1 := TempStruct{
		Name:  IfThen(len("ABCD") < 4, "1", "0"),
		Value: 2,
	}
	expected := TempStruct{
		Name:  "0",
		Value: 2,
	}
	assert.EqualValues(t, expected, tempInfo1)

	tempInfo1 = TempStruct{
		Name:  IfThen(len("ABCD") == 4, "ABCD", ""),
		Value: IfThen(len("ABCD") > 4, 4, 0),
	}
	expected = TempStruct{
		Name:  "ABCD",
		Value: 0,
	}
	assert.EqualValues(t, expected, tempInfo1)
}
func TestAll(t *testing.T) {
	t.Run("TestAllWithIntegers", func(t *testing.T) {
		// Test with integers
		ints := []int{2, 4, 6, 8}
		isEven := func(n int) bool {
			return n%2 == 0
		}
		assert.True(t, ForAll(ints, isEven))

		ints = []int{2, 4, 6, 7}
		assert.False(t, ForAll(ints, isEven))
	})

	t.Run("TestAllWithStrings", func(t *testing.T) {

		// Test with strings
		strings := []string{"apple", "apricot", "avocado"}
		startsWithA := func(s string) bool {
			return s[0] == 'a'
		}
		assert.True(t, ForAll(strings, startsWithA))

		strings = []string{"apple", "banana", "avocado"}
		assert.False(t, ForAll(strings, startsWithA))
	})

	t.Run("TestAllWithStructs", func(t *testing.T) {
		// Test with custom struct
		type TempStruct struct {
			Name  string
			Value int
		}
		structs := []TempStruct{
			{Name: "one", Value: 1},
			{Name: "two", Value: 2},
			{Name: "three", Value: 3},
		}
		valueLessThanFour := func(ts TempStruct) bool {
			return ts.Value < 4
		}
		assert.True(t, ForAll(structs, valueLessThanFour))

		structs = []TempStruct{
			{Name: "one", Value: 1},
			{Name: "two", Value: 2},
			{Name: "four", Value: 4},
		}
		assert.False(t, ForAll(structs, valueLessThanFour))
	})

}
