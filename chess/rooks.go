package chess

import "math/bits"

var RookMagics = [64]uint64{
	0x8a80104000800020, 0x140002000100040, 0x2801896a00020011, 0x3908002600000068,
	0x10020220010020, 0x1000880142706, 0x200114000200018, 0x200880a42202000,
	0x100180208010, 0x82001002000040, 0x2008204000c02010, 0x480822000400801,
	0x820104242104, 0x82402011000108, 0x1000282000800011, 0x4002021008210080,
	0x81010080800104, 0x8104040108421, 0x4020800100404, 0x1020802010842,
	0x810010120120400, 0x120800202201004, 0x104041112200100, 0x404010020010008,
	0x810404008010, 0x4208022004010, 0x81440040200101, 0x4210080100021,
	0x420080024040, 0x2804001020200, 0x10810001200881, 0x80010214010081,
	0x48408020801008, 0x81200204010104, 0x201402002082200, 0x200010220110800,
	0x882002010801, 0x808140412108, 0x2211010100808, 0x22108004041000,
	0x84010010801, 0x1021004020208, 0x8004104100404, 0x4041041001010,
	0x240208108010, 0x4420108080002, 0x41002028280, 0x2004011400201,
	0x1110020240, 0x81400024040100, 0x411081042200, 0x8080404010804,
	0x8200041201, 0x408000420001, 0x212008404002, 0x4120021122002,
	0x8404008100408, 0x1800102422080, 0x401000100404, 0x1102080242,
	0x22018804240, 0x2800420124000, 0x800801080001, 0x44021040802,
}

var RookMasks [64]uint64
var RookAttacks [64][4096]uint64

func initSliders() {
	for sq := 0; sq < 64; sq++ {
		BishopMasks[sq] = maskBishopOccupancy(uint8(sq))
		RookMasks[sq] = maskRookOccupancy(uint8(sq))
		bishopRelevantBits := bits.OnesCount64(BishopMasks[sq])
		rookRelevantBits := bits.OnesCount64(RookMasks[sq])
		bishopOccupancyIndices := 1 << bishopRelevantBits
		rookOccupancyIndices := 1 << rookRelevantBits
		for i := 0; i < bishopOccupancyIndices; i++ {
			occupancy := setOccupancy(i, bishopRelevantBits, BishopMasks[sq])
			attacks := bishopAttacksOnTheFly(uint8(sq), occupancy)
			magicIndex := (occupancy * BishopMagics[sq]) >> (64 - bishopRelevantBits)
			BishopAttacks[sq][magicIndex] = attacks
		}
		for i := 0; i < rookOccupancyIndices; i++ {
			occupancy := setOccupancy(i, rookRelevantBits, RookMasks[sq])
			attacks := rookAttacksOnTheFly(uint8(sq), occupancy)
			magicIndex := (occupancy * RookMagics[sq]) >> (64 - rookRelevantBits)
			RookAttacks[sq][magicIndex] = attacks
		}

	}
}

func maskRookOccupancy(sq uint8) uint64 {
	var mask uint64 = 0
	targetRank := int(sq / 8)
	targetFile := int(sq % 8)
	//N
	for r := targetRank + 1; r <= 6; r++ {
		SetBit(&mask, uint8(r*8+targetFile))
	}
	//E
	for r := targetFile + 1; r <= 6; r++ {
		SetBit(&mask, uint8(targetRank*8+r))
	}
	//S
	for r := 1; r < targetRank; r++ {
		SetBit(&mask, uint8(r*8+targetFile))
	}
	//W
	for r := 1; r < targetFile; r++ {
		SetBit(&mask, uint8(targetRank*8+r))
	}
	return mask
}

func setOccupancy(index int, bitsInMask int, attackMask uint64) uint64 {
	var occupancy uint64 = 0
	for count := 0; count < bitsInMask; count++ {
		sq := PopBit(&attackMask)
		if (index & (1 << count)) != 0 {
			occupancy |= 1 << sq
		}
	}
	return occupancy
}

func rookAttacksOnTheFly(sq uint8, block uint64) uint64 {
	var attacks uint64 = 0
	targetRank := int(sq / 8)
	targetFile := int(sq % 8)

	//N
	for r := targetRank + 1; r <= 7; r++ {
		square := uint8(r*8 + targetFile)
		SetBit(&attacks, square)
		if (block & (1 << square)) != 0 {
			break
		}
	}
	//S
	for r := targetRank - 1; r >= 0; r-- {
		square := uint8(r*8 + targetFile)
		SetBit(&attacks, square)
		if (block & (1 << square)) != 0 {
			break
		}
	}
	//E
	for r := targetFile + 1; r <= 7; r++ {
		square := uint8(targetRank*8 + r)
		SetBit(&attacks, square)
		if (block & (1 << square)) != 0 {
			break
		}
	}
	//W
	for r := targetFile - 1; r >= 0; r-- {
		square := uint8(targetRank*8 + r)
		SetBit(&attacks, square)
		if (block & (1 << square)) != 0 {
			break
		}
	}
	return attacks
}

func GetRookAttacks(sq uint8, occupancy uint64) uint64 {
	blockers := occupancy & RookMasks[sq]
	magicIndex := (blockers * RookMagics[sq]) >> (64 - bits.OnesCount64(RookMasks[sq]))
	return RookAttacks[sq][magicIndex]
}
