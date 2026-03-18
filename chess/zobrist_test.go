package chess

import (
	"testing"
)

func TestZobristHashing(t *testing.T) {
	InitZobrist()

	fen1 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board1, _ := ParseFEN(fen1)
	board2, _ := ParseFEN(fen1)

	if board1.Hash == 0 {
		t.Fatal("Hash is 0. Did you forget to add board.Hash = board.GenerateHash() to ParseFEN?")
	}
	if board1.Hash != board2.Hash {
		t.Errorf("Determinism failed: Identical boards have different hashes: %x vs %x", board1.Hash, board2.Hash)
	}

	fen3 := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1" // After 1. e4
	board3, _ := ParseFEN(fen3)
	if board1.Hash == board3.Hash {
		t.Error("Collision detected: Different board states produced the same hash!")
	}

	e4Move := NewMove(12, 28, 1)

	board1Copy := *board1
	if !board1Copy.MakeMove(e4Move) {
		t.Fatal("Failed to make move e4")
	}

	if board1Copy.Hash != board3.Hash {
		t.Errorf("Incremental hashing failed!\nStarting FEN: %s\nMove: e2e4\nExpected Hash (from scratch): %x\nGot Hash (incremental):     %x", fen1, board3.Hash, board1Copy.Hash)
	}

	board4, _ := ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
	kd1Move := NewMove(4, 3, 0)
	board4Copy := *board4
	board4Copy.MakeMove(kd1Move)

	board5, _ := ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R2K3R b kq - 1 1")

	if board4Copy.Hash != board5.Hash {
		t.Errorf("Castling rights hash update failed!\nExpected: %x\nGot:      %x", board5.Hash, board4Copy.Hash)
	}
}
