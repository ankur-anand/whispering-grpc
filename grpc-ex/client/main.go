package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"github.com/ankur-anand/grpc-go/grpc-ex/model"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stdout, "specify action: list or add")
		os.Exit(1)
	}

	con, err := grpc.Dial(":3002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to the model server %v", err)
	}

	client := model.NewGetUserClient(con)
	switch action := flag.Arg(0); action {
	case "add":
		if flag.NArg() != 4 {
			err = fmt.Errorf("expected three agrument for add action, [%s]", "name age email")
		} else {
			err = add(flag.Args()[1:], context.Background(), client)
		}
	case "list":
		err = list(context.Background(), client)
	default:
		err = fmt.Errorf("unkown action [%s]", action)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func list(ctx context.Context, client model.GetUserClient) error {
	u, err := client.List(ctx, &model.Void{})
	if err != nil {
		return fmt.Errorf("could not get any user : %v", err)
	}
	for _, value := range u.Users {
		fmt.Println(value.Name, value.Age, value.Email)
	}
	return nil
}

func add(args []string, ctx context.Context, client model.GetUserClient) error {
	age, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("expected third arg to be valid age number [%s]",
			err.Error())
	}
	user := &model.User{
		Name:  args[0],
		Age:   uint32(age),
		Email: args[2],
	}
	_, err = client.Add(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
