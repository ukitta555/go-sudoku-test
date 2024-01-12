package main

/*
Struct that represent coordinates of a cell in a sudoku board.
Coordinates (programmers love to count from 0 for some reason, so last coordinate would be 8):
(0, 0) (0, 1) ... (0, 8)
(1, 0) (1, 1) ... (1, 8)
........................
(8, 0) (8, 1) ... (8, 8)
*/
type BoardCellCoordinates struct {
	Row int
	Col int
}

/*
Struct that represent coordinates of a SQUARE in a sudoku board. 
This is semantically different from the BoardCellCoordinates struct.
The struct in question encodes the coordinates of one of the nine 3x3 squares in the sudoku board.
This struct is used for finding the coordinates of a top-left cell in a sudoku board.

Coordinates:
(0, 0) (0, 1) (0, 2)
(1, 0) (1, 1) (1, 2)
(2, 0) (2, 1) (2, 2)
*/
type BoardSquareCoordinates struct {
	Row int
	Col int
}

// Finds empty cells (duh!) in the provided board.
func FindEmptyCells(board [][]int) []BoardCellCoordinates {
	var emptyCells []BoardCellCoordinates

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				emptyCells = append(emptyCells, BoardCellCoordinates{row, col})
			}
		}
	}
	return emptyCells
}

// Check whether there are any conflicts in the row where the backtracking algorithm just made a change.
func RowIsValid(board [][]int, cellToValidate BoardCellCoordinates) bool {
	var numbersPresent [10]bool
	for col_index := 0; col_index < 9; col_index++ {
		var currentNumber int = board[cellToValidate.Row][col_index]

		// don't care about cells that haven't been filled out yet
		if currentNumber == 0 {
			continue
		}

		// check whether we have already seen this number before while scanning
		if numbersPresent[currentNumber] == false {
			numbersPresent[currentNumber] = true // set this number as present in the row
		} else {
			return false // if already present, there is a conflict; return false
		}
	}
	// no conflicts; return true
	return true
}

// Check whether there are any conflicts in the column where the backtracking algorithm just made a change.
func ColIsValid(board [][]int, cellToValidate BoardCellCoordinates) bool {
	var numbersPresent [10]bool

	for row_index := 0; row_index < 9; row_index++ {
		var currentNumber int = board[row_index][cellToValidate.Col]
		// don't care about cells that haven't been filled out yet
		if currentNumber == 0 {
			continue
		}

		// check whether we have already seen this number before while scanning
		if numbersPresent[currentNumber] == false {
			numbersPresent[currentNumber] = true // set this number as present in the column
		} else {
			return false // if already present, there is a conflict; return false
		}
	}
	// no conflicts; return true
	return true
}

func getTopLeftCellOfSquare(cell BoardCellCoordinates) BoardSquareCoordinates {
	return BoardSquareCoordinates{cell.Row / 3, cell.Col / 3}
}

func SquareIsValid(board [][]int, cellToValidate BoardCellCoordinates) bool {
	var numbersPresent [10]bool
	var topLeftCellOfSquare BoardSquareCoordinates = getTopLeftCellOfSquare(cellToValidate)

	// iterate through the square
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var currentNumber int = board[topLeftCellOfSquare.Row*3+i][topLeftCellOfSquare.Col*3+j]
			// don't care about cells that haven't been filled out yet
			if currentNumber == 0 {
				continue
			}

			// check whether we have already seen this number before while scanning
			if numbersPresent[currentNumber] == false {
				numbersPresent[currentNumber] = true // set this number as present in the column
			} else {
				return false // if already present, there is a conflict; return false
			}
		}
	}
	// no conflicts; return true
	return true
}

// Use backtracking approach to try all possibilities for each of the empty cells.
func recursive_bruteforce(board [][]int, emptyCells []BoardCellCoordinates, emptyCellToProcess int) bool {
	// base case - in case we filled all cells without breaking the rules, the sudoku is solved.
	if len(emptyCells) == emptyCellToProcess {
		return true
	}
	currentEmptyCell := emptyCells[emptyCellToProcess]
	var solutionFound bool = false

	for valueToInsertInTheCell := 1; valueToInsertInTheCell <= 9; valueToInsertInTheCell++ {
		board[currentEmptyCell.Row][currentEmptyCell.Col] = valueToInsertInTheCell
		processedCell := currentEmptyCell
		if RowIsValid(board, processedCell) && ColIsValid(board, processedCell) && SquareIsValid(board, processedCell) {
			solutionFound = recursive_bruteforce(board, emptyCells, emptyCellToProcess+1)
		}
		if !solutionFound {
			board[processedCell.Row][processedCell.Col] = 0
		} else {
			return solutionFound // short-circuit the search in case a solution was found
		}
	}
	return false
}


func SolveSudoku(board [][]int) [][]int {
	// don't know whether I'm able to change the function interface to return an error in case of bad input...
	// I'll leave it here just in case

	// for row := 0; row < 9; row++ {
	// 	for col := 0; col < 9; col++ {
	// 		cellToValidate := BoardCellCoordinates {Row: row, Col: col}
	// 		if !RowIsValid(board, cellToValidate) || !ColIsValid(board, cellToValidate) || !SquareIsValid(board, cellToValidate) {
	// 			panic("Bad sudoku puzzle provided!")
	// 		}
	// 	}
	// }

	// assuming that sudoku board is correct when submitted to this function
	emptyCells := FindEmptyCells(board)
	// sudoku already solved
	if len(emptyCells) == 0 {
		return board
	}
	recursive_bruteforce(board, emptyCells, 0)

	return board
}
