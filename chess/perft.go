package chess

import "fmt"

// Perft (Performance test) is used to evaluate the correctness by checking pseudolegal moves at each depth
func (board *Board) Perft(depth int) uint64 {
	if depth == 0 {
		return 1
	}
	moves := board.GeneratePseudoLegalMoves()
	var total uint64 = 0
	for i := 0; i < moves.Count; i++ {
		move := moves.Moves[i]
		boardCopy := *board
		ok := boardCopy.MakeMove(move)
		if ok {
			total += boardCopy.Perft(depth - 1)
		}
	}
	return total
}

// PerftDivide is used to debug the moves
func (board *Board) PerftDivide(depth int) {
	moves := board.GeneratePseudoLegalMoves()
	var total uint64 = 0
	fmt.Printf("\n Perft Divide Depth %d\n", depth)
	for i := 0; i < moves.Count; i++ {
		move := moves.Moves[i]
		boardCopy := *board
		if boardCopy.MakeMove(move) {
			nodes := boardCopy.Perft(depth - 1)
			fromStr, _ := ParseSquareI2S(move.From())
			toStr, _ := ParseSquareI2S(move.To())
			fmt.Printf("%s%s: %d\n", fromStr, toStr, nodes)
			total += nodes
		}
	}
	fmt.Printf("\n Perft Divide Total %d\n", total)
}
