package chess

import "math/bits"

func (board *Board) AlphaBeta(depth, alpha, beta int, isMax bool) int {
	if depth == 0 {
		return board.Evaluate()
	}
	moves := board.GenerateLegalMoves()
	color := board.SideToMove
	if moves.Count == 0 {
		king := bits.TrailingZeros64(board.Colors[color] & board.Pieces[King])
		if board.IsSquareAttacked(uint8(king), color^1) {
			if isMax {
				return ColorScores[White] + depth
			} else {
				return ColorScores[Black] - depth
			}
		} else {
			return 0
		}
	}
	if isMax {
		bestScore := ColorScores[White]
		for i := 0; i < moves.Count; i++ {
			boardCopy := *board
			boardCopy.MakeMove(moves.Moves[i])
			score := boardCopy.AlphaBeta(depth-1, alpha, beta, false)
			if score > bestScore {
				bestScore = score
			}
			if score > alpha {
				alpha = score
			}
			if alpha >= beta {
				break
			}
		}
		return bestScore
	} else {
		bestScore := ColorScores[Black]
		for i := 0; i < moves.Count; i++ {
			boardCopy := *board
			boardCopy.MakeMove(moves.Moves[i])
			score := boardCopy.AlphaBeta(depth-1, alpha, beta, true)
			if score < bestScore {
				bestScore = score
			}
			if score < beta {
				beta = score
			}
			if beta <= alpha {
				break
			}
		}
		return bestScore
	}
}
