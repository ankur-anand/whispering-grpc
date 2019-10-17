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
	flatbuffers "github.com/google/flatbuffers/go"
)

// START OMIT

// OnekRequest represents an request of size (approzimately)
// one KB in go type
type OnekRequest struct { // HL
	ID   int64      // HL
	Data [1024]byte // HL
} // HL

// END OMIT

func printGoSize() {
	// size in go type 1024 + 8 = 1032 no padding is needed
	req := OnekRequest{
		ID:   1234,
		Data: [1024]byte{},
	}

	// verify using unsafe package
	goSize := unsafe.Sizeof(req) // HL
	fmt.Println(goSize)
	// OUTPUT 1032 // HL
}

func printJSONSize() {
	// size in go type 1024 + 8 = 1032 no padding is needed
	req := OnekRequest{
		ID:   1234,
		Data: [1024]byte{},
	}

	// verify using unsafe package
	var jsonXML bytes.Buffer
	jsonEnc := json.NewEncoder(&jsonXML)
	jsonEnc.Encode(req)
	fmt.Println(jsonXML.Len()) // 2069 // HL
}

func printXMLSize() {
	// size in go type 1024 + 8 = 1032 no padding is needed
	req := OnekRequest{
		ID:   1234,
		Data: [1024]byte{},
	}

	var bufxml bytes.Buffer
	xmlEnc := xml.NewEncoder(&bufxml)
	xmlEnc.Encode(req)
	fmt.Println(bufxml.Len()) // 3125 // HL
}

func printProtobuffSize() {
	var reqProto Onekreq
	sl := make([]byte, 1024)
	for i := range sl {
		sl[i] = byte(i / 255)
	}
	reqProto = Onekreq{
		Id:   1234,
		Data: sl,
	}
	mb, _ := proto.Marshal(&reqProto)
	fmt.Println(len(mb)) // 1030 // HL
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

	// flatbuffer

	b := flatbuffers.NewBuilder(0)
	OneKReqStartDataVector(b, 1024)
	for i := range sl {
		b.PrependByte(byte(i / 255))
	}
	data := b.EndVector(1024)
	OneKReqStart(b)
	OneKReqAddId(b, 1234)
	OneKReqAddData(b, data)
	reqF := OneKReqEnd(b)
	b.Finish(reqF)
	bytesF := b.Bytes[b.Head():]
	fmt.Println(len(bytesF))

}
