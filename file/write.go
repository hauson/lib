package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func WriteFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, data)
	return err
}

func WriteFile1(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}

	return writer.Flush()
}

func WriteFile2(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	return err
}

func WriteFile3(filename, data string) error {
	d := []byte(data)
	return ioutil.WriteFile(filename, d, 0666)
}

func WritFileByStrings(filename string, ss []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, s := range ss {
		if _, err := file.WriteString(s + "\n"); err != nil {
			return err
		}
	}
	return file.Sync()
}

func WritFileByStrings1(filename string, ss []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, s := range ss {
		if _, err := writer.WriteString(s + "\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}

