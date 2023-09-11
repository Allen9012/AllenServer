package bigsave

import (
	"bytes"
	"encoding/gob"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc: 大数据类型可以提供压缩
  @modified by:
**/

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
