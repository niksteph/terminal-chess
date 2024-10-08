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
	owner, isEmpty := playerOf(board[from.file][from.row])
	if isEmpty {
		return fmt.Errorf("Square %v is empty!", from)
	}
	if position.turn != owner {
		return fmt.Errorf("Not %v's turn!", owner)
	}
	if !board.validateMove(from, to) {
		return fmt.Errorf("Invalid move from %v to %v!", from, to)
	}

	position.turn = !position.turn
	return nil
}

func (board *board) validateMove(from, to square) bool {
	fromPlayer, isEmpty := playerOf(board[from.file][from.row])
	if isEmpty {
		return false
	}
	if toPlayer, isEmpty := playerOf(board[to.file][to.row]); !isEmpty && fromPlayer == toPlayer {
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
			panic(fmt.Sprintf("Impossible pawn position: %v", from))
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
		}
		if fileDiff > 1 {
			return false
		} else if fileDiff == 0 && !(1 <= moveRow && ((from.row == startRow && moveRow <= 2) || moveRow == 1)) {
			return false
		} else if toPlayer, isEmpty := playerOf(board[to.file][to.row]); fileDiff == 1 &&
			!(moveRow == 1 && !isEmpty && toPlayer == opponent) {
			return false
		}
	}
	return true
}

func (board *board) squareAttackedByPlayer(sq square, attacker player) bool {
	orthogonals := []square{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	diagonals := []square{
		{1, 1},
		{1, -1},
		{-1, -1},
		{-1, 1},
	}
	for _, d := range orthogonals {
		file, row := sq.file+d.file, sq.row+d.row
		if file < _a || _h < file || row < _1 || _8 < row {
			continue
		}
		piece := board[file][row]
		if ((piece == wqueen || piece == wrook || piece == wking) && attacker == white) ||
			((piece == bqueen || piece == brook || piece == bking) && attacker == black) {
			return true
		} else if piece != empty {
			continue
		}
		for i := 2; i <= 7; i++ {
			file, row = sq.file+d.file*i, sq.row+d.row*i
			if file < _a || _h < file || row < _1 || _8 < row {
				break
			}
			piece = board[file][row]
			if ((piece == wqueen || piece == wrook) && attacker == white) ||
				((piece == bqueen || piece == brook) && attacker == black) {
				return true
			} else if piece != empty {
				break
			}
		}
	}
	for _, d := range diagonals {
		file, row := sq.file+d.file, sq.row+d.row
		if file < _a || _h < file || row < _1 || _8 < row {
			continue
		}
		piece := board[file][row]
		if ((piece == wqueen || piece == wbishop || piece == wking) && attacker == white) ||
			((piece == bqueen || piece == bbishop || piece == bking) && attacker == black) {
			return true
		} else if piece != empty {
			continue
		}
		for i := 2; i <= 7; i++ {
			file, row = sq.file+d.file*i, sq.row+d.row*i
			if file < _a || _h < file || row < _1 || _8 < row {
				break
			}
			piece = board[file][row]
			if ((piece == wqueen || piece == wbishop) && attacker == white) ||
				((piece == bqueen || piece == bbishop) && attacker == black) {
				return true
			} else if piece != empty {
				break
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

	movFile := 1
	movRow := 2
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			file := sq.file + movFile
			row := sq.row + movRow
			if (_a <= file && file <= _h) && (_1 <= row && row <= _8) {
				piece := board[file][row]
				if (piece == wknight && attacker == white) || (piece == bknight && attacker == black) {
					return true
				}
			}
			tmp := -movFile
			movFile = movRow
			movRow = tmp
		}
		tmp := movFile
		movFile = movRow
		movRow = tmp
	}
	return false
}

func (position *position) generateValidMoves() (moves map[square][]square) {
	moves = make(map[square][]square)
	player := position.turn
	board := position.board
	orthogonals := []square{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	diagonals := []square{
		{1, 1},
		{1, -1},
		{-1, -1},
		{-1, 1},
	}
	for file := range board {
		for row, piece := range board[file] {
			if owner, isEmpty := playerOf(piece); !isEmpty && owner == player {
				from := square{file, row}
				if piece == wking || piece == bking {
					for _, d := range orthogonals {
						to := square{from.file + d.file, from.row + d.row}
						if !withinBounds(to) {
							continue
						}
						isCheckedAfter, _ := board.kingIsCheckedAfter(from, to)
						if owner, _ := playerOf(board[to.file][to.row]); (board[to.file][to.row] == empty ||
							owner != player) &&
							!isCheckedAfter {

							_, ok := moves[from]
							if !ok {
								moves[from] = []square{to}
							} else {
								moves[from] = append(moves[from], to)
							}
						}
					}
					for _, d := range diagonals {
						to := square{from.file + d.file, from.row + d.row}
						if !withinBounds(to) {
							continue
						}
						isCheckedAfter, _ := board.kingIsCheckedAfter(from, to)
						if owner, _ := playerOf(board[to.file][to.row]); (board[to.file][to.row] == empty ||
							owner != player) &&
							!isCheckedAfter {
							_, ok := moves[from]
							if !ok {
								moves[from] = []square{to}
							} else {
								moves[from] = append(moves[from], to)
							}
						}
					}
				}
				if piece == wrook || piece == brook || piece == wqueen || piece == bqueen {
					for _, d := range orthogonals {
						for i := 1; i <= 7; i++ {
							to := square{from.file + d.file*i, from.row + d.row*i}
							if !withinBounds(to) {
								break
							}
							isCheckedAfter, _ := board.kingIsCheckedAfter(from, to)
							owner, isEmpty := playerOf(board[to.file][to.row])
							if isEmpty && !isCheckedAfter {
								_, ok := moves[from]
								if !ok {
									moves[from] = []square{to}
								} else {
									moves[from] = append(moves[from], to)
								}
							} else if !isEmpty && owner != player && !isCheckedAfter {
								_, ok := moves[from]
								if !ok {
									moves[from] = []square{to}
								} else {
									moves[from] = append(moves[from], to)
								}
								break
							} else if !isEmpty && owner == player {
								break
							}
						}
					}
				}
				if piece == wbishop || piece == bbishop || piece == wqueen || piece == bqueen {
					for _, d := range diagonals {
						for i := 1; i <= 7; i++ {
							to := square{from.file + d.file*i, from.row + d.row*i}
							if !withinBounds(to) {
								break
							}
							isCheckedAfter, _ := board.kingIsCheckedAfter(from, to)
							owner, isEmpty := playerOf(board[to.file][to.row])
							if isEmpty && !isCheckedAfter {
								_, ok := moves[from]
								if !ok {
									moves[from] = []square{to}
								} else {
									moves[from] = append(moves[from], to)
								}
							} else if !isEmpty && owner != player && !isCheckedAfter {
								_, ok := moves[from]
								if !ok {
									moves[from] = []square{to}
								} else {
									moves[from] = append(moves[from], to)
								}
								break
							} else if !isEmpty && owner == player {
								break
							}
						}
					}
				}
				if piece == wknight || piece == bknight {
					movFile := 1
					movRow := 2
					for i := 0; i < 2; i++ {
						for j := 0; j < 4; j++ {
							to := square{from.file + movFile, from.row + movRow}
							if !withinBounds(to) {
								tmp := -movFile
								movFile = movRow
								movRow = tmp
								continue
							}
							isCheckedAfter, _ := board.kingIsCheckedAfter(from, to)
							if owner, isEmpty := playerOf(board[to.file][to.row]); (isEmpty ||
								owner != player) && !isCheckedAfter {
								_, ok := moves[from]
								if !ok {
									moves[from] = []square{to}
								} else {
									moves[from] = append(moves[from], to)
								}
							}
							tmp := -movFile
							movFile = movRow
							movRow = tmp
						}
						tmp := movFile
						movFile = movRow
						movRow = tmp
					}
				}
				if piece == wpawn || piece == bpawn {
					if from.row == _1 || from.row == _8 {
						panic(fmt.Sprintf("Impossible pawn position: %v", from))
					}
					var dir, startRow int
					if piece == wpawn {
						dir = 1
						startRow = _2
					} else if piece == bpawn {
						dir = -1
						startRow = _7
					}
					to := square{from.file, from.row + dir}
					isCheckedAfter, _ := board.kingIsCheckedAfter(from, to)
					if board[to.file][to.row] == empty && !isCheckedAfter {
						_, ok := moves[from]
						if !ok {
							moves[from] = []square{to}
						} else {
							moves[from] = append(moves[from], to)
						}
						to = square{to.file, to.row + dir}
						isCheckedAfter, _ = board.kingIsCheckedAfter(from, to)
						if board[to.file][to.row] == empty && from.row == startRow && !isCheckedAfter {
							moves[from] = append(moves[from], to)
						}
					}
					to = square{from.file + 1, from.row + dir}
					isCheckedAfter, _ = board.kingIsCheckedAfter(from, to)
					if to.file <= _h {
						if owner, isEmpty := playerOf(board[to.file][to.row]); !isEmpty &&
							owner != position.turn && !isCheckedAfter {
							_, ok := moves[from]
							if !ok {
								moves[from] = []square{to}
							} else {
								moves[from] = append(moves[from], to)
							}
						}
					}
					to = square{from.file - 1, from.row + dir}
					isCheckedAfter, _ = board.kingIsCheckedAfter(from, to)
					if _a <= to.file {
						if owner, isEmpty := playerOf(board[to.file][to.row]); !isEmpty &&
							owner != position.turn && !isCheckedAfter {
							_, ok := moves[from]
							if !ok {
								moves[from] = []square{to}
							} else {
								moves[from] = append(moves[from], to)
							}
						}
					}
				}
			}
		}
	}
	return
}

func (board *board) kingIsCheckedAfter(from, to square) (bool, error) {
	player, isEmpty := playerOf(board[from.file][from.row])
	if isEmpty {
		return false, fmt.Errorf("Square %v is empty.", from)
	}

	var kingPos square
	if board[from.file][from.row] == wking || board[from.file][from.row] == bking {
		kingPos = to
	} else {
		var err error
		kingPos, err = board.findKingOf(player)
		if err != nil {
			return false, err
		}
	}
	tmpPiece := board[to.file][to.row]
	board[to.file][to.row] = board[from.file][from.row]
	board[from.file][from.row] = empty

	isChecked := board.squareAttackedByPlayer(kingPos, !player)

	board[from.file][from.row] = board[to.file][to.row]
	board[to.file][to.row] = tmpPiece
	return isChecked, nil
}

func withinBounds(sq square) bool {
	return _a <= sq.file && sq.file <= _h && _1 <= sq.row && sq.row <= _8
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func playerOf(piece piece) (player, bool) {
	if piece == empty {
		return white, true
	}
	if wking <= piece && piece <= wpawn {
		return white, false
	}
	if bking <= piece && piece <= bpawn {
		return black, false
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

func (sq square) String() string {
	return fmt.Sprintf("%c%c", sq.file+fileUnicodeOffset, sq.row+rowUnicodeOffset)
}
