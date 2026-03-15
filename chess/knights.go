package chess

var KnightAttacks [64]uint64

func maskKnightAttacks(sq uint8) uint64 {
	var bitboard uint64 = 0
	var attacks uint64 = 0
	SetBit(&bitboard, sq)
	//NNE
	if (bitboard & NotHFile) != 0 {
		attacks |= bitboard << 17
	}
	//NEE
	if (bitboard & NotGHFile) != 0 {
		attacks |= bitboard << 10
	}
	//NNW
	if (bitboard & NotAFile) != 0 {
		attacks |= bitboard << 15
	}
	//NWW
	if (bitboard & NotABFile) != 0 {
		attacks |= bitboard << 6
	}
	//SSE
	if (bitboard & NotHFile) != 0 {
		attacks |= bitboard >> 15
	}
	//SEE
	if (bitboard & NotGHFile) != 0 {
		attacks |= bitboard >> 6
	}
	//SSW
	if (bitboard & NotAFile) != 0 {
		attacks |= bitboard >> 17
	}
	//SWW
	if (bitboard & NotABFile) != 0 {
		attacks |= bitboard >> 10
	}
	return attacks
}
