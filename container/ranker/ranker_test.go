package ranker

import (
	"testing"
	"fmt"
)

type TKey int

func (s TKey) Cmp(other Key) int {
	o := other.(TKey)
	return int(s - o)
}

type TScore int

func (s TScore) Add(other Value) Value {
	o := other.(TScore)
	return s + o
}

func TestRankerAdd(t *testing.T) {
	ranker := New()
	ranker.Add(TKey(1), TScore(7))
	keys, values := ranker.Tops(10)
	for i, key := range keys {
		fmt.Println(key, values[i])
	}
}

func TestRankerDel(t *testing.T) {
	ranker := New()
	ranker.Add(TKey(1), TScore(7))
	ranker.Delete(TKey(1))
	keys, values := ranker.Tops(-1)
	for i, key := range keys {
		fmt.Println(key, values[i])
	}
}

func TestRankerUpdate(t *testing.T) {
	ranker := New()
	ranker.Add(TKey(1), TScore(7))
	ranker.Update(TKey(1), TScore(-2))
	keys, values := ranker.Tops(-1)
	for i, key := range keys {
		fmt.Println(key, values[i])
	}
}

func TestRankerRank(t *testing.T) {
	ranker := New()
	ranker.Add(TKey(1), TScore(7))
	ranker.Update(TKey(2), TScore(3))
	rank, value := ranker.Rank(TKey(2))
	fmt.Println(rank, value)

	keys, values := ranker.Tops(-1)
	for i, key := range keys {
		fmt.Println(key, values[i])
	}
}
