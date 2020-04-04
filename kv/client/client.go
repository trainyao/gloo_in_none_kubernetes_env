package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/trainyao/gloo_in_none_kubernetes_env/kv/kv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


func main() {
	port := flag.Int("port", 8081, "grpc port")
	host := flag.String("host", "0.0.0.0", "grpc host")
	method := flag.String("do", "get", "kv method")

	flag.Parse()

	conn ,err := grpc.Dial(fmt.Sprintf("%s:%d", *host ,*port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := kv.NewKVClient(conn)
	if strings.Index(*method, ":") >= 0 {
		s := strings.SplitN(*method, ":", 2)
		key := s[0]
		value := s[1]
		if resp, err := client.Set(context.Background(), &kv.SetRequest{
			Key:                  key,
			Value:                value,
		}); err != nil {
			panic(err)
		} else {
			log.Print(resp)
		}
	} else {
		key := *method
		if resp, err := client.Get(context.Background(), &kv.GetRequest{
			Key:                  key,
		}); err != nil {
			panic(err)
		} else {
			log.Print(resp)
		}
	}
	panic("do nothing")
	return
}
