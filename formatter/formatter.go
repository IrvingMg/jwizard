package formatter

import (
	"bytes"
	"encoding/json"
)

func Beautify(data []byte) ([]byte, error) {
	var dst bytes.Buffer
	err := json.Indent(&dst, data, "", "\t")
	if err != nil {
		return nil, err
	}

	return dst.Bytes(), nil
}

func Minify(data []byte) ([]byte, error) {
	var dst bytes.Buffer
	err := json.Compact(&dst, data)
	if err != nil {
		return nil, err
	}

	return dst.Bytes(), nil
}
