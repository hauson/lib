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
		sumer.Wait()
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
		sumer.Wait()
		i, err := sumer.Result()
		if err != nil {
			fmt.Println(err)
		}

		num := i.(int)
		fmt.Println("g2:", num)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		sumer := NewSumer(2, 2)
		sumer.expire = time.Now().Add(time.Second)
		singer.Flight(sumer)
		sumer.Wait()
		i, err := sumer.Result()
		if err != nil {
			fmt.Println(err)
		}

		num := i.(int)
		fmt.Println("g3:", num)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		sumer := NewSumer(2, 2)
		sumer.expire = time.Now().Add(time.Second)
		singer.Flight(sumer)
		sumer.Wait()
		i, err := sumer.Result()
		if err != nil {
			fmt.Println(err)
		}

		num := i.(int)
		fmt.Println("g4:", num)
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("%d second \n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(20 * time.Second)
}

type Sumer struct {
	num1   int
	num2   int
	done   chan int
	result int
	err    error
	expire time.Time
}

func NewSumer(num1, num2 int) *Sumer {
	return &Sumer{
		num1:   num1,
		num2:   num2,
		done:   make(chan int),
		expire: time.Now().Add(5 * time.Second),
	}
}

func (s *Sumer) Key() string {
	return fmt.Sprintf("sum:%d+%d", s.num1, s.num2)
}

func (s *Sumer) Wait() {
	<-s.done
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
	fmt.Println("sum done:", s.Key())
	time.Sleep(5 * time.Second)
	s.result = s.num1 + s.num2
	s.err = nil
	s.Close()
}

func (s *Sumer) Expire() time.Time {
	return s.expire
}

func (s *Sumer) Close() {
	close(s.done)
}
