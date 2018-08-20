package main

import (
	"api"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:8080", grpc.WithInsecure())

	defer conn.Close()
	client := api.NewGreetingClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	r, _ := client.SayHello(ctx, &api.Request{RequestApi: "Janshair"})

	fmt.Println(r.ResponseApi)
}
