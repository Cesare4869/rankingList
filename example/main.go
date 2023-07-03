package main

import (
	"context"
	"fmt"

	rankList "github.com/Cesare4869/rankingList"
	"github.com/redis/go-redis/v9"
)

func main() {

	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.TODO()
	key := "rankList:example"
	defer rds.Del(ctx, key)

	rankList, err := rankList.New(rds, key)
	if err != nil {
		panic(err)
	}

	score, err := rankList.Update(ctx, 1, 100)
	if err != nil {
		panic(err)
	}
	fmt.Println(score)

}
