package sequence

type Sequence struct {
	max uint64
	seq uint64
}

func New(max uint64) *Sequence {
	return &Sequence{max: max}
}

func (s *Sequence) Next() uint64 {
	s.seq = (s.seq + 1) % s.max
	return s.seq
}
