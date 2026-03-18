package chess

var ZobristPieces [2][6][64]uint64
var ZobristSideToMove uint64
var ZobristCastling [16]uint64
var ZobristEnPassant [8]uint64

func InitZobrist() {
	for color := White; color <= Black; color++ {
		for piece := Pawn; piece <= King; piece++ {
			for sq := 0; sq < 64; sq++ {
				ZobristPieces[color][piece][sq] = randomUint64()
			}
		}
	}
	ZobristSideToMove = randomUint64()
	for i := 0; i < 16; i++ {
		ZobristCastling[i] = randomUint64()
	}
	for i := 0; i < 8; i++ {
		ZobristEnPassant[i] = randomUint64()
	}
}
func (board *Board) GenerateHash() uint64 {
	hash := uint64(0)
	for color := White; color <= Black; color++ {
		for piece := Pawn; piece <= King; piece++ {
			bb := board.Colors[color] & board.Pieces[piece]
			for bb != 0 {
				sq := PopBit(&bb)
				hash ^= ZobristPieces[color][piece][sq]
			}
		}
	}
	if board.SideToMove == Black {
		hash ^= ZobristSideToMove
	}
	hash ^= ZobristCastling[board.CastlingRights]
	if board.EnPassantSquare != 255 {
		file := board.EnPassantSquare % 8
		hash ^= ZobristEnPassant[file]
	}
	return hash
}
