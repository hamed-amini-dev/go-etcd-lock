package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.uber.org/zap"
)

func main() {
	var name = flag.String("name", "", "give a name")
	flag.Parse()

	// Create an etcd client with increased timeout and logging
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379", "http://localhost:2380"},
		DialTimeout: 5 * time.Second,
		LogConfig: &zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
			Development: true,
			Encoding:    "json",
			OutputPaths: []string{"stdout"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// Test connectivity with a simple operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = cli.Put(ctx, "test-key", "test-value")
	if err != nil {
		log.Fatal(err)
	}

	// Create a session to acquire a lock
	s, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	l := concurrency.NewMutex(s, "/distributed-lock/")
	ctx = context.Background()

	// Acquire lock (or wait to have it)
	if err := l.Lock(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for", *name)
	fmt.Println("Do some work in", *name)
	time.Sleep(15 * time.Second)

	// Release lock
	if err := l.Unlock(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for", *name)
}
