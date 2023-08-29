package main

import (
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type cell struct {
	isMined     bool
	numMinNeigh int
	isClicked   bool
	button      *widget.Button
}

type matrix [][]cell

func (board matrix) SetButtonLog(content *fyne.Container, nMines int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			r, c := i, j
			board[r][c].button.OnTapped = func() {
				if firstClick == false {
					firstClick = true
					extremes := getExtremes(board, r, c)
					board.Fill(nMines, extremes)
				}
				over := board.onClick(r, c)
				if over {
					time.Sleep(time.Second * 5)
					board.Empty()
					firstClick = false
				}
			}
			content.Add(board[r][c].button)
		}
	}
}

func (board matrix) onClick(r, c int) bool {
	if board[r][c].isMined {
		for i, row := range board {
			for j, cell := range row {
				if cell.isMined {
					board[i][j].isClicked = true
					board[i][j].button.Text = "   "
					board[i][j].button.Importance = widget.DangerImportance
					board[i][j].button.Refresh()
				} else {
					recursiveClicking(board, i, j)
				}
			}
		}
		return true
	}
	recursiveClicking(board, r, c)
	return false
}

// if cell has 0 neighbors, function clicks them, recursive process
func recursiveClicking(board matrix, r, c int) {
	if fMinedNeigh(board, r, c) == 0 && board[r][c].isClicked == false {
		board[r][c].isClicked = true
		neigh := getExtremes(board, r, c)
		for row := neigh[0]; row <= neigh[2]; row++ {
			for col := neigh[1]; col <= neigh[3]; col++ {
				recursiveClicking(board, row, col)
			}
		}
	}

	if board[r][c].numMinNeigh == 0 {
		board[r][c].button.Importance = widget.LowImportance
		board[r][c].button.Disable()
	} else if board[r][c].numMinNeigh < 3 {
		board[r][c].button.Importance = widget.HighImportance
		board[r][c].button.Text = strconv.Itoa(board[r][c].numMinNeigh)
	} else {
		board[r][c].button.Importance = widget.WarningImportance
		board[r][c].button.Text = strconv.Itoa(board[r][c].numMinNeigh)
	}
	board[r][c].button.Refresh()
}

// calculates number of mined neighbors
func fMinedNeigh(board matrix, r, c int) int {
	n := 0
	neigh := getExtremes(board, r, c)

	for row := neigh[0]; row <= neigh[2]; row++ {
		for col := neigh[1]; col <= neigh[3]; col++ {
			if board[row][col].isMined {
				n++
			}
		}
	}
	return n
}

// returns coordinates (row,column) of left top and right bottom, depending on situation may return itself
func getExtremes(board matrix, r, c int) []int {
	extremes := make([]int, 0, 4)
	if r == 0 {
		extremes = append(extremes, r)
	} else {
		extremes = append(extremes, r-1)
	}
	if c == 0 {
		extremes = append(extremes, c)
	} else {
		extremes = append(extremes, c-1)
	}
	if r == len(board)-1 {
		extremes = append(extremes, r)
	} else {
		extremes = append(extremes, r+1)
	}
	if c == len(board[0])-1 {
		extremes = append(extremes, c)
	} else {
		extremes = append(extremes, c+1)
	}
	return extremes
}

// fills the board with 'mines' number of mines, avoid coordinate in blocked slice, which contains the coordinates of neighbors of first clicked cell
func (board matrix) Fill(mines int, blocked []int) {
	for i := 0; i < mines; i++ {
		var r, c int
		for {
			r, c = rand.Intn(len(board)), rand.Intn(len(board[0]))
			if (r < blocked[0] || r > blocked[2]) || (c < blocked[1] || c > blocked[3]) {
				if board[r][c].isMined == false {
					break
				}
			}
		}
		board[r][c].isMined = true
	}
	for i := 0; i < cap(board); i++ {
		for j := 0; j < cap(board[i]); j++ {
			if board[i][j].isMined {
				board[i][j].numMinNeigh = -1
				continue
			}
			board[i][j].numMinNeigh = fMinedNeigh(board, i, j)
		}
	}
}

// resets all the config of board to start game again
func (board matrix) Empty() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j].isClicked = false
			board[i][j].isMined = false
			board[i][j].numMinNeigh = 0
			board[i][j].button.Text = "   "
			board[i][j].button.Importance = widget.MediumImportance
			board[i][j].button.Enable()
			board[i][j].button.Refresh()
		}
	}
}

// returns empty board m to n
func EmptyBoard(m, n int) matrix {
	board := make(matrix, 0, m)
	for i := 0; i < cap(board); i++ {
		line := make([]cell, 0, n)
		for j := 0; j < cap(line); j++ {
			line = append(line, cell{button: widget.NewButton("     ", func() {})})
		}
		board = append(board, line)
	}
	return board
}
