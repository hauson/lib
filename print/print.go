package print

import (
	"fmt"
	"time"
	"sync"
)

/*
	demo:
	Println("key1", "hello,world")
	Println("key2", "hello,world")
*/

var selector = newColorSelector()

type colorSelector struct {
	colors      []string
	topicColors map[string]string
	idx         int
	sync.Mutex
}

func newColorSelector() *colorSelector {
	return &colorSelector{
		colors:      []string{Green, Yellow, Red, Blue, Magenta, White, Cyan},
		topicColors: map[string]string{},
	}
}

func (c *colorSelector) get(topic string) string {
	c.Lock()
	defer c.Unlock()

	if color, ok := c.topicColors[topic]; ok {
		return color
	}
	color := c.colors[c.idx]
	c.topicColors[topic] = color

	c.idx++
	if c.idx > len(c.colors) {
		c.idx = 0
	}

	return color
}

func Println(topic string, a ...interface{}) {
	color := selector.get(topic)
	topicAndTime := fmt.Sprintf("[%s:%s]", topic, time.Now().Format("2006-01-02 15:04:05"))
	a = append([]interface{}{color, topicAndTime}, a...)
	fmt.Println(a...)
}

func Printf(topic, format string, a ...interface{}) {
	color := selector.get(topic)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	format = color + "[%s:%s]" + format
	a = append([]interface{}{topic, timeStr}, a...)
	fmt.Printf(format, a...)
}
