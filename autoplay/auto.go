package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/widget"
)

func (board matrix) clickNotMined() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			//if is clicked and nummined neigh equal to (notclicked - foundMined) click the noclicks
			cell := board[i][j]
			if cell.isClicked && cell.numMinNeigh != 0 {
				neigh := getExtremes(board, i, j)
				fundthamined := fNumMinedNeig(board, neigh)

				if cell.numMinNeigh == (fundthamined) {
					for row := neigh[0]; row <= neigh[2]; row++ {
						for col := neigh[1]; col <= neigh[3]; col++ {
							if board[row][col].isClicked == false && board[row][col].foundMined == false {
								time.Sleep(time.Second / 30)

								if board[row][col].isMined {
									fmt.Println("invalid try", row, col)
									fmt.Println("commande came from", i, j)
									fmt.Println(cell.numMinNeigh, fundthamined)
									for row := neigh[0]; row <= neigh[2]; row++ {
										for col := neigh[1]; col <= neigh[3]; col++ {
											if board[row][col].foundMined {
												fmt.Println("found mined", row, col)
											}
										}
									}
									time.Sleep(time.Second * 100)
								}
								_ = board.onClick(row, col)
								board.findMined()
								board.clickNotMined()

							}
						}
					}
				}
			}
		}
	}
}

// find 100% mined ones and make it yellow
func (board matrix) findMined() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cell := board[i][j]
			//if clicked and it mined neig is same as not clicked neig make the neigh yellow
			if cell.isClicked {
				neigh := getExtremes(board, i, j)
				notClicked := fNotClicked(board, neigh)

				if cell.numMinNeigh == notClicked && notClicked != 0 {
					for row := neigh[0]; row <= neigh[2]; row++ {
						for col := neigh[1]; col <= neigh[3]; col++ {
							if board[row][col].isClicked == false {
								board[row][col].button.Importance = widget.WarningImportance
								board[row][col].foundMined = true
								board[row][col].button.Refresh()
							}
						}
					}
				}
			}
		}
	}
}

func fNotClicked(board matrix, neigh []int) int {
	notClicked := 0
	for row := neigh[0]; row <= neigh[2]; row++ {
		for col := neigh[1]; col <= neigh[3]; col++ {
			if board[row][col].isClicked == false {
				notClicked++
			}
		}
	}
	return notClicked
}

func fNumMinedNeig(board matrix, neigh []int) int {
	foundthatismined := 0
	for row := neigh[0]; row <= neigh[2]; row++ {
		for col := neigh[1]; col <= neigh[3]; col++ {
			if board[row][col].foundMined {
				foundthatismined++
			}
		}
	}
	return foundthatismined
}
