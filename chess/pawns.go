package chess

var PawnAttacks [2][64]uint64

func maskPawnAttacks(color uint8, sq uint8) uint64 {
	var bitboard uint64 = 0
	var attacks uint64 = 0
	SetBit(&bitboard, sq)
	if color == White {
		//NE
		if (bitboard & NotHFile) != 0 {
			attacks |= bitboard << 9
		}
		//NW
		if (bitboard & NotAFile) != 0 {
			attacks |= bitboard << 7
		}
	} else {
		//SE
		if (bitboard & NotHFile) != 0 {
			attacks |= bitboard >> 7
		}
		//SW
		if (bitboard & NotAFile) != 0 {
			attacks |= bitboard >> 9
		}
	}
	return attacks
}
