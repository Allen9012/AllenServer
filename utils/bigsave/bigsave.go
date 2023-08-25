package bigsave

import (
	"bytes"
	"encoding/gob"
)

func Encoder[T any](data []T) string {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(&data)
	if err != nil {
		return ""
	}
	b := GZipBytes(buf.Bytes())
	return string(b[:])
}

func Decoder[T any](i string) []T {
	byteEn := UGZipBytes([]byte(i))
	decoder := gob.NewDecoder(bytes.NewReader(byteEn))
	ret := []T{}
	err := decoder.Decode(&ret)
	if err != nil {
		return nil
	}
	return ret
}
