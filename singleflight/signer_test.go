package singleflight

import (
	"testing"
	"time"
	"fmt"
)

func TestSingleflight(t *testing.T) {
	singer := New()
	go func() {
		sumer := NewSumer(1, 2)
		singer.Flight(sumer)
		<-sumer.Done()
		i, err := sumer.Result()
		if err != nil {
			fmt.Println(err)
		}

		num := i.(int)
		fmt.Println("g1:", num)
	}()

	go func() {
		sumer := NewSumer(1, 2)
		singer.Flight(sumer)
		<-sumer.Done()
		i, err := sumer.Result()
		if err != nil {
			fmt.Println(err)
		}

		num := i.(int)
		fmt.Println("g2:", num)
	}()
	go func() {
		sumer := NewSumer(1, 2)
		singer.Flight(sumer)
		<-sumer.Done()
		i, err := sumer.Result()
		if err != nil {
			fmt.Println(err)
		}

		num := i.(int)
		fmt.Println("g3:", num)
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("%d second \n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(13 * time.Second)
}

type Sumer struct {
	num1   int
	num2   int
	done   chan int
	result int
	err    error
	valid  bool
}

func NewSumer(num1, num2 int) *Sumer {
	return &Sumer{
		num1: num1,
		num2: num2,
		done: make(chan int),
	}
}

func (s *Sumer) Key() string {
	return fmt.Sprintf("sum:%d+%d", s.num1, s.num2)
}

func (s *Sumer) Done() <-chan int {
	return s.done
}

func (s *Sumer) Result() (interface{}, error) {
	return s.result, s.err
}

func (s *Sumer) CopyResult(i Executor) {
	e := i.(*Sumer)
	s.result = e.result
	s.err = e.err
}

func (s *Sumer) Do() {
	fmt.Println("sum done")
	time.Sleep(5 * time.Second)
	s.result = s.num1 + s.num2
	s.err = nil
	s.SetValid(false)
	s.Close()
}

func (s *Sumer) Valid() bool {
	return s.valid
}

func (s *Sumer) SetValid(b bool) {
	s.valid = b
}

func (s *Sumer) Close() {
	close(s.done)
}
