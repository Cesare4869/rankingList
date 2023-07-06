package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	rankList "github.com/Cesare4869/rankingList"
	"github.com/Cesare4869/rankingList/rank"
	proto3 "github.com/golang/protobuf/proto"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

// No Need, MyHanlderFunc is moved to rankList.go
func (h *Handler) MyHanlderFunc(resp http.ResponseWriter, req *http.Request) {
	contentLenth := req.ContentLength
	fmt.Printf("Content Length Received : %v\n", contentLenth)
	request := &rank.UpdatePlayerRankInfoReq{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	proto3.Unmarshal(data, request)
	roleid := request.GetRoleid()
	score := request.GetScore()
	fmt.Println(roleid, score)
	ctx := context.Background()
	err = h.RedisClient.Set(ctx, "score2", 200, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	result := &rank.UpdatePlayerRankInfoRes{RetCode: 1}
	response, err := proto3.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	resp.Write(response)
}

type Handler struct {
	RedisClient *redis.Client
} //暂时不需要，使用rankList定义type

func main() {

	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", //localhost:6379 for local, redis-container:6379 for docker
	})
	ctx := context.TODO()
	key := "rankList:example"
	defer rds.Del(ctx, key)

	// handler := &Handler{
	// 	RedisClient: rds,
	// }

	rankList, err := rankList.New(rds, key)
	if err != nil {
		panic(err)
	}

	// API Server
	fmt.Println("Starting the API Server")
	r := mux.NewRouter()
	r.HandleFunc("/update", rankList.MyHanlderFuncUpdate).Methods("POST")
	r.HandleFunc("/query/rank", rankList.MyHanlderFuncQueryRank).Methods("POST")
	r.HandleFunc("/query/score", rankList.MyHanlderFuncQueryScore).Methods("POST")
	r.HandleFunc("/query/total", rankList.MyHanlderFuncQueryTotal).Methods("GET")
	r.HandleFunc("/query/top5", rankList.MyHanlderFuncQueryTop5).Methods("GET")
	r.HandleFunc("/delete", rankList.MyhandlerFuncDelete).Methods("POST")
	r.HandleFunc("/clear", rankList.MyhandleFuncClear).Methods("GET")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
