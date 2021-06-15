package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"sync"
)

func ReadFileLinesEx(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var s []string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		s = append(s, string(line))
	}
	return s, nil
}

func ReadFileWords(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}

func ReadWholeFile1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	if _, err := file.Read(buffer); err != nil {
		return "", err
	}

	return string(buffer), nil
}

func ReadWholeFile(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ReadWholeFileEx(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

func ReadByBlocks(filePath string, concurrency int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileSize := int(fileInfo.Size())
	bufferSize := int(math.Ceil(float64(fileSize) / float64(concurrency)))

	blockStrings := new(sync.Map)
	var wg sync.WaitGroup
	wg.Add(concurrency)

	errChan := make(chan error, concurrency)
	for i := 0; i < concurrency; i++ {
		chunkSize := bufferSize
		if i == concurrency-1 {
			chunkSize = fileSize - (i * bufferSize)
		}

		go func(i, bufSize int, offset int64) {
			defer wg.Done()

			buffer := make([]byte, bufSize)
			if _, err := file.ReadAt(buffer, offset); err != nil {
				errChan <- err
			} else {
				blockStrings.Store(i, string(buffer))
			}
		}(i, chunkSize, int64(i*bufferSize))
	}
	wg.Wait()
	close(errChan)

	for err := range errChan {
		return "", err
	}

	var str string
	for i := 0; i < concurrency; i++ {
		s, _ := blockStrings.Load(i)
		str += s.(string)
	}

	return str, nil
}

// split words by space
func ReadWords(str string) []string {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}
