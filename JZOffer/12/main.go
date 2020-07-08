package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			flag := make([][]bool, len(board))
			for v := 0; v < len(board); v++ {
				tmp := make([]bool, len(board[0]))
				flag[v] = tmp
			}
			//fmt.Println(flag)
			if cursive(board, flag, []byte(word), 0, i, j) {
				return true
			}
		}
	}
	return false
}

func cursive(board [][]byte, flag [][]bool, word []byte, i, posX, posY int) bool {
	if i >= len(word) {
		return true
	}
	if posY < 0 || posY >= len(board[0]) {
		return false
	}
	if posX < 0 || posX >= len(board) {
		return false
	}

	if flag[posX][posY] == false && board[posX][posY] == word[i] {
		flag[posX][posY] = true
		if cursive(board, flag, word, i+1, posX+1, posY) ||
			cursive(board, flag, word, i+1, posX-1, posY) ||
			cursive(board, flag, word, i+1, posX, posY+1) ||
			cursive(board, flag, word, i+1, posX, posY-1) {
			return true
		}
		flag[posX][posY] = false
	}
	return false
}

func main() {
	board := make([][]byte, 2)
	board[0] = []byte("ab")
	board[1] = []byte("cd")
	//board[2] = []byte("")
	b := exist(board, "dcba")
	fmt.Println(b)
}
