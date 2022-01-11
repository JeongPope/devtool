package devtool

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"io/ioutil"
)

func DecompressFlate(data []byte) ([]byte, error) {
	tReader := bytes.NewReader(data)
	tFlateReader := flate.NewReader(tReader)

	return ioutil.ReadAll(tFlateReader)
}

func DecompressGzip(data []byte) (resData []byte, err error) {
	tReader := bytes.NewReader(data)

	r, err := gzip.NewReader(tReader)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}
