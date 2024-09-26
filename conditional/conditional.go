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

// IfThen executes an if-else operation in a single line.
// Examples:
//   - IfThen(len("ABC") < 4, "less than 4", "more than 4") returns "less than 4".
//   - IfThen(len("ABD") < 4, 0, 1) returns 0.
//   - IfThen(len("ABCD") < 4, "0", "1") returns "1".
func IfThen[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func ForAll[T any](elements []T, condition func(T) bool) bool {
	for _, e := range elements {
		if !condition(e) {
			return false
		}
	}
	return true
}
