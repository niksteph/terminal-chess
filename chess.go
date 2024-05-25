package main

import (
	"fmt"
)

type piece = rune
const empty piece = 0x0020
const (
	wking piece = 0x2654 + iota
	wqueen
	wrook
	wbishop
	wknight
	wpawn
	bking
	bqueen
	brook
	bbishop
	bknight
	bpawn
)

const (
	_a int = iota
	_b
	_c
	_d
	_e
	_f
	_g
	_h
)
const fileUnicodeOffset = 0x0061

const (
	_1 int = iota
	_2
	_3
	_4
	_5
	_6
	_7
	_8
)
const rowUnicodeOffset = 0x0031

type player string
const (
	white player = "white"
	black player = "black"
)
type board [8][8]piece
type position struct {
	board board
	turn player
}

type square struct {
	file int
	row int
}

func main() {
	var position position
	position.startingPos()
	err := position.move(square{_e,_2}, square{_e,_4})
	fmt.Println(err)
	err = position.move(square{_d,_7}, square{_d,_5})
	fmt.Println(err)
	err = position.move(square{_e,_4}, square{_d,_5})
	fmt.Println(err)
	fmt.Println(position.board.formatb(true))
}

func (position *position) startingPos() {
	var board *board = &(position.board)
	board.clear()
	board[_a][_1] = wrook
	board[_b][_1] = wknight
	board[_c][_1] = wbishop
	board[_d][_1] = wqueen
	board[_e][_1] = wking
	board[_f][_1] = wbishop
	board[_g][_1] = wknight
	board[_h][_1] = wrook

	board[_a][_8] = brook
	board[_b][_8] = bknight
	board[_c][_8] = bbishop
	board[_d][_8] = bqueen
	board[_e][_8] = bking
	board[_f][_8] = bbishop
	board[_g][_8] = bknight
	board[_h][_8] = brook

	for file := _a; file <= _h; file++ {
		board[file][_2] = wpawn
		board[file][_7] = bpawn
	}

	position.turn = white
}

func (board *board) clear() {
	for row := _1; row <= _8; row++ {
		for file := _a; file <= _h; file++ {
			board[file][row] = 0x0020
		}
	}
}

func (board *board) formatb(withLabels bool) (s string) {
	for row := _8; row >= _1; row-- {
		if row < _8 {
			s += "\n"
		}
		if withLabels {
			s += fmt.Sprintf("%d ", row+1)
		}
		s += "\033[38;5;0m"
		for file := _a; file <= _h; file++ {
			if (row+file)%2 == 0 {
				s += fmt.Sprintf("\033[48;5;250m%c ", (*board)[file][row])
			} else {
				s += fmt.Sprintf("\033[48;5;15m%c ", (*board)[file][row])
			}
		}
		s += "\033[0m"
	}
	if withLabels {
		s += "\n  a b c d e f g h"
	}
	return
}

func (position *position) move(from, to square) error {
	var board *board = &(position.board)
	if board[from.file][from.row] ==  empty {
		return fmt.Errorf("Square %c%c is empty!", from.file+fileUnicodeOffset, from.row+rowUnicodeOffset)
	}
	playerOfPiece, err := playerOf(board[from.file][from.row])
	if err != nil {
		return err
	}
	if position.turn != playerOfPiece {
		return fmt.Errorf("Not %v's turn!", playerOfPiece)
	}
	
	board[to.file][to.row] = board[from.file][from.row]
	board[from.file][from.row] = empty
	if position.turn == white {
		position.turn = black
	} else {
		position.turn = white
	}
	return nil
}

func playerOf(piece piece) (player, error) {
	if wking <= piece && piece <= wpawn {
		return white, nil
	} else if bking <= piece && piece <= bpawn {
		return black, nil
	}
	return "", fmt.Errorf("Not an actual piece: %v", piece)
}
