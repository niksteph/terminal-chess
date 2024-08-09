package main

import (
	"testing"
)

func TestValidateMoveTargetSameColor(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wking
	board[_e][_5] = wqueen
	from := square{_e, _4}
	to := square{_e, _5}
	legal := board.validateMove(from, to)
	if legal {
		t.Error("Move to a square occupied by a piece of the same color should be illegal but is legal.")
	}
}

func TestValidateMoveKingLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	pieces := []piece{wking, bking}
	for _, piece := range pieces {
		board[_e][_4] = piece
		from := square{_e, _4}
		tos := []square{
			square{_d, _3},
			square{_e, _3},
			square{_f, _3},
			square{_d, _4},
			square{_f, _4},
			square{_d, _5},
			square{_e, _5},
			square{_f, _5},
		}
		for _, to := range tos {
			legal := board.validateMove(from, to)
			if !legal {
				t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
					from.file+fileUnicodeOffset,
					from.row+rowUnicodeOffset,
					to.file+fileUnicodeOffset,
					to.row+rowUnicodeOffset)
			}
		}
	}
}

func TestValidateMoveKingIllegalDistance(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wking
	from := square{_e, _4}
	tos := []square{
		square{_e, _2},
		square{_e, _6},
		square{_c, _4},
		square{_g, _4},
		square{_g, _6},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveRookLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wrook
	from := square{_e, _4}
	tos := []square{
		square{_a, _4},
		square{_b, _4},
		square{_c, _4},
		square{_d, _4},
		square{_f, _4},
		square{_g, _4},
		square{_h, _4},
		square{_e, _1},
		square{_e, _2},
		square{_e, _3},
		square{_e, _5},
		square{_e, _6},
		square{_e, _7},
		square{_e, _8},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveRookIllegalDiagonal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wrook
	from := square{_e, _4}
	to := square{_f, _5}
	legal := board.validateMove(from, to)
	if legal {
		t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}
}

func TestValidateMoveRookIllegalObstructed(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wrook
	board[_e][_7] = wknight
	board[_e][_2] = wknight
	board[_b][_4] = bbishop
	board[_g][_4] = brook
	from := square{_e, _4}
	tos := []square{
		square{_e, _8},
		square{_e, _1},
		square{_a, _4},
		square{_h, _4},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveBishopLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wbishop
	from := square{_e, _4}
	tos := []square{
		square{_b, _1},
		square{_c, _2},
		square{_d, _3},
		square{_a, _8},
		square{_b, _7},
		square{_c, _6},
		square{_d, _5},
		square{_f, _5},
		square{_g, _6},
		square{_h, _7},
		square{_f, _3},
		square{_g, _2},
		square{_h, _1},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveBishopIllegalOrthogonal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wbishop
	from := square{_e, _4}
	to := square{_e, _3}
	legal := board.validateMove(from, to)
	if legal {
		t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}
}

func TestValidateMoveQueenLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wqueen
	from := square{_e, _4}
	tos := []square{
		square{_a, _4},
		square{_b, _4},
		square{_c, _4},
		square{_d, _4},
		square{_f, _4},
		square{_g, _4},
		square{_h, _4},
		square{_e, _1},
		square{_e, _2},
		square{_e, _3},
		square{_e, _5},
		square{_e, _6},
		square{_e, _7},
		square{_e, _8},
		square{_b, _1},
		square{_c, _2},
		square{_d, _3},
		square{_a, _8},
		square{_b, _7},
		square{_c, _6},
		square{_d, _5},
		square{_f, _5},
		square{_g, _6},
		square{_h, _7},
		square{_f, _3},
		square{_g, _2},
		square{_h, _1},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveQueenIllegal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wqueen
	from := square{_e, _4}
	to := square{_f, _2}
	legal := board.validateMove(from, to)
	if legal {
		t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}
}

func TestValidateMoveKnightLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wknight
	from := square{_e, _4}
	tos := []square{
		square{_c, _3},
		square{_c, _5},
		square{_d, _2},
		square{_d, _6},
		square{_f, _2},
		square{_f, _6},
		square{_g, _3},
		square{_g, _5},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveKnightIllegal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wknight
	from := square{_e, _4}
	to := square{_e, _5}
	legal := board.validateMove(from, to)
	if legal {
		t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}
}

func TestValidateMoveWhitePawnLegalStartingPos(t *testing.T) {
	var board board
	board.clear()
	board[_e][_2] = wpawn
	from := square{_e, _2}
	tos := []square{
		square{_e, _3},
		square{_e, _4},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveWhitePawnLegalStandard(t *testing.T) {
	var board board
	board.clear()
	board[_e][_3] = wpawn
	from := square{_e, _3}
	to := square{_e, _4}
	legal := board.validateMove(from, to)
	if !legal {
		t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}
}

func TestValidateMoveWhitePawnLegalTaking(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wpawn
	board[_d][_5] = bpawn
	board[_f][_5] = bknight
	from := square{_e, _4}
	tos := []square{
		square{_d, _5},
		square{_f, _5},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveWhitePawnIllegal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_2] = wpawn
	from := square{_e, _2}
	tos := []square{
		square{_e, _1},
		square{_e, _5},
		square{_d, _2},
		square{_d, _3},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveBlackPawnLegalStartingPos(t *testing.T) {
	var board board
	board.clear()
	board[_e][_7] = bpawn
	from := square{_e, _7}
	tos := []square{
		square{_e, _6},
		square{_e, _5},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveBlackPawnLegalStandard(t *testing.T) {
	var board board
	board.clear()
	board[_e][_6] = bpawn
	from := square{_e, _6}
	to := square{_e, _5}
	legal := board.validateMove(from, to)
	if !legal {
		t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}
}

func TestValidateMoveBlackPawnLegalTaking(t *testing.T) {
	var board board
	board.clear()
	board[_e][_5] = bpawn
	board[_d][_4] = wpawn
	board[_f][_4] = wknight
	from := square{_e, _5}
	tos := []square{
		square{_d, _4},
		square{_f, _4},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestValidateMoveBlackPawnIllegal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_7] = bpawn
	from := square{_e, _7}
	tos := []square{
		square{_e, _8},
		square{_e, _4},
		square{_d, _7},
		square{_d, _6},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}
