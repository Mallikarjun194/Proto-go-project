package main

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

func ToJson(pb proto.Message) string {
	// To print it in multiple lines we use Multiline
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)
	// out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatal("Error while marshelling", err)
		return ""
	}
	return string(out)
}
func FromJson(in string, pb proto.Message) {
	//
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	// if err := protojson.Unmarshal([]byte(in), pb); err != nil
	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatal("Error while unmarshelling", err)
	}
	fmt.Println(pb)
}
