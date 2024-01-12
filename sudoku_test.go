package main

import (
	"reflect"
	"testing"
)

func TestFindEmptyCells(t *testing.T) {
	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	expected := []BoardCellCoordinates{
		{Row: 0, Col: 2},
		{Row: 0, Col: 3},
		{Row: 0, Col: 5},
		{Row: 0, Col: 6},
		{Row: 0, Col: 7},
		{Row: 0, Col: 8},

		{Row: 1, Col: 1},
		{Row: 1, Col: 2},
		{Row: 1, Col: 6},
		{Row: 1, Col: 7},
		{Row: 1, Col: 8},

		{Row: 2, Col: 0},
		{Row: 2, Col: 3},
		{Row: 2, Col: 4},
		{Row: 2, Col: 5},
		{Row: 2, Col: 6},
		{Row: 2, Col: 8},

		{Row: 3, Col: 1},
		{Row: 3, Col: 2},
		{Row: 3, Col: 3},
		{Row: 3, Col: 5},
		{Row: 3, Col: 6},
		{Row: 3, Col: 7},

		{Row: 4, Col: 1},
		{Row: 4, Col: 2},
		{Row: 4, Col: 4},
		{Row: 4, Col: 6},
		{Row: 4, Col: 7},

		{Row: 5, Col: 1},
		{Row: 5, Col: 2},
		{Row: 5, Col: 3},
		{Row: 5, Col: 5},
		{Row: 5, Col: 6},
		{Row: 5, Col: 7},

		{Row: 6, Col: 0},
		{Row: 6, Col: 2},
		{Row: 6, Col: 3},
		{Row: 6, Col: 4},
		{Row: 6, Col: 5},
		{Row: 6, Col: 8},

		{Row: 7, Col: 0},
		{Row: 7, Col: 1},
		{Row: 7, Col: 2},
		{Row: 7, Col: 6},
		{Row: 7, Col: 7},

		{Row: 8, Col: 0},
		{Row: 8, Col: 1},
		{Row: 8, Col: 2},
		{Row: 8, Col: 3},
		{Row: 8, Col: 5},
		{Row: 8, Col: 6},
	}
	result := FindEmptyCells(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Empty cells list does not match the expected one. Expected:\n%v\n\nGot:\n%v", expected, result)
	}
}

func TestRowIsValid(t *testing.T) {
	good_input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0}, // no duplicate
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	if RowIsValid(good_input, BoardCellCoordinates{Row: 0, Col: 0}) == false {
		t.Errorf("RowIsValid() returns false, although has to return true")
	}

	bad_input := [][]int{
		{5, 3, 0, 9, 9, 0, 0, 0, 0}, // duplicate
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if RowIsValid(bad_input, BoardCellCoordinates{Row: 0, Col: 0}) == true {
		t.Errorf("RowIsValid() returns true, although has to return false")
	}
}

func TestSquareIsValid(t *testing.T) {
	good_input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0}, // no duplicate, no zeros
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 1, 2, 3, 0, 0, 3},
		{4, 0, 0, 4, 5, 6, 0, 0, 7},
		{7, 0, 0, 7, 8, 9, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	if SquareIsValid(good_input, BoardCellCoordinates{Row: 4, Col: 4}) == false {
		t.Errorf("SquareIsValid() returns false, although has to return true")
	}

	good_input = [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 1, 2, 3, 0, 0, 3},
		{4, 0, 0, 0, 5, 0, 0, 0, 7}, // no duplicate, has zeros
		{7, 0, 0, 7, 8, 9, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	if SquareIsValid(good_input, BoardCellCoordinates{Row: 4, Col: 4}) == false {
		t.Errorf("SquareIsValid() returns false, although has to return true")
	}

	bad_input := [][]int{
		// duplicate in col 3
		{5, 3, 0, 0, 0, 0, 0, 0, 0},
		{6, 0, 0, 0, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 6, 6, 6, 0, 0, 3},
		{4, 0, 0, 6, 9, 6, 0, 0, 1},
		{7, 0, 0, 6, 6, 6, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if ColIsValid(bad_input, BoardCellCoordinates{Row: 4, Col: 4}) == true {
		t.Errorf("ColIsValid() returns true, although has to return false")
	}
}

func TestColIsValid(t *testing.T) {
	good_input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0}, // no duplicate
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	if ColIsValid(good_input, BoardCellCoordinates{Row: 0, Col: 3}) == false {
		t.Errorf("ColIsValid() returns false, although has to return true")
	}

	bad_input := [][]int{
		// duplicate in col 3
		{5, 3, 0, 9, 0, 0, 0, 0, 0},
		{6, 0, 0, 9, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if ColIsValid(bad_input, BoardCellCoordinates{Row: 0, Col: 3}) == true {
		t.Errorf("ColIsValid() returns true, although has to return false")
	}
}

func TestSolveSudoku(t *testing.T) {
	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	expected := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}
