package pagefile

import (
	"os"
	"sync"
)

type PageFile struct {
	page    *Pager
	namer   *Namer
	fd      *os.File
	lineCh  chan string
	exitSig chan int
	wg      sync.WaitGroup
}

func New(file string, linesPage int) (*PageFile, error) {
	pageFile := &PageFile{
		page:    NewPager(linesPage),
		namer:   NewNamer(file),
		lineCh:  make(chan string, 2000),
		exitSig: make(chan int),
	}

	if err := pageFile.openOrCreate(); err != nil {
		return nil, err
	}

	pageFile.wg.Add(1)
	go pageFile.run()

	return pageFile, nil
}

func (file *PageFile) WriteLines(lines ...string) {
	for _, line := range lines {
		file.lineCh <- line
	}
}

func (file *PageFile) openOrCreate() error {
	if file.fd != nil {
		file.fd.Close()
	}

	fileName := file.namer.makeName(file.page.num)
	fd, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	file.fd = fd
	return nil
}
