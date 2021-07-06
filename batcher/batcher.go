package batcher

import "github.com/hauson/lib/statisitcs"

type Batcher struct {
	total  int
	cntPer int
}

func New(total, cntPer int) *Batcher {
	return &Batcher{
		total:  total,
		cntPer: cntPer,
	}
}

func (b *Batcher) Range(fn func(offset, limit int)) {
	for offset := 0; offset < b.total; offset += b.cntPer {
		limit := statisitcs.MinInts(b.cntPer, b.total-offset)
		fn(offset, limit)
	}
}
