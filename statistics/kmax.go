package statisitcs

import (
	"errors"
	"fmt"
)

// KMax the first k is from 0, and numbers remain unchanged
func KMax(numbers []float64, k int) (float64, error) {
	cp := make([]float64, len(numbers))
	copy(cp, numbers)
	return kMaxDo(cp, k)
}

func adjust(list []float64) (mid int) {
	left, right := 0, len(list)-1
	for left < right {
		for right > mid && list[right] <= list[mid] {
			right--
		}

		if right > mid {
			list[right], list[mid] = list[mid], list[right]
			mid = right
		}

		for left < mid && list[left] >= list[mid] {
			left++
		}

		if left < mid {
			list[left], list[mid] = list[mid], list[left]
			mid = left
		}
	}
	return
}

func kMaxDo(numbers []float64, k int) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("numbers is empty")
	}

	if k > len(numbers)-1 {
		return 0.0, fmt.Errorf("k %d exceed number last index %d", k, len(numbers)-1)
	}

	mid := adjust(numbers)
	switch {
	case mid > k:
		return kMaxDo(numbers[:mid], k)
	case mid < k:
		return kMaxDo(numbers[mid+1:], k-(mid+1))
	default:
		return numbers[mid], nil
	}
}
