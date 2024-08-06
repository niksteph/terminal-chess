package main

import (
	"fmt"
	"testing"
)

func TestVerifyMoveTargetSameColor(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wking
	board[_e][_5] = wqueen
	from := square{_e, _4}
	to := square{_e, _5}
	legal := board.verifyMove(from, to)
	if legal {
		t.Error("Move to a square occupied by a piece of the same color should be illegal but is legal.")
	}
}

func TestVerifyMoveKingLegalEmpty(t *testing.T) {
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
			legal := board.verifyMove(from, to)
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

func TestVerifyMoveKingIllegalDistance(t *testing.T) {
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
		legal := board.verifyMove(from, to)
		if legal {
			t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestVerifyMoveRookLegalEmpty(t *testing.T) {
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
		legal := board.verifyMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestVerifyMoveRookIllegalDiagonal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wrook
	from := square{_e, _4}
	to := square{_f, _5}
	legal := board.verifyMove(from, to)
	if legal {
		t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}
}

func TestVerifyMoveBishopLegalEmpty(t *testing.T) {
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
		legal := board.verifyMove(from, to)
		if !legal {
			t.Errorf("Move from %c%c to %c%c should be legal but is illegal.",
				from.file+fileUnicodeOffset,
				from.row+rowUnicodeOffset,
				to.file+fileUnicodeOffset,
				to.row+rowUnicodeOffset)
		}
	}
}

func TestVerifyMoveBishopIllegalOrthogonal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wbishop
	from := square{_e, _4}
	to := square{_e, _3}
	legal := board.verifyMove(from, to)
	if legal {
		t.Errorf("Move from %c%c to %c%c should be illegal but is legal.",
			from.file+fileUnicodeOffset,
			from.row+rowUnicodeOffset,
			to.file+fileUnicodeOffset,
			to.row+rowUnicodeOffset)
	}
}
