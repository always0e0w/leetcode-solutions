package main

import (
	"strings"
)

// N皇后
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	var backtrack func([]string, int)
	backtrack = func(board []string, row int) {
		if row == n {
			b := make([]string, n)
			copy(b, board)
			res = append(res, b)
			return
		}
		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}
			curRow := []byte(board[row])
			old := curRow[col]
			curRow[col] = 'Q'
			board[row] = string(curRow)
			backtrack(board, row+1)
			curRow[col] = old
			board[row] = string(curRow)
		}
	}
	board := make([]string, n)
	for i := range board {
		board[i] = strings.Repeat(".", n)
	}
	backtrack(board, 0)
	return res
}

func isValid(board []string, row, col int) bool {
	for i := range board {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 右上方
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 左上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
