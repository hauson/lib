package sortmap

import (
	"fmt"
	"testing"
)

func TestSortMap(t *testing.T) {
	m := NewSortedMap()

	//1. IsExist
	k := NKey(7)
	fmt.Println("nkey of 7 is exist:", m.IsExist(k))

	//2. Add
	m.Add(k, 700)
	fmt.Println("nkey of 7 is exist:", m.IsExist(k))

	//2-1. Add
	m.Add(k, 701)
	fmt.Println("nkey of 7 is exist:", m.IsExist(k))

	//3. Get
	result, ok := m.Get(k)
	fmt.Println("result:", result.(int), " ok:", ok)

	//4. inner Sort
	for i := 0; i < 10; i++ {
		k := NKey(i)
		v := i * 101
		m.Add(k, v)
	}

	//5. Del
	m.Del(NKey(7))

	//6. Range
	for _, item := range m.Range() {
		fmt.Println(item.K, item.V)
	}
}

type NKey int

func (k NKey) Less(other Key) bool {
	return k < other.(NKey)
}

func (k NKey) Equal(other Key) bool {
	return k == other.(NKey)
}
