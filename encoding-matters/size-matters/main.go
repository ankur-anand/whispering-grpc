package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"unsafe"

	proto "github.com/golang/protobuf/proto"
)

// OnekRequest represents an request of size (approzimately)
// one KB in go type
type OnekRequest struct {
	ID   int64
	Data [1024]byte
}

func main() {
	req := OnekRequest{
		ID:   1234,
		Data: [1024]byte{},
	} // size in go type 1024 + 8 = 1032 no padding is needed

	// verify using unsafe package
	goSize := unsafe.Sizeof(req)
	fmt.Println(goSize)

	var bufxml bytes.Buffer
	xmlEnc := xml.NewEncoder(&bufxml)
	xmlEnc.Encode(req)
	// what is the size of the encoded XML ?
	fmt.Println(bufxml.Len()) // 3125

	var jsonXML bytes.Buffer
	jsonEnc := json.NewEncoder(&jsonXML)
	jsonEnc.Encode(req)
	fmt.Println(jsonXML.Len()) // 2069

	// Let's try binary encoding
	var gobBuffer bytes.Buffer
	gobEnc := gob.NewEncoder(&gobBuffer)
	gobEnc.Encode(req)
	fmt.Println(gobBuffer.Len()) // 1111

	// Let's encode the same in protobuff
	var reqProto Onekreq
	sl := make([]byte, 1024)
	for i := range sl {
		sl[i] = byte(i / 255)
	}
	reqProto = Onekreq{
		Id:   1234,
		Data: sl,
	}

	mb, err := proto.Marshal(&reqProto)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(mb)) // 1030

}
