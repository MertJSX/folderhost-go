package utils

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"sync"
)

func CreateFileAsync(path string, content string, wg *sync.WaitGroup, ch chan<- error) {
	go func() {
		defer wg.Done()
		buf := make([]byte, 32*1024) // 32KB buffer
		file, _ := os.Create(path)
		defer file.Close()

		writer := bufio.NewWriterSize(file, len(buf))
		_, err := io.CopyBuffer(writer, bytes.NewReader([]byte(content)), buf)
		ch <- err
	}()
}
