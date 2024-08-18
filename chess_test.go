package main

import (
	"slices"
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
			{_d, _3},
			{_e, _3},
			{_f, _3},
			{_d, _4},
			{_f, _4},
			{_d, _5},
			{_e, _5},
			{_f, _5},
		}
		for _, to := range tos {
			legal := board.validateMove(from, to)
			if !legal {
				t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		{_e, _2},
		{_e, _6},
		{_c, _4},
		{_g, _4},
		{_g, _6},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
		}
	}
}

func TestValidateMoveRookLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wrook
	from := square{_e, _4}
	tos := []square{
		{_a, _4},
		{_b, _4},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
		{_e, _1},
		{_e, _2},
		{_e, _3},
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
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
		{_e, _8},
		{_e, _1},
		{_a, _4},
		{_h, _4},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
		}
	}
}

func TestValidateMoveBishopLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wbishop
	from := square{_e, _4}
	tos := []square{
		{_b, _1},
		{_c, _2},
		{_d, _3},
		{_a, _8},
		{_b, _7},
		{_c, _6},
		{_d, _5},
		{_f, _5},
		{_g, _6},
		{_h, _7},
		{_f, _3},
		{_g, _2},
		{_h, _1},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
	}
}

func TestValidateMoveBishopIllegalObstructed(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wbishop
	board[_b][_7] = wknight
	board[_c][_2] = wknight
	board[_g][_6] = bbishop
	board[_g][_2] = brook
	from := square{_e, _4}
	tos := []square{
		{_a, _8},
		{_b, _1},
		{_h, _7},
		{_h, _1},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
		}
	}
}

func TestValidateMoveQueenLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wqueen
	from := square{_e, _4}
	tos := []square{
		{_a, _4},
		{_b, _4},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
		{_e, _1},
		{_e, _2},
		{_e, _3},
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
		{_b, _1},
		{_c, _2},
		{_d, _3},
		{_a, _8},
		{_b, _7},
		{_c, _6},
		{_d, _5},
		{_f, _5},
		{_g, _6},
		{_h, _7},
		{_f, _3},
		{_g, _2},
		{_h, _1},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
	}
}

func TestValidateMoveQueenIllegalObstructed(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wqueen
	board[_b][_7] = wknight
	board[_c][_2] = wknight
	board[_g][_6] = bbishop
	board[_g][_2] = brook
	board[_e][_7] = wknight
	board[_e][_2] = wknight
	board[_b][_4] = bbishop
	board[_g][_4] = brook
	from := square{_e, _4}
	tos := []square{
		{_a, _8},
		{_b, _1},
		{_h, _7},
		{_h, _1},
		{_e, _8},
		{_e, _1},
		{_a, _4},
		{_h, _4},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
		}
	}
}

func TestValidateMoveKnightLegalEmpty(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wknight
	from := square{_e, _4}
	tos := []square{
		{_c, _3},
		{_c, _5},
		{_d, _2},
		{_d, _6},
		{_f, _2},
		{_f, _6},
		{_g, _3},
		{_g, _5},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
	}
}

func TestValidateMoveWhitePawnLegalStartingPos(t *testing.T) {
	var board board
	board.clear()
	board[_e][_2] = wpawn
	from := square{_e, _2}
	tos := []square{
		{_e, _3},
		{_e, _4},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		{_d, _5},
		{_f, _5},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
		}
	}
}

func TestValidateMoveWhitePawnIllegal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_2] = wpawn
	from := square{_e, _2}
	tos := []square{
		{_e, _1},
		{_e, _5},
		{_d, _2},
		{_d, _3},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
		}
	}
}

func TestValidateMoveBlackPawnLegalStartingPos(t *testing.T) {
	var board board
	board.clear()
	board[_e][_7] = bpawn
	from := square{_e, _7}
	tos := []square{
		{_e, _6},
		{_e, _5},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
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
		{_d, _4},
		{_f, _4},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if !legal {
			t.Errorf("Move from %v to %v should be legal but is illegal.", from, to)
		}
	}
}

func TestValidateMoveBlackPawnIllegal(t *testing.T) {
	var board board
	board.clear()
	board[_e][_7] = bpawn
	from := square{_e, _7}
	tos := []square{
		{_e, _8},
		{_e, _4},
		{_d, _7},
		{_d, _6},
	}
	for _, to := range tos {
		legal := board.validateMove(from, to)
		if legal {
			t.Errorf("Move from %v to %v should be illegal but is legal.", from, to)
		}
	}
}

func TestSquareAttackedByPlayerQueenWhite(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := white
	attackers := []square{
		{_a, _4},
		{_b, _4},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
		{_e, _1},
		{_e, _2},
		{_e, _3},
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
		{_b, _1},
		{_c, _2},
		{_d, _3},
		{_a, _8},
		{_b, _7},
		{_c, _6},
		{_d, _5},
		{_f, _5},
		{_g, _6},
		{_h, _7},
		{_f, _3},
		{_g, _2},
		{_h, _1},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = wqueen
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerQueenBlack(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := black
	attackers := []square{
		{_a, _4},
		{_b, _4},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
		{_e, _1},
		{_e, _2},
		{_e, _3},
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
		{_b, _1},
		{_c, _2},
		{_d, _3},
		{_a, _8},
		{_b, _7},
		{_c, _6},
		{_d, _5},
		{_f, _5},
		{_g, _6},
		{_h, _7},
		{_f, _3},
		{_g, _2},
		{_h, _1},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = bqueen
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerRookWhite(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := white
	attackers := []square{
		{_a, _4},
		{_b, _4},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
		{_e, _1},
		{_e, _2},
		{_e, _3},
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = wrook
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerRookBlack(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := black
	attackers := []square{
		{_a, _4},
		{_b, _4},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
		{_e, _1},
		{_e, _2},
		{_e, _3},
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = brook
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerRookFalseDiagonal(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := black
	attackers := []square{
		{_b, _1},
		{_c, _2},
		{_d, _3},
		{_a, _8},
		{_b, _7},
		{_c, _6},
		{_d, _5},
		{_f, _5},
		{_g, _6},
		{_h, _7},
		{_f, _3},
		{_g, _2},
		{_h, _1},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = brook
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if attacked {
			t.Errorf("Square %v should not be attacked by %v but is.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerBishopWhite(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := white
	attackers := []square{
		{_b, _1},
		{_c, _2},
		{_d, _3},
		{_a, _8},
		{_b, _7},
		{_c, _6},
		{_d, _5},
		{_f, _5},
		{_g, _6},
		{_h, _7},
		{_f, _3},
		{_g, _2},
		{_h, _1},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = wbishop
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerBishopBlack(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := black
	attackers := []square{
		{_b, _1},
		{_c, _2},
		{_d, _3},
		{_a, _8},
		{_b, _7},
		{_c, _6},
		{_d, _5},
		{_f, _5},
		{_g, _6},
		{_h, _7},
		{_f, _3},
		{_g, _2},
		{_h, _1},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = bbishop
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerBishopFalseOrthogonal(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := black
	attackers := []square{
		{_a, _4},
		{_b, _4},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
		{_e, _1},
		{_e, _2},
		{_e, _3},
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = bbishop
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if attacked {
			t.Errorf("Square %v should not be attacked by %v but is.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerBlackObstructed(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := black
	board[_a][_8] = bqueen
	board[_a][_4] = bqueen
	board[_e][_8] = brook
	board[_h][_1] = bbishop
	board[_d][_4] = bknight
	board[_d][_5] = wknight
	board[_e][_5] = wknight
	board[_f][_3] = bknight
	attacked := board.squareAttackedByPlayer(sq, player)
	if attacked {
		t.Errorf("Square %v should not be attacked but is.", sq)
	}
}

func TestSquareAttackedByPlayerPawnWhite(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := white
	attackers := []square{
		{_d, _3},
		{_f, _3},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = wpawn
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerPawnBlack(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := black
	attackers := []square{
		{_d, _5},
		{_f, _5},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = bpawn
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerPawnFalse(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := white
	attackers := []square{
		{_e, _3},
		{_e, _2},
		{_d, _5},
		{_f, _5},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = wpawn
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if attacked {
			t.Errorf("Square %v should not be attacked by %v but is.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerKnightWhite(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := white
	attackers := []square{
		{_d, _2},
		{_f, _2},
		{_c, _3},
		{_g, _3},
		{_c, _5},
		{_g, _5},
		{_d, _6},
		{_f, _6},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = wknight
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestSquareAttackedByPlayerKnightBlack(t *testing.T) {
	var board board
	board.clear()
	sq := square{_e, _4}
	player := black
	attackers := []square{
		{_d, _2},
		{_f, _2},
		{_c, _3},
		{_g, _3},
		{_c, _5},
		{_g, _5},
		{_d, _6},
		{_f, _6},
	}
	for _, attacker := range attackers {
		board[attacker.file][attacker.row] = bknight
		attacked := board.squareAttackedByPlayer(sq, player)
		board[attacker.file][attacker.row] = empty
		if !attacked {
			t.Errorf("Square %v should be attacked by %v but is not.", sq, attacker)
		}
	}
}

func TestGenerateValidMovesKing(t *testing.T) {
	var pos position
	pos.board.clear()
	pos.board[_e][_4] = wking
	pos.turn = white
	want := []square{
		{_d, _3},
		{_d, _4},
		{_d, _5},
		{_e, _3},
		{_e, _5},
		{_f, _3},
		{_f, _4},
		{_f, _5},
	}
	got := pos.generateValidMoves()
	from := square{_e, _4}
	if !equivalent(want, got[from]) {
		t.Errorf("Generated moves are wrong. Want %v but got %v.", want, got)
	}
}

func TestGenerateValidMovesKingNoMoves(t *testing.T) {
	var pos position
	pos.board.clear()
	pos.board[_e][_4] = wking
	pos.board[_f][_3] = wpawn
	pos.board[_e][_6] = bking
	pos.board[_e][_3] = bbishop
	pos.board[_d][_1] = brook
	pos.board[_f][_4] = bpawn
	pos.turn = white
	want := []square{}
	got := pos.generateValidMoves()
	from := square{_e, _4}
	if !equivalent(want, got[from]) {
		t.Errorf("Generated moves are wrong. Want %v but got %v.", want, got)
	}
}

func TestGenerateValidMovesRook(t *testing.T) {
	var pos position
	pos.board.clear()
	pos.board[_e][_4] = wrook
	pos.turn = white
	want := []square{
		{_e, _1},
		{_e, _2},
		{_e, _3},
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
		{_a, _4},
		{_b, _4},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
	}
	got := pos.generateValidMoves()
	from := square{_e, _4}
	if !equivalent(want, got[from]) {
		t.Errorf("Generated moves are wrong. Want %v but got %v.", want, got)
	}
}

func TestGenerateValidMovesRookObstructed(t *testing.T) {
	var pos position
	pos.board.clear()
	pos.board[_e][_4] = wrook
	pos.board[_e][_3] = wpawn
	pos.board[_c][_4] = brook
	pos.turn = white
	want := []square{
		{_e, _5},
		{_e, _6},
		{_e, _7},
		{_e, _8},
		{_c, _4},
		{_d, _4},
		{_f, _4},
		{_g, _4},
		{_h, _4},
	}
	got := pos.generateValidMoves()
	from := square{_e, _4}
	if !equivalent(want, got[from]) {
		t.Errorf("Generated moves are wrong. Want %v but got %v.", want, got)
	}
}

func TestParseMoveOk(t *testing.T) {
	move := "e2-a5"
	gotFrom, gotTo, err := parseMove(move)
	wantFrom, wantTo := square{_e, _2}, square{_a, _5}
	if gotFrom != wantFrom || gotTo != wantTo || err != nil {
		t.Errorf("Move %q is not parsed correctly. From: %v, To: %v, Error: %v", move, gotFrom, gotTo, err)
	}
}

func TestParseMoveError(t *testing.T) {
	moves := []string{
		"e2-i5",
		"e9-e5",
		"xyz",
		"ab-cd",
		"##e2-e5##",
	}
	for _, move := range moves {
		_, _, err := parseMove(move)
		if err == nil {
			t.Errorf("Parsing move %q should error, but does not", move)
		}
	}
}

func TestFindKingOf(t *testing.T) {
	var board board
	board.clear()
	board[_e][_4] = wking
	board[_e][_6] = bking
	_, wErr := board.findKingOf(white)
	_, bErr := board.findKingOf(black)
	if wErr != nil || bErr != nil {
		t.Error(wErr, bErr)
	}
}

func TestFindKingOfErr(t *testing.T) {
	var board board
	board.clear()
	_, err := board.findKingOf(white)
	if err == nil {
		t.Error("Missing king should error but does not")
	}
}

func equivalent(a, b []square) bool {
	if len(a) != len(b) {
		return false
	}
	for _, sq := range a {
		if !slices.Contains(b, sq) {
			return false
		}
	}
	return true
}
