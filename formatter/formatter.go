package formatter

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func Beautify(data []byte) ([]byte, error) {
	var dst bytes.Buffer
	err := json.Indent(&dst, data, "", "\t")
	if err != nil {
		return nil, err
	}

	return dst.Bytes(), nil
}

func Escape(data string) string {
	return strconv.Quote(data)
}

func Minify(data []byte) ([]byte, error) {
	var dst bytes.Buffer
	err := json.Compact(&dst, data)
	if err != nil {
		return nil, err
	}

	return dst.Bytes(), nil
}

func Unescape(data string) (string, error) {
	return strconv.Unquote(data)
}
