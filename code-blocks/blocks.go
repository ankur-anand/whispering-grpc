package main

func printXMLSize() {
	...
	fmt.Println(bufxml.Len()) // 3125 // HL
}

func printProtobuffSize() {
	...
	fmt.Println(len(mb)) // 1030 // HL
}