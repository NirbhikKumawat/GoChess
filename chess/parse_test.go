package chess

import (
	"testing"
)

func TestFENRoundtrip(t *testing.T) {
	fens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",             // Starting pos
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", // Kiwipete
		"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1",                            // Endgame with no castling
	}

	for _, fen := range fens {
		board, err := ParseFEN(fen)
		if err != nil {
			t.Fatalf("Failed to parse FEN: %v", err)
		}

		outFEN := board.GenerateFEN()
		if outFEN != fen {
			t.Errorf("FEN roundtrip failed.\nExpected: %s\nGot:      %s", fen, outFEN)
		}
	}
}

func TestSANRoundtrip(t *testing.T) {
	board, _ := ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")

	moves := board.GeneratePseudoLegalMoves()
	for i := 0; i < moves.Count; i++ {
		m := moves.Moves[i]
		boardCopy := *board

		if !boardCopy.MakeMove(m) {
			continue
		}

		san := board.MoveToSAN(m)
		parsedMove, err := board.ParseSAN(san)

		if err != nil {
			t.Errorf("ParseSAN returned error for %s: %v", san, err)
			continue
		}

		if parsedMove != m {
			t.Errorf("SAN roundtrip failed for %s. Expected internal move %v, got %v", san, m, parsedMove)
		}
	}
}

func TestParsePGN(t *testing.T) {
	pgn := `[Event "FIDE World Cup 2023"]
[Site "Baku AZE"]
[Date "2023.08.22"]
[White "Carlsen, Magnus"]
[Black "Praggnanandhaa, R."]
[Result "1/2-1/2"]

1. e4 e5 2. Nf3 Nc6 3. Bc4 Bc5 4. c3 Nf6 1/2-1/2`

	game, err := ParsePGN(pgn)
	if err != nil {
		t.Fatalf("Failed to parse PGN: %v", err)
	}

	if game.Headers["White"] != "Carlsen, Magnus" {
		t.Errorf("Expected White to be Carlsen, Magnus, got %s", game.Headers["White"])
	}

	if len(game.Moves) != 8 {
		t.Errorf("Expected 8 moves, got %d", len(game.Moves))
	}

	board, _ := ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	firstMoveSAN := board.MoveToSAN(game.Moves[0])
	if firstMoveSAN != "e4" {
		t.Errorf("Expected first move to be e4, got %s", firstMoveSAN)
	}
}
