message onekreq {
    uint64 id = 1; // HL
    bytes data = 2;
}

func main() {
	var reqProto Onekreq
	...
	mb, err := proto.Marshal(&reqProto)
	...
	fmt.Println(len(mb)) // 1030 // HL
}