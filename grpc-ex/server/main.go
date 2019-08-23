package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/ankur-anand/grpc-go/grpc-ex/model"
	proto "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

const (
	fileDB  = "person.db"
	sizeLen = 8
)

type length int64

var endianness = binary.LittleEndian

type UserModel struct {
}

func (u UserModel) List(ctx context.Context, void *model.Void) (*model.Users,
	error) {
	b, err := ioutil.ReadFile(fileDB)
	if err != nil {
		return &model.Users{}, fmt.Errorf("could not read %s: %v", fileDB, err)
	}
	var users model.Users
	for {
		if len(b) == 0 {
			return &users, nil
		} else if len(b) < sizeLen {
			return &users, fmt.Errorf("db improperly formatted")
		}
		var l length
		if err := binary.Read(bytes.NewReader(b[:sizeLen]), endianness, &l); err != nil {
			return &users, fmt.Errorf("error decoding the message %v", err)
		}
		b = b[sizeLen:]
		var person model.User
		if err := proto.Unmarshal(b[:l], &person); err == io.EOF {
			return &users, nil
		} else if err != nil {
			return &users, fmt.Errorf("could not read person: %v", err)
		}
		users.Users = append(users.Users, &person)
		b = b[l:]
		fmt.Println(person)
	}
}

func (u UserModel) Add(ctx context.Context, user *model.User) (*model.Void,
	error) {

	b, err := proto.Marshal(user)
	if err != nil {
		return &model.Void{}, fmt.Errorf("could not encode task: %v", err)
	}

	f, err := os.OpenFile(fileDB, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return &model.Void{}, fmt.Errorf("could not open file %s: %v", fileDB, err)
	}
	// store the length of the person
	if err := binary.Write(f, endianness, length(len(b))); err != nil {
		return &model.Void{},
			fmt.Errorf("could not encode length of message %d: %v", len(b), err)
	}
	_, err = f.Write(b)
	if err != nil {
		return &model.Void{}, fmt.Errorf("could not write task to file: %v",
			err)
	}
	if err := f.Close(); err != nil {
		return &model.Void{}, fmt.Errorf("could not close the file %s: %v",
			fileDB, err)
	}
	return &model.Void{}, nil
}

func main() {
	svc := grpc.NewServer()
	var um UserModel
	model.RegisterGetUserServer(svc, um)
	l, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("could not start server on [%s]: %w", "3002", err)
	}
	err = svc.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
