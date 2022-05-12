package main

import (
	"fmt"
)

var columnNumber = 4
var rowNumber = 4

func canBePlaced(board [][]int, row, column int) bool {
	for i := 0; i < column; i++ {
		if board[row][i] == 1 {
			return false
		}
	}

	//chequear diagonal superior
	for i, j := row, column; i >= 0; i-- {
		if j < 0 {
			break
		}
		if board[i][j] == 1 {
			return false
		}
		j--
	}

	//chequear diagonal inferior
	for i, j := row, column; j >= 0; j-- {
		if i >= rowNumber {
			break
		}
		if board[i][j] == 1 {
			return false
		}
		i++
	}

	return true
}

func solveNQueens(board [][]int, column int) bool {
	if column >= columnNumber {
		return true
	}

	for i := 0; i < columnNumber; i++ {
		if canBePlaced(board, i, column) {
			board[i][column] = 1

			if solveNQueens(board, column+1) == true {
				return true
			}

			board[i][column] = 0
		}

	}
	return false
}

func main() {
	var board [][]int = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	if solveNQueens(board, 0) == false {
		fmt.Println("No se encontro solucion")
	} else {
		fmt.Println("Solucion encontrada")
		fmt.Println(board)
	}
}