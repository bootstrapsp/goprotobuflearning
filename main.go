package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"

	"github.com/bootstrapsp/goplayground/goProtobufLearning/src/simple"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	// write to file func
	// read from file()

	writeToFileName("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("reading the message from the buffer", sm2)

	smAsString := toJSON(sm)

	fmt.Println(smAsString)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func from()  {
	
}
func writeToFileName(fname string, pb proto.Message) error {

	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	wr := ioutil.WriteFile(fname, out, 0644)
	if wr != nil {
		log.Fatalln("Can't write to file", err)
		return err

	}

	fmt.Println("File successfully writen")
	return nil

}

func readFromFile(fname string, pb proto.Message) error {

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Failed to read from file", err)
		return err
	}

	rf := proto.Unmarshal(in, pb)

	if rf != nil {
		log.Fatalln("Failed to read from the file", rf)
		return rf
	}

	return nil
}
func doSimple() *simplepb.SimpleMessage {
	// This is a constructor
	sm := simplepb.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       "Simple Message",
		SampleList: []int32{1, 4, 5, 6, 3},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)

	fmt.Println("This is the ID", sm.GetId())

	return &sm
}
