package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
	"time"

	"github.com/trainyao/gloo_in_none_kubernetes_env/kv/kv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 8081, "grpc port")
	host := flag.String("host", "0.0.0.0", "grpc host")
	method := flag.String("do", "get", "kv method")
	header := flag.String("header", "train123", "header")

	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *host, *port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	client := kv.NewKVClient(conn)
	for {
		if strings.Index(*method, ":") >= 0 {
			s := strings.SplitN(*method, ":", 2)
			key := s[0]
			value := s[1]
			if resp, err := client.Set(context.Background(), &kv.SetRequest{
				Key:   key,
				Value: value,
			}); err != nil {
				panic(err)
			} else {
				log.Print(resp)
			}
		} else {
			key := *method
			req := &kv.GetRequest{
				Key: key,
			}

			h := metadata.New(map[string]string{
				"train": *header,
			})

			ctx := metadata.NewOutgoingContext(context.Background(), h)

			if resp, err := client.Get(ctx, req); err != nil {
				panic(err)
			} else {
				log.Print(resp)
			}
		}

		time.Sleep(1 * time.Second)
	}

}
