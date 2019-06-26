package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func uintToByte(num uint64) []byte {

	var buffer bytes.Buffer

	//使用二进制编码
	err := binary.Write(&buffer, binary.LittleEndian, &num)
	if err != nil {
		fmt.Println("binary.Write err", err)
		return nil
	}

	return buffer.Bytes()
}
