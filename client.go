package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"math/rand"
	"time"
)

//func startClient(target string, msgLen int, delayBetweenMsgsMs int, writeBufferSize int, initialWindowSize int32, initialConnWindowSize int32) {
func startClient(target string, msgLen int, delayBetweenMsgsMs int) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	str := buildMessage(msgLen)

	var sum, count int64

	for true {
		time.Sleep(time.Duration(delayBetweenMsgsMs) * time.Millisecond)

		start := time.Now()
		_, err = c.SayHello(context.Background(), &pb.HelloRequest{Name: str})
		if err != nil {
			log.Fatalf("failed to greet: %v", err)
		}
		duration := time.Now().Sub(start)
		sum += duration.Microseconds()
		count++
		log.Printf("Greeting: rate of every %v ms, msgLen = %v, duration = %v, count = %v, avg = %v",
			delayBetweenMsgsMs, msgLen, duration, count, sum/count)
	}
}

func randStringRunes(n int) string {
	var charsArray = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = charsArray[rand.Intn(len(charsArray))]
	}
	return string(b)
}

func buildMessage(len int) string {
	if len == 0 {
		return "hello"
	}
	return randStringRunes(len)
}
