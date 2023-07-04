package main

import (
	"context"
	"fmt"
	"log"

	// "time"
	rankList "github.com/Cesare4869/rankingList"
	"github.com/Cesare4869/rankingList/rank"
	proto3 "github.com/golang/protobuf/proto"
	"github.com/redis/go-redis/v9"
	// "github.com/gorilla/mux"
)

// func Update(resp http.ResponseWriter, req *http.Request) {
// 	contentLenth := req.ContentLength
// 	fmt.Printf("Content Length Received : %v\n", contentLenth)
// 	request := &proto.UpdatePlayerRankInfoReq{}
// 	data, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		log.Fatalf("Unable to read message from request : %v", err)
// 	}
// 	proto3.Unmarshal(data, request)
// 	// roleid := request.GetRoleid()
// 	// score := request.GetScore()
// 	result := &proto.UpdatePlayerRankInfoRes{RetCode: 1}
// 	response, err := proto3.Marshal(result)
// 	if err != nil {
// 		log.Fatalf("Unable to marshal response : %v", err)
// 	}
// 	resp.Write(response)
// }

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
	rankList.Update(ctx, 21, 100)

	req := &rank.UpdatePlayerRankInfoReq{Roleid: 21, Score: 100}
	data, err := proto3.Marshal(req)
	if err != nil {
		log.Fatalf("Error while marshalling the object : %v", err)
	}
	fmt.Println(data)
	// fmt.Println(req)

	res := &rank.UpdatePlayerRankInfoReq{}
	err = proto3.Unmarshal(data, res)
	fmt.Println(res.GetScore())
	if err != nil {
		log.Fatalf("Error while un-marshalling the object : %v", err)
	}
	// fmt.Printf("Value from un-marshalled data is %v", res.GetRoleid())

	// // API Server
	// fmt.Println("Starting the API Server")
	// r := mux.NewRouter()
	// r.HandleFunc("/update", Update).Methods("POST")

	// server := &http.Server{
	// 	Handler:      r,
	// 	Addr:         "0.0.0.0:8080",
	// 	WriteTimeout: 2 * time.Second,
	// 	ReadTimeout:  2 * time.Second,
	// }

	// log.Fatal(server.ListenAndServe())

}
