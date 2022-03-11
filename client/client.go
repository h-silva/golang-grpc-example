package main

import (
	"context"
	"fmt"
	"log"

	"github.com/h-silva/golang-grpc-example/api/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello World Client")

	conn, err := grpc.Dial(fmt.Sprintf(":%d", 3000), nil)

	if err != nil {
		log.Fatalf("Failed to dial at port %d: %s", 3000, err.Error())
	}

	defer conn.Close()

	newObject := pb.NewFieldServiceClient(conn)

	resp, err := newObject.Create(context.Background(), &pb.FieldRequest{Farm: 1, Zone: 2, Field: 3})

	if err != nil {
		log.Fatal("FUDEU")
	}

	log.Println(resp.Status)

}
