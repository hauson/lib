package batcher

import (
	"reflect"
)

func Split(s []int, limit int) [][]int {
	var ss [][]int
	for i:=0;i<len(s);i+=limit {
		if (i+limit) <= len(s) {
			ss = append(ss, s[i:i+limit])
		}else {
			ss = append(ss, s[i:])
		}
	}

	return ss
}

func Join(ss [][]int) []int {
	var s []int
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}
