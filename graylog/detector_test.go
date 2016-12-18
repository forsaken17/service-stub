package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"testing"
	"time"

	"fmt"

	"github.com/duythinht/gelf"
	"github.com/duythinht/gelf/chunk"
)

type Test struct {
	in  []byte
	out string
}

var tests = []Test{
	{makeGzipMsg("test gzip"), "test gzip"},
	{makeZlibMsg("test zlib"), "test zlib"},
}

func TestCheck(t *testing.T) {
	for _, test := range tests {
		var out = check(test.in)
		if out != test.out {
			t.Error("unexpexted: " + out)
		}
	}

}
func TestChunked(t *testing.T) {
	var expected = "test zlib1"
	var buffers = makeGelfChunked(expected)
	for _, chunked := range buffers {
		fmt.Printf("%#x\n", chunked)
		fmt.Printf("%#x\n", chunked[0:2])
		fmt.Printf("%#x\n", chunked[2:10])
		fmt.Printf("%#x\n", chunked[10:11])
		fmt.Printf("%#x\n", chunked[11:12])
		out := check(chunked)
		if out != expected {
			t.Error("unexpexted: " + out)
		}
	}
}
func makeGelfChunked(input string) [][]byte {
	message := gelf.Create(input).
		SetTimestamp(time.Now().Unix()).
		SetFullMessage("This is full message").
		SetLevel(3).
		SetHost("chat Server").
		ToJSON()
	ZippedMessage := chunk.ZipMessage(message)
	var MaxChunkSize int = 50
	var buffers [][]byte
	fmt.Println(len(ZippedMessage))
	if len(ZippedMessage) > MaxChunkSize {
		buffers = chunk.GetGelfChunks(ZippedMessage, MaxChunkSize)
	}
	return buffers
}

func makeGzipMsg(input string) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte(input))
	w.Close()
	return b.Bytes()
}

func makeZlibMsg(input string) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte(input))
	w.Close()
	return b.Bytes()
}
