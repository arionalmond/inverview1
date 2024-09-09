package main

import (
	"fmt"
)

const x = 0
const y = 1

func main() {

	board := [][]int{
		{0, 0, 0, -1, -1},
		{0, 0, -1, 0, 0},
		{0, -1, 0, -1, 0},
		{0, 0, -1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	test := [][]int{
		{1, 1}, {5, 3}, {5, 1}, {6, 0}, {6, 4}, {0, 0}, {2, 2},
	}

	i := 0
	result := [][]int{}

	for i < len(test) {
		cord := test[i]

		// err := validateInput(board, cord)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// UP/DOWN {-1/+1, cord[y]}
		if up := cord[x] - 1; up >= 0 {
			if board[up][cord[y]] != -1 {
				r := []int{up, cord[y]}
				result = append(result, r)
			}
		}

		if down := cord[x] + 1; down < len(board) {
			if board[down][cord[y]] != -1 {
				r := []int{down, cord[y]}
				result = append(result, r)
			}
		}

		// LEFT/RIGHT {cord[x], -1/+1}
		if left := cord[y] - 1; left >= 0 {
			if board[cord[x]][left] != -1 {
				r := []int{cord[x], left}
				result = append(result, r)
			}
		}

		if right := cord[y] + 1; right <= len(board[i])-1 {
			if board[cord[x]][right] != -1 {
				r := []int{cord[x], right}
				result = append(result, r)
			}
		}

		// (0, 1), (1, 0)
		fmt.Printf("\ncord {%d,%d} \nresults: %v",
			cord[x], cord[y],
			result,
		)
		result = [][]int{}
		i++
	}
}

// func validateInput(board [][]int, pos []int) error {
// 	if pos[0] < 0 || pos[0] > len(board)-1 {
// 		err := errors.New("row position")
// 		return err
// 	}

// 	if pos[1] < 0 || pos[1] > len(board[pos[0]])-1 {
// 		err := errors.New("col position")
// 		return err
// 	}
// 	return nil
// }
