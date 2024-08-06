package main

import (
	"fmt"
	"testing"
)

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
