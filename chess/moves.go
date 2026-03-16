package chess

type Move uint16
type MoveList struct {
	Moves [256]Move
	Count int
}

func (ml *MoveList) Add(m Move) {
	if ml.Count >= 256 {
		panic("Too many moves")
	}
	ml.Moves[ml.Count] = m
	ml.Count++
}

func NewMove(from uint8, to uint8, flags uint16) Move {
	return Move(flags<<12 | uint16(to)<<6 | uint16(from))
}
func (m Move) From() uint8 {
	return uint8(m & 63)
}
func (m Move) To() uint8 {
	return uint8((m >> 6) & 63)
}
func (m Move) Flags() uint16 {
	return uint16(m >> 12)
}
