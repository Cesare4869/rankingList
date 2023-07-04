package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Cesare4869/rankingList/rank"
	"github.com/golang/protobuf/proto"
)

func makeRequest(request *rank.UpdatePlayerRankInfoReq) *rank.UpdatePlayerRankInfoRes {

	req, err := proto.Marshal(request)
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

	respObj := &rank.UpdatePlayerRankInfoRes{}
	proto.Unmarshal(respBytes, respObj)
	return respObj

}

func main() {

	request := &rank.UpdatePlayerRankInfoReq{Roleid: 21, Score: 100}
	resp := makeRequest(request)
	fmt.Printf("Response from API is : %v\n", resp.GetMessage())
}
