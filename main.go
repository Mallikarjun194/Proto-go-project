package main

import (
	pb "Proto-go-project/proto"
	"fmt"
	"google.golang.org/protobuf/proto"
	"reflect"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:       42,
		IsSimple: true,
		Name:     "A name",
		Lists:    []int32{1, 2, 3, 4},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy1{Id: 42, Name: "My Name-1"},
		MultipleDummies: []*pb.Dummy1{
			{Id: 43, Name: "My Name-2"},
			{Id: 44, Name: "My Name-3"},
		},
	}
}

func doEnum() *pb.Enumerations {
	return &pb.Enumerations{
		EyeColor: 2,
		//pb.EyeColor_EYE_COLOR_BLUE
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println("Msg is id", x.Id)
	case *pb.Result_Message:
		fmt.Println("Message is msg", x.Message)
	default:
		fmt.Errorf("Msg is of unexpected type %v", x)

	}
}

func doMaps() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myId-1": {Id: 42},
			"myId-2": {Id: 43},
			"myId-3": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "sample.bin"
	WriteToFile(path, p)
	msg := &pb.Simple{}
	ReadFromFile(path, msg)
	fmt.Println(msg)
}

func doToJson(p proto.Message) string {
	jsonString := ToJson(p)
	fmt.Println(jsonString)
	return jsonString
}

func doFromJson(in string, t reflect.Type) proto.Message {
	msg := reflect.New(t).Interface().(proto.Message)
	FromJson(in, msg)
	return msg
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doComplex())
	fmt.Println(doEnum())
	doOneOf(&pb.Result_Id{Id: 45})
	doOneOf(&pb.Result_Message{Message: "Hello my msg"})
	fmt.Println(doMaps())
	doFile(doSimple())
	fmt.Println("####")
	jsonString := doToJson(doSimple())
	msg := doFromJson(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString)
	fmt.Println(msg)

	fmt.Println("####")
	jsonString = doToJson(doComplex())
	msg = doFromJson(jsonString, reflect.TypeOf(pb.Complex{}))
	fmt.Println(jsonString)
	fmt.Println(msg)

	fmt.Println("@@#@#@#@")
	// Discard unknown field n prints only id
	fmt.Println(doFromJson(`{"id":42, "unknown":"true"}`, reflect.TypeOf(pb.Simple{})))

}
