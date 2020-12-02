package main

import (
	"clientGRPCJsonXml"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func getPersonByPhoneNumber(ctx context.Context, client clientGRPCJsonXml.RPCServiceClient, number string, phoneType int) *clientGRPCJsonXml.Person {
	phone := &clientGRPCJsonXml.PhoneNumber{Number: number, Kind: phoneType}
	format := "json"
	payload, _ := json.Marshal(phone)
	payloadString := string(payload)
	message := &clientGRPCJsonXml.Message{Format: &format, Payload: &payloadString}
	response, err := client.GetPersonByPhoneNumber(ctx, message)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var person clientGRPCJsonXml.Person
	_ = json.Unmarshal([]byte(*response.Payload), &person)
	fmt.Println("getPerson, name: " + person.Name)
	return &person
}

func editPeople(ctx context.Context, client clientGRPCJsonXml.RPCServiceClient, people []clientGRPCJsonXml.Person) {
	stream, err := client.EditPeople(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, person := range people {
		marshalPerson, _ := json.Marshal(person)
		marshalPersonString := string(marshalPerson)
		format := "json"
		message := clientGRPCJsonXml.Message{Format: &format, Payload: &marshalPersonString}
		if err := stream.Send(&message); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, person, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	var response clientGRPCJsonXml.ResponseEdit
	_ = json.Unmarshal([]byte(*reply.Payload), &response)
	if response.Result {
		fmt.Println("People edited successfully")
	} else {
		fmt.Println("Error editing people")
	}
}
func listPeople(ctx context.Context, client clientGRPCJsonXml.RPCServiceClient, number clientGRPCJsonXml.PhoneNumber) {
	format := "json"
	marshalNumber, _ := json.Marshal(number)
	marshalNumberString := string(marshalNumber)
	reqMessage := clientGRPCJsonXml.Message{Format: &format, Payload: &marshalNumberString}
	stream, err := client.ListPeopleByPhoneType(ctx, &reqMessage)
	if err != nil {
		fmt.Println("error with stream, " + err.Error())
		return
	}
	for {
		rcvStream, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListPeople(_) = _, %v", client, err)
		}
		var person clientGRPCJsonXml.Person
		_ = json.Unmarshal([]byte(*rcvStream.Payload), &person)
		fmt.Println(person.Name)
	}
}
func getPeopleById(ctx context.Context, client clientGRPCJsonXml.RPCServiceClient, ids []clientGRPCJsonXml.RequestId) {
	stream, err := client.GetPeopleById(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a request : %v", err)
			}
			var person clientGRPCJsonXml.Person
			_ = json.Unmarshal([]byte(*in.Payload), &person)
			log.Printf("Got person %s", person.Name)
		}
	}()
	for _, id := range ids {
		format := "json"
		marshalId, _ := json.Marshal(id)
		marshalIdString := string(marshalId)
		sendMessage := clientGRPCJsonXml.Message{Format: &format, Payload: &marshalIdString}
		if err := stream.Send(&sendMessage); err != nil {
			log.Fatalf("Failed to send a request: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
func main() {
	var conn *grpc.ClientConn
	ctx := context.Background()
	conn, err := grpc.Dial("10.8.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := clientGRPCJsonXml.NewRPCServiceClient(conn)

	number := "4365365432"
	phoneType := 1
	fmt.Println("--------- GET A PERSON WITH NUMBER: " + number + " ----------")
	person := getPersonByPhoneNumber(ctx, client, number, phoneType)

	fmt.Println("-------- Editing John Doe in Giovanni Doe and Mario Rossi in Mario Bianchi --------")
	newName := "Giovanni Doe"
	person.Name = newName
	number = "452376467"
	person2 := getPersonByPhoneNumber(ctx, client, number, phoneType)
	newName2 := "Mario Bianchi"
	person2.Name = newName2
	editPeople(ctx, client, []clientGRPCJsonXml.Person{*person, *person2})
	getPersonByPhoneNumber(ctx, client, number, phoneType)

	fmt.Println("------ Listing People with at least a number of type HOME --------")
	listPeople(ctx, client, clientGRPCJsonXml.PhoneNumber{Number: number, Kind: phoneType})

	fmt.Println("------- Get People with ID 1 and 2 ---------")
	var id1 int32 = 1
	var id2 int32 = 2
	reqId1 := clientGRPCJsonXml.RequestId{Id: id1}
	reqId2 := clientGRPCJsonXml.RequestId{Id: id2}
	ids := []clientGRPCJsonXml.RequestId{reqId1, reqId2}
	ids[0].Id = id1
	ids[1].Id = id2
	getPeopleById(ctx, client, ids)

}
