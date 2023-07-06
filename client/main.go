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

func makeRequestUpdate(request *rank.UpdatePlayerRankInfoReq) *rank.UpdatePlayerRankInfoRes {

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

func makeRequestQueryRank(request *rank.QueryPlayerRankReq) *rank.QueryPlayerRankRes {

	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8080/query/rank", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &rank.QueryPlayerRankRes{}
	proto.Unmarshal(respBytes, respObj)
	return respObj

}

func makeRequestQueryScore(request *rank.QueryPlayerScoreReq) *rank.QueryPlayerScoreRes {

	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8080/query/score", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &rank.QueryPlayerScoreRes{}
	proto.Unmarshal(respBytes, respObj)
	return respObj

}

func makeRequestQueryTotal(request *rank.QueryPlayerNumberReq) *rank.QueryPlayerNumberRes {

	// req, err := proto.Marshal(request)
	// if err != nil {
	// 	log.Fatalf("Unable to marshal request : %v", err)
	// }

	resp, err := http.Get("http://0.0.0.0:8080/query/total")
	// resp, err := http.Post("http://0.0.0.0:8080/query/score", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &rank.QueryPlayerNumberRes{}
	proto.Unmarshal(respBytes, respObj)
	return respObj

}

func makeRequestQueryTop5Rank(request *rank.QueryTop5RankReq) *rank.QueryTop5RankRes {
	resp, err := http.Get("http://0.0.0.0:8080/query/top5")

	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &rank.QueryTop5RankRes{}
	proto.Unmarshal(respBytes, respObj)
	return respObj
}

func makeRequestDelete(request *rank.DeletePlayerRankReq) *rank.DeletePlayerRankRes {
	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8080/delete", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &rank.DeletePlayerRankRes{}
	proto.Unmarshal(respBytes, respObj)
	return respObj
}

func makeRequestClear(request *rank.ClearRankInofReq) *rank.ClearRankInofRes {
	resp, err := http.Get("http://0.0.0.0:8080/clear")

	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &rank.ClearRankInofRes{}
	proto.Unmarshal(respBytes, respObj)
	return respObj
}

func main() {

	request := &rank.UpdatePlayerRankInfoReq{Roleid: 21, Score: 100}
	resp := makeRequestUpdate(request)
	fmt.Printf("Update Player sucessfully and the Return Code from API is : %v\n", resp.RetCode)

	request2 := &rank.UpdatePlayerRankInfoReq{Roleid: 22, Score: 200}
	resp2 := makeRequestUpdate(request2)
	fmt.Printf("Update Player sucessfully and the Return Code from API is : %v\n", resp2.RetCode)

	request3 := &rank.QueryPlayerRankReq{Roleid: 21}
	resp3 := makeRequestQueryRank(request3)
	fmt.Printf("Query Rank for this player is : %v\n", resp3.Info.Rank)

	request5 := &rank.QueryPlayerScoreReq{Roleid: 21}
	resp5 := makeRequestQueryScore(request5)
	fmt.Printf("Query Score for this player is : %v\n", resp5.Info.Rank)

	request4 := &rank.QueryPlayerRankReq{Roleid: 22}
	resp4 := makeRequestQueryRank(request4)
	fmt.Printf("Query Rank for this player is : %v\n", resp4.Info.Rank)

	request6 := &rank.QueryPlayerScoreReq{Roleid: 22}
	resp6 := makeRequestQueryScore(request6)
	fmt.Printf("Query Score for this player is : %v\n", resp6.Info.Rank)

	request7 := &rank.QueryPlayerNumberReq{}
	resp7 := makeRequestQueryTotal(request7)
	fmt.Printf("The total number of this rank list is : %v\n", resp7.Total)

	//add 4 more players
	request8 := &rank.UpdatePlayerRankInfoReq{Roleid: 23, Score: 300}
	resp8 := makeRequestUpdate(request8)
	fmt.Printf("Update Player sucessfully and the Return Code from API is : %v\n", resp8.RetCode)

	request9 := &rank.UpdatePlayerRankInfoReq{Roleid: 24, Score: 400}
	resp9 := makeRequestUpdate(request9)
	fmt.Printf("Update Player sucessfully and the Return Code from API is : %v\n", resp9.RetCode)

	request10 := &rank.UpdatePlayerRankInfoReq{Roleid: 25, Score: 500}
	resp10 := makeRequestUpdate(request10)
	fmt.Printf("Update Player sucessfully and the Return Code from API is : %v\n", resp10.RetCode)

	request11 := &rank.UpdatePlayerRankInfoReq{Roleid: 26, Score: 600}
	resp11 := makeRequestUpdate(request11)
	fmt.Printf("Update Player sucessfully and the Return Code from API is : %v\n", resp11.RetCode)

	//request for top5
	request12 := &rank.QueryTop5RankReq{}
	resp12 := makeRequestQueryTop5Rank(request12)
	fmt.Printf("The top 5 of this rank list is %vand the return code is %v\n", resp12.Info, resp12.RetCode)

	//delete member 26
	request13 := &rank.DeletePlayerRankReq{Roleid: 26}
	resp13 := makeRequestDelete(request13)
	fmt.Printf("This member has been removed from the rank list successfully and the return code is : %v\n", resp13.RetCode)

	//request for top5
	request14 := &rank.QueryTop5RankReq{}
	resp14 := makeRequestQueryTop5Rank(request14)
	fmt.Printf("The top 5 of this rank list is %vand the return code is %v\n", resp14.Info, resp14.RetCode)

	//request for clear the rank list
	request15 := &rank.ClearRankInofReq{}
	resp15 := makeRequestClear(request15)
	fmt.Printf("Already clear the rank list and the return code is %v\n", resp15.RetCode)

	//request for top5
	request16 := &rank.QueryTop5RankReq{}
	resp16 := makeRequestQueryTop5Rank(request16)
	fmt.Printf("The top 5 of this rank list is %vand the return code is %v\n", resp16.Info, resp16.RetCode)
}
