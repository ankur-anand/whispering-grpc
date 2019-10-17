package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
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
