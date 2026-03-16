package chess

import (
	"fmt"
)

const (
	White = iota
	Black
)
const (
	Pawn = iota
	Knight
	Bishop
	Rook
	Queen
	King
	Empty
)

var (
	WhiteKingside  uint8 = 1
	WhiteQueenside uint8 = 2
	BlackKingside  uint8 = 4
	BlackQueenside uint8 = 8
)

const (
	NotAFile  uint64 = 0xfefefefefefefefe
	NotHFile  uint64 = 0x7f7f7f7f7f7f7f7f
	NotABFile uint64 = 0xfcfcfcfcfcfcfcfc
	NotGHFile uint64 = 0x3f3f3f3f3f3f3f3f
)

var pieceChars = [2][6]rune{
	{'P', 'N', 'B', 'R', 'Q', 'K'},
	{'p', 'n', 'b', 'r', 'q', 'k'},
}
var pieceInfo = map[rune][2]int{
	'P': {0, 0},
	'N': {0, 1},
	'B': {0, 2},
	'R': {0, 3},
	'Q': {0, 4},
	'K': {0, 5},
	'p': {1, 0},
	'n': {1, 1},
	'b': {1, 2},
	'r': {1, 3},
	'q': {1, 4},
	'k': {1, 5},
}

type Board struct {
	Colors          [2]uint64
	Pieces          [6]uint64
	SideToMove      uint8
	CastlingRights  uint8
	EnPassantSquare uint8
	HalfMoveClock   uint8
	FullMoveNumber  uint16
}

func init() {
	for sq := 0; sq < 64; sq++ {
		KnightAttacks[sq] = maskKnightAttacks(uint8(sq))
		KingAttacks[sq] = maskKingAttacks(uint8(sq))
		PawnAttacks[White][sq] = maskPawnAttacks(White, uint8(sq))
		PawnAttacks[Black][sq] = maskPawnAttacks(Black, uint8(sq))
	}
	initSliders()
}
func (board *Board) Print() {
	pieceChars := [2][6]rune{
		{'P', 'N', 'B', 'R', 'Q', 'K'},
		{'p', 'n', 'b', 'r', 'q', 'k'},
	}
	fmt.Println()
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d ", rank+1)
		for file := 0; file < 8; file++ {
			square := uint8(rank*8 + file)
			char := '.'
			for color := White; color <= Black; color++ {
				if GetBit(board.Colors[color], square) != 0 {
					for piece := Pawn; piece <= King; piece++ {
						if GetBit(board.Pieces[piece], square) != 0 {
							char = pieceChars[color][piece]
							break
						}
					}
					break
				}
			}
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
	fmt.Println("\n  a b c d e f g h")
}
func (board *Board) GeneratePseudoLegalMoves() MoveList {
	ml := MoveList{}
	board.generatePawnMoves(&ml)
	board.generateKnightMoves(&ml)
	board.generateKingMoves(&ml)
	board.generateSliderMoves(&ml, Bishop)
	board.generateSliderMoves(&ml, Rook)
	board.generateSliderMoves(&ml, Queen)

	return ml
}
func (board *Board) IsSquareAttacked(sq uint8, attackerColor uint8) bool {
	var pawn uint64
	if attackerColor == White {
		pawn = PawnAttacks[Black][sq] & board.Colors[attackerColor] & board.Pieces[Pawn]
	}
	if attackerColor == Black {
		pawn = PawnAttacks[White][sq] & board.Colors[attackerColor] & board.Pieces[Pawn]
	}
	if pawn != 0 {
		return true
	}
	knight := KnightAttacks[sq] & board.Colors[attackerColor] & board.Pieces[Knight]
	if knight != 0 {
		return true
	}
	king := KingAttacks[sq] & board.Colors[attackerColor] & board.Pieces[King]
	if king != 0 {
		return true
	}
	occupied := board.Colors[White] | board.Colors[Black]
	bishop := GetBishopAttacks(sq, occupied) & (board.Pieces[Bishop] | board.Pieces[Queen]) & board.Colors[attackerColor]
	if bishop != 0 {
		return true
	}
	rook := GetRookAttacks(sq, occupied) & (board.Pieces[Rook] | board.Pieces[Queen]) & board.Colors[attackerColor]
	if rook != 0 {
		return true
	}
	return false
}

// 0000	0	Quiet move (Default)
// 0001	1	Double pawn push
// 0010	2	King-side castle
// 0011	3	Queen-side castle
// 0100	4	Standard capture
// 0101	5	En Passant capture
// 1000	8	Knight promotion
// 1001	9	Bishop promotion
// 1010	10	Rook promotion
// 1011	11	Queen promotion
// 1100	12	Knight promotion + capture
// 1101	13	Bishop promotion + capture
// 1110	14	Rook promotion + capture
// 1111	15	Queen promotion + capture
