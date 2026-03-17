package chess

import "math/bits"

var BishopMagics = [64]uint64{
	0x40040844404084, 0x2004208a004208, 0x10190041080202, 0x108060845042010,
	0x581104180800210, 0x211208044620004, 0x411047368482806, 0x81018051080000,
	0x40028040405102, 0x100410020084000, 0x808020050880, 0x800200100410044,
	0x8000021200084, 0x200215104080200, 0x40022220410020, 0x11021448002020,
	0x11002001020610, 0x401024008204224, 0x82200212040000, 0x8101014028240,
	0x804002201402000, 0x202020220201404, 0x11000508104841, 0x4441014120810,
	0x24240080800021, 0x21008026110020, 0x100020200010040, 0x10082400080082,
	0x24021204481804, 0x20200122114020, 0x10410024440840, 0x8848041011400,
	0x20042080024, 0x1001002202048, 0x240182802200, 0x410220010008,
	0x2100200882081, 0x22108000108, 0x828448880504100, 0x2024100842004,
	0x80041080020420, 0x800820010048, 0x21204120804000, 0x801014022011,
	0x140281001202, 0x81000010424281, 0x401401041000200, 0x88182048082010,
	0x10421041008, 0x180024001142, 0x2014002012024, 0x220108208401,
	0x101440212882, 0x8082024040101, 0x40200210204, 0x101108200080104,
	0x880004101140, 0x20480040204, 0x8804008202010, 0x210242200041200,
	0x48102081028114, 0x8820082020040, 0x4440401084204, 0x410408110082,
}

var BishopMasks [64]uint64
var BishopAttacks [64][512]uint64

func maskBishopOccupancy(sq uint8) uint64 {
	var mask uint64 = 0
	targetRank := int(sq / 8)
	targetFile := int(sq % 8)
	//SE
	for r, f := targetRank-1, targetFile+1; r >= 1 && f <= 6; r, f = r-1, f+1 {
		SetBit(&mask, uint8(r*8+f))
	}
	//NE
	for r, f := targetRank+1, targetFile+1; r <= 6 && f <= 6; r, f = r+1, f+1 {
		SetBit(&mask, uint8(r*8+f))
	}
	//NW
	for r, f := targetRank+1, targetFile-1; r <= 6 && f >= 1; r, f = r+1, f-1 {
		SetBit(&mask, uint8(r*8+f))
	}
	//SW
	for r, f := targetRank-1, targetFile-1; r >= 1 && f >= 1; r, f = r-1, f-1 {
		SetBit(&mask, uint8(r*8+f))
	}
	return mask
}

func bishopAttacksOnTheFly(sq uint8, block uint64) uint64 {
	var attacks uint64 = 0
	targetRank := int(sq / 8)
	targetFile := int(sq % 8)
	//SE
	for r, f := targetRank-1, targetFile+1; r >= 0 && f <= 7; r, f = r-1, f+1 {
		square := uint8(r*8 + f)
		SetBit(&attacks, square)
		if (block & (1 << square)) != 0 {
			break
		}
	}
	//NE
	for r, f := targetRank+1, targetFile+1; r <= 7 && f <= 7; r, f = r+1, f+1 {
		square := uint8(r*8 + f)
		SetBit(&attacks, square)
		if (block & (1 << square)) != 0 {
			break
		}
	}
	//NW
	for r, f := targetRank+1, targetFile-1; r <= 7 && f >= 0; r, f = r+1, f-1 {
		square := uint8(r*8 + f)
		SetBit(&attacks, square)
		if (block & (1 << square)) != 0 {
			break
		}
	}
	//SW
	for r, f := targetRank-1, targetFile-1; r >= 0 && f >= 0; r, f = r-1, f-1 {
		square := uint8(r*8 + f)
		SetBit(&attacks, square)
		if (block & (1 << square)) != 0 {
			break
		}
	}

	return attacks
}

func GetBishopAttacks(sq uint8, occupancy uint64) uint64 {
	return bishopAttacksOnTheFly(sq, occupancy)
	blockers := occupancy & BishopMasks[sq]
	magicIndex := (blockers * BishopMagics[sq]) >> (64 - bits.OnesCount64(BishopMasks[sq]))
	return BishopAttacks[sq][magicIndex]
}
