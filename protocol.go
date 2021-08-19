package main

import (
	"encoding/binary"
)


func DecodeProto(buf []byte) (uint16, string) {
	data_len := binary.BigEndian.Uint16(buf)
	//fmt.Println("in Decode",buf)
	str := string(buf[2:data_len + 2])
	//fmt.Println("Decode len=",data_len, buf[:data_len + 2], str)
	return data_len, str
}

func EncodeProto(s string) (uint16,[]byte) {
	buf := make([]byte,256)
	//fmt.Println("in encoded :",  []byte(s))
	var len uint16 = uint16(len(s))
	binary.BigEndian.PutUint16(buf, len)
	var i uint16 = 2
	for _,c := range([]byte(s)) {
		buf[i] = c
		i++
	}
	return  len, buf[:len+2]
}
