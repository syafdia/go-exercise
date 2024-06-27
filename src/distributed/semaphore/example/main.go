package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/syafdia/distributed/semaphore"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	sem, err := semaphore.New(20, "example", redisClient, 2*time.Minute)
	if err != nil {
		panic(err)
	}

	totalProceses := 11
	for i := 0; i < totalProceses; i++ {
		n := i
		go runHeavyProcess(sem, n)
	}

	waitForeverC := make(chan bool, 1)
	<-waitForeverC
}

func runHeavyProcess(sem *semaphore.Semaphore, n int) {
	ctx := context.Background()

	resource, err := sem.Acquire(ctx)
	if err != nil {
		fmt.Println("Got error", err)
		return
	}

	fmt.Println("Heavy process started, id:", n)
	time.Sleep(10 * time.Second)
	fmt.Println("Heavy process finished, id:", n)

	err = sem.Release(ctx, resource)
	if err != nil {
		fmt.Println("Got error", err)
		return
	}
}
