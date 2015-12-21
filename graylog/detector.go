package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
)

func check(testBytes []byte) string {

	b := bytes.NewBufferString(string(testBytes))
	var r io.ReadCloser
	var err interface{} = nil
	out := new(bytes.Buffer)

	if testBytes[0] == 0x1f && testBytes[1] == 0x8b {
		r, err = gzip.NewReader(b)
	} else if testBytes[0] == 0x78 && testBytes[1] == 0x9c {
		r, err = zlib.NewReader(b)
	} else {
		r = ioutil.NopCloser(b)
	}
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer r.Close()
	out.ReadFrom(r)

	return out.String()

}
