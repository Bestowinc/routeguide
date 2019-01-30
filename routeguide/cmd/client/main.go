package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ihcsim/grpc-101/routeguide"
	pb "github.com/ihcsim/grpc-101/routeguide/proto"

	"google.golang.org/grpc"
)

func main() {
	var (
		addr       = ""
		port       = "8080"
		opts       = []grpc.DialOption{grpc.WithInsecure()}
		ctxTimeout = time.Second * 20
		client     = routeguide.Client{}
	)

	override, exist := os.LookupEnv("SERVER_HOST")
	if exist {
		addr = override
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", addr, port), opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	grpcClient := pb.NewRouteGuideClient(conn)
	client.GRPC = grpcClient

	for {
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		// randomly pick one of the 4 APIs
		n := rand.Intn(10)
		if n < 3 {
			if err := client.GetFeature(ctx); err != nil {
				log.Println(err)
			}
		} else if n < 5 && n >= 3 {
			if err := client.ListFeatures(ctx); err != nil {
				log.Println(err)
			}
		} else if n < 7 && n >= 5 {
			if err := client.RecordRoute(ctx); err != nil {
				log.Println(err)
			}
		} else {
			if err := client.RouteChat(ctx); err != nil {
				log.Println(err)
			}
		}

		time.Sleep(time.Second * 1)
	}
}
