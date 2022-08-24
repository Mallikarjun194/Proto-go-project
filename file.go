package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func WriteToFile(fileName string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Error while marshalling", err)
		return
	}

	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatal("Can't write to file", err)
		return
	}
	fmt.Println("Data is written successfully")
}

func ReadFromFile(fileName string, pb proto.Message) {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error while reading", err)
		return
	}
	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatal("Error while unmarshelling", err)
		return
	}

}
