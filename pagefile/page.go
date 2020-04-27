package pagefile

import (
	"strings"
	"strconv"
)

type Pager struct {
	num       int
	linesPage int
	lineCnt   int
}

func NewPager(linesPage int) *Pager {
	return &Pager{
		num:       1,
		linesPage: linesPage,
	}
}

func (p *Pager) next() *Pager {
	return &Pager{
		num:       p.num + 1,
		linesPage: p.linesPage,
	}
}

func (p *Pager) isFull() bool {
	return p.lineCnt > p.linesPage
}

type Namer struct {
	prefix string
	suffix string
}

func NewNamer(file string) *Namer {
	prefix, suffix := splitFileName(file)
	return &Namer{
		prefix: prefix,
		suffix: suffix,
	}
}

func (n *Namer) makeName(page int) string {
	return n.prefix + strconv.Itoa(page) + n.suffix
}

func splitFileName(file string) (name, suffix string) {
	ss := strings.Split(file, ".")
	switch n := len(ss); n {
	case 0:
		return "", ""
	case 1:
		return file, ""
	default:
		return strings.Join(ss[:n-1], "."), "." + ss[n-1]
	}
}
