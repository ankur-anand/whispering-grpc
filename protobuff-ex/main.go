package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ankur-anand/grpc-go/protobuff-ex/persondb"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stdout, "specify action: list or add")
		os.Exit(1)
	}

	var err error
	switch action := flag.Arg(0); action {
	case "add":
		if flag.NArg() != 4 {
			err = fmt.Errorf("expected three agrument for add action, [%s]", "name age email")
		} else {
			err = add(flag.Args()[1:])
		}
	case "list":
		err = list()
	default:
		err = fmt.Errorf("unkown action [%s]", action)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func list() error {
	return nil
}

// what save the list inside the file.
// So the question of efficiency arise with the encoding.
// we can save this as a plain text, efficient ?

func add(args []string) error {
	age, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("expected third arg to be valid age number [%s]",
			err.Error())
	}
	p := &persondb.Person{
		Name:  args[0],
		Age:   uint32(age),
		Email: args[2],
	}
	fmt.Println(proto.MarshalTextString(p))
	return nil
}
