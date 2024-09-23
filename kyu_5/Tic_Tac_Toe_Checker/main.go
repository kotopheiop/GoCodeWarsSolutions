package main

import "fmt"

func IsSolved(board [3][3]int) int {
	// Строки, столбцы и диагонали
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != 0 {
			return board[i][0]
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != 0 {
			return board[0][i]
		}
	}

	// Диагонали
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != 0 {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != 0 {
		return board[0][2]
	}

	// Пустые клетки
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				return -1 // Игра не окончена
			}
		}
	}

	return 0 // Ничья
}

func main() {
	fmt.Println(IsSolved([3][3]int{ // -1
		{0, 0, 1},
		{0, 1, 2},
		{2, 1, 0},
	}))

	fmt.Println(IsSolved([3][3]int{ // 1
		{1, 1, 1},
		{0, 2, 2},
		{0, 0, 0},
	}))

	fmt.Println(IsSolved([3][3]int{ // 0
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 1},
	}))
}
