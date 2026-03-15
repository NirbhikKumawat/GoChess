package chess

var KingAttacks [64]uint64

func maskKingAttacks(sq uint8) uint64 {
	var bitboard uint64 = 0
	var attacks uint64 = 0
	SetBit(&bitboard, sq)
	//N
	attacks |= bitboard << 8
	//S
	attacks |= bitboard >> 8
	//E
	if (bitboard & NotHFile) != 0 {
		attacks |= bitboard << 1
	}
	//W
	if (bitboard & NotAFile) != 0 {
		attacks |= bitboard >> 1
	}
	//NW
	if (bitboard & NotAFile) != 0 {
		attacks |= bitboard << 7
	}
	//NE
	if (bitboard & NotHFile) != 0 {
		attacks |= bitboard << 9
	}
	//SE
	if (bitboard & NotHFile) != 0 {
		attacks |= bitboard >> 7
	}
	//SW
	if (bitboard & NotAFile) != 0 {
		attacks |= bitboard >> 9
	}
	return attacks
}
