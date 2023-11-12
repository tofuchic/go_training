package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuf-lesson/pb"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
);

func main() {
	person := &pb.Person{
		Id: 1,
		Name: "Namonashi",
		Email: "test@example.com",
		Occupation: pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project: map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Person_Text{
			Text: "My name is Naomonashi",
		},
		Birthday: &pb.Date{
			Year: 2000,
			Month: 1,
			Day: 1,
		},
	}

	// シリアライズ
	binData, err := proto.Marshal(person)
	if err != nil{
		log.Fatalln("Can't serialize", err)
	}

	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil{
		log.Fatalln("Can't write", err)
	}

	in, err := ioutil.ReadFile("test.bin")
	if err != nil{
		log.Fatalln("Can't read file", err)
	}

	readPerson := &pb.Person{}
	err = proto.Unmarshal(in, readPerson)
	if err != nil {
		log.Fatalln("Can't deserialise", err)
	}

	fmt.Println(readPerson)


	// JSONマッピング
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(person)
	if err != nil{
		log.Fatalln("Can't marshal to json", err)
	}

	fmt.Println(out)

	readPerson2 := &pb.Person{}
	if err := jsonpb.UnmarshalString(out, readPerson2); err != nil {
		log.Fatalln("Can't unmarshal from json", err)
	}

	fmt.Println(readPerson2)
}