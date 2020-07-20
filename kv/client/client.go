

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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	port := flag.Int("port", 8081, "grpc port")
	host := flag.String("host", "0.0.0.0", "grpc host")
	method := flag.String("do", "key", "--do key == get 'key'; or --do key:value == set 'key: value'")
	header := flag.String("header", "-", "header")

	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *host, *port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer func() {
		if err := recover(); err != nil {
			// ouch!
			// lets print the gRPC error message
			// which is "Length of `Name` cannot be more than 10 characters"
			errStatus, _ := status.FromError(err.(error))
			fmt.Println(errStatus.Message())
			// lets print the error code which is `INVALID_ARGUMENT`
			fmt.Println(errStatus.Code())
			// Want its int version for some reason?
			// you shouldn't actullay do this, but if you need for debugging,
			// you can do `int(status_code)` which will give you `3`
			//
			// Want to take specific action based on specific error?
			if codes.InvalidArgument == errStatus.Code() {
				// do your stuff here
				log.Fatal()
			}
		}
	}()
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
				"Header": *header,
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

