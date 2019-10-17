message onekreq1 {
    uint64 id = 17; // HL
    bytes data = 2;
}

func main() {
	var reqProto1 onekreq1
	...
	mb1, err := proto.Marshal(&reqProto1)
	...
	fmt.Println(len(mb1)) // 1031 // HL
}