package main

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/rotemtam/pbreflect-example/model"
)

func main() {
	bytes := genBinary()
	pb := getProto("com.rotemtam.User", bytes)
	fmt.Println(pb)
}

func getProto(messageType string, messageBytes []byte) proto.Message {
	pbtype := proto.MessageType(messageType)
	msg := reflect.New(pbtype.Elem()).Interface().(proto.Message)
	proto.Unmarshal(messageBytes, msg)
	return msg
}

func genBinary() []byte {
	user := &model.User{UserName: "John Doe"}
	bytes, _ := proto.Marshal(user)
	return bytes
}
