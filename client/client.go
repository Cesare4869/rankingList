package main

import (
	"bytes"
	"fmt"
	"log"

	"io/ioutil"
	"net/http"

	"github.com/Cesare4869/rankingList/proto"
	proto3 "github.com/golang/protobuf/proto"
)

func makeRequest(request *proto.UpdatePlayerRankInfoReq) *proto.UpdatePlayerRankInfoRes {

	req, err := proto3.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8080/update", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &proto.UpdatePlayerRankInfoRes{}
	proto3.Unmarshal(respBytes, respObj)
	return respObj
}

func main() {

	request := &proto.UpdatePlayerRankInfoReq{Roleid: 21, Score: 100}
	resp := makeRequest(request)
	fmt.Print("Response from API is :&v\n", resp.GetRetCode())

}
