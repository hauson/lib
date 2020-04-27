package pagefile

import (
	"testing"
	"fmt"
	"time"
)

func TestFile(t *testing.T) {
	fd, err := New("cui.txt", 10)
	if err != nil {
		t.Error(err)
	}
	defer fd.Close()

	for i := 0; i < 37; i++ {
		fd.WriteLines(fmt.Sprintf("hello,world %d", i))
	}
}

func TestFileWithGoroutine(t *testing.T) {
	fd, err := New("cui.txt", 40)
	if err != nil {
		t.Error(err)
	}
	defer fd.Close()

	go func() {
		for i := 0; i < 37; i++ {
			fd.WriteLines(fmt.Sprintf("hello, world %d", i))
		}
	}()

	go func() {
		for i := 0; i < 37; i++ {
			fd.WriteLines(fmt.Sprintf("hi, girl %d", i))
		}
	}()

	go func() {
		for i := 0; i < 37; i++ {
			fd.WriteLines(fmt.Sprintf("how do you do , boy %d", i))
		}
	}()

	time.Sleep(1 * time.Minute)
}

func TestSpiltFileNameNil(t *testing.T) {
	fmt.Println(splitFileName(""))
}

func TestSpiltFileNameNoSuffix(t *testing.T) {
	fmt.Println(splitFileName("cui"))
}

func TestSpiltFileName(t *testing.T) {
	fmt.Println(splitFileName("hao.cui.txt"))
}
