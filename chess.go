package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
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

type player bool

const (
	white player = true
	black player = false
)

type board [8][8]piece
type position struct {
	board board
	turn  player
}

type square struct {
	file int
	row  int
}

func main() {
	var position position
	position.startingPos()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(position.board.formatb())
	for scanner.Scan() {
		move := scanner.Text()
		from, to, err := parseMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = position.move(from, to)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(position.board.formatb())
	}
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

func (board *board) formatb() (s string) {
	for row := _8; row >= _1; row-- {
		if row < _8 {
			s += "\n"
		}
		s += fmt.Sprintf("%d ", row+1)
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
	s += "\n  a b c d e f g h"
	return
}

func (position *position) move(from, to square) error {
	if from.file < 0 || from.file > 7 {
		return fmt.Errorf("File %v out of bounds.", from.file)
	}
	if from.row < 0 || from.row > 7 {
		return fmt.Errorf("Row %v out of bounds.", from.row)
	}
	if from == to {
		return fmt.Errorf("Target square is same as origin square.")
	}
	var board *board = &(position.board)
	if board[from.file][from.row] == empty {
		return fmt.Errorf("Square %c%c is empty!", from.file+fileUnicodeOffset, from.row+rowUnicodeOffset)
	}
	playerOfPiece := playerOf(board[from.file][from.row])
	if position.turn != playerOfPiece {
		return fmt.Errorf("Not %v's turn!", playerOfPiece)
	}
	if !board.validateMove(from, to) {
		return fmt.Errorf("Invalid move from %c%c to %c%c!",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}

	board[to.file][to.row] = board[from.file][from.row]
	board[from.file][from.row] = empty

	kingPos, err := board.findKingOf(position.turn)
	if err != nil {
		panic(err)
	}
	isChecked := board.squareAttackedByPlayer(kingPos, !position.turn)
	if isChecked {
		board[from.file][from.row] = board[to.file][to.row]
		board[to.file][to.row] = empty
		return fmt.Errorf("Invalid move from %c%c to %c%c, king is left in check!",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}

	position.turn = !position.turn
	return nil
}

func (board *board) validateMove(from, to square) bool {
	if board[to.file][to.row] != empty && playerOf(board[from.file][from.row]) == playerOf(board[to.file][to.row]) {
		return false
	}
	fileDiff := abs(to.file - from.file)
	rowDiff := abs(to.row - from.row)
	switch piece := board[from.file][from.row]; piece {
	case wking, bking:
		if fileDiff > 1 || rowDiff > 1 {
			return false
		}
	case wrook, brook:
		if fileDiff != 0 && rowDiff != 0 {
			return false
		}
		moveFile := to.file - from.file
		moveRow := to.row - from.row
		chebyDist := max(fileDiff, rowDiff)
		moveFile /= chebyDist
		moveRow /= chebyDist
		currFile, currRow := from.file+moveFile, from.row+moveRow
		for currFile != to.file || currRow != to.row {
			if board[currFile][currRow] != empty {
				return false
			}
			currFile, currRow = currFile+moveFile, currRow+moveRow
		}
	case wbishop, bbishop:
		if fileDiff != rowDiff {
			return false
		}
		moveFile := to.file - from.file
		moveRow := to.row - from.row
		chebyDist := max(fileDiff, rowDiff)
		moveFile /= chebyDist
		moveRow /= chebyDist
		currFile, currRow := from.file+moveFile, from.row+moveRow
		for currFile != to.file || currRow != to.row {
			if board[currFile][currRow] != empty {
				return false
			}
			currFile, currRow = currFile+moveFile, currRow+moveRow
		}
	case wqueen, bqueen:
		if fileDiff != 0 && rowDiff != 0 && fileDiff != rowDiff {
			return false
		}
		moveFile := to.file - from.file
		moveRow := to.row - from.row
		chebyDist := max(fileDiff, rowDiff)
		moveFile /= chebyDist
		moveRow /= chebyDist
		currFile, currRow := from.file+moveFile, from.row+moveRow
		for currFile != to.file || currRow != to.row {
			if board[currFile][currRow] != empty {
				return false
			}
			currFile, currRow = currFile+moveFile, currRow+moveRow
		}
	case wknight, bknight:
		if !((fileDiff == 2 && rowDiff == 1) || (fileDiff == 1 && rowDiff == 2)) {
			return false
		}
	case wpawn, bpawn:
		if from.row == _1 || from.row == _8 {
			panic(fmt.Sprintf("Impossible pawn position: %c%c", from.file+fileUnicodeOffset, from.row+rowUnicodeOffset))
		}
		var moveRow, startRow int
		var opponent player
		if piece == wpawn {
			moveRow = to.row - from.row
			startRow = _2
			opponent = black
		} else if piece == bpawn {
			moveRow = from.row - to.row
			startRow = _7
			opponent = white
		} else {
			panic(fmt.Sprintf("Not a pawn: %d (0x%x)", piece, piece))
		}
		if fileDiff > 1 {
			return false
		} else if fileDiff == 0 && !(1 <= moveRow && ((from.row == startRow && moveRow <= 2) || moveRow == 1)) {
			return false
		} else if fileDiff == 1 &&
			!(moveRow == 1 && board[to.file][to.row] != empty && playerOf(board[to.file][to.row]) == opponent) {
			return false
		}
	}
	return true
}

func (board *board) squareAttackedByPlayer(sq square, attacker player) bool {
	orthogonals := []square{
		square{0, 1},
		square{1, 0},
		square{0, -1},
		square{-1, 0},
	}
	diagonals := []square{
		square{1, 1},
		square{1, -1},
		square{-1, -1},
		square{-1, 1},
	}
	for _, d := range orthogonals {
		for i := 1; i <= 7; i++ {
			file, row := sq.file+d.file*i, sq.row+d.row*i
			if _a <= file && file <= _h && _1 <= row && row <= _8 {
				piece := board[file][row]
				if ((piece == wqueen || piece == wrook) && attacker == white) ||
					((piece == bqueen || piece == brook) && attacker == black) {
					return true
				} else if piece != empty {
					break
				}
			}
		}
	}
	for _, d := range diagonals {
		for i := 1; i <= 7; i++ {
			file, row := sq.file+d.file*i, sq.row+d.row*i
			if _a <= file && file <= _h && _1 <= row && row <= _8 {
				piece := board[file][row]
				if ((piece == wqueen || piece == wbishop) && attacker == white) ||
					((piece == bqueen || piece == bbishop) && attacker == black) {
					return true
				} else if piece != empty {
					break
				}
			}
		}
	}
	var row int
	if attacker == white {
		row = sq.row - 1
	} else {
		row = sq.row + 1
	}
	for i := -1; i <= 1; i += 2 {
		file := sq.file + i
		if _a <= file && file <= _h && _1 <= row && row <= _8 {
			piece := board[file][row]
			if (piece == wpawn && attacker == white) || (piece == bpawn && attacker == black) {
				return true
			}
		}
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func playerOf(piece piece) player {
	if wking <= piece && piece <= wpawn {
		return white
	} else if bking <= piece && piece <= bpawn {
		return black
	}
	panic(fmt.Sprintf("Not an actual piece: %d (0x%x)", piece, piece))
}

func parseMove(s string) (from, to square, err error) {
	regex := regexp.MustCompile(`^[a-h][1-8]-[a-h][1-8]$`)
	if !regex.MatchString(s) {
		return square{}, square{}, fmt.Errorf("Move %q does not match format", s)
	}
	squareStrings := strings.Split(s, "-")
	if len(squareStrings) != 2 {
		panic("Size of parsed squares is not 2.")
	}
	squareRunes := make([][]rune, 2)
	squareRunes[0] = []rune(squareStrings[0])
	squareRunes[1] = []rune(squareStrings[1])
	from = square{int(squareRunes[0][0] - fileUnicodeOffset), int(squareRunes[0][1] - rowUnicodeOffset)}
	to = square{int(squareRunes[1][0] - fileUnicodeOffset), int(squareRunes[1][1] - rowUnicodeOffset)}
	return from, to, nil
}

func (board *board) findKingOf(player player) (square, error) {
	for file := range board {
		for row := range board[file] {
			if (player == white && board[file][row] == wking) || (player == black && board[file][row] == bking) {
				return square{file, row}, nil
			}
		}
	}
	return square{0, 0}, fmt.Errorf("%v's king not found.", player)
}

func (p player) String() string {
	if p == white {
		return "white"
	}
	return "black"
}
