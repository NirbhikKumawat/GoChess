package chess

const (
	TTExact = iota
	TTAlpha
	TTBeta
)

// TTEntry is a structure to store already calculated best move
type TTEntry struct {
	Hash     uint64
	Depth    int
	Score    int
	Flag     int
	BestMove Move
}

const TTSize = 1000000

// TranspositionTable stores calculated value for a board state
var TranspositionTable [TTSize]TTEntry

// ProbeTT returns value from the table
func (board *Board) ProbeTT() (int, int, int, Move, bool) {
	currEntry := TranspositionTable[board.Hash%TTSize]
	if currEntry.Hash == board.Hash {
		return currEntry.Score, currEntry.Flag, currEntry.Depth, currEntry.BestMove, true
	}
	return 0, 0, 0, 0, false
}

// StoreTT stores value into the table
func (board *Board) StoreTT(depth, score, flag int, bestMove Move) {
	entry := TTEntry{
		Hash:     board.Hash,
		Depth:    depth,
		Score:    score,
		Flag:     flag,
		BestMove: bestMove,
	}
	TranspositionTable[board.Hash%TTSize] = entry
}
