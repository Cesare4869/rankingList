package rankList

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	_ "strings"
	_ "time"

	"github.com/Cesare4869/rankingList/rank"
	proto3 "github.com/golang/protobuf/proto"
	"github.com/redis/go-redis/v9"
)

type RankingList struct {
	Redis *redis.Client
	Key   string
}

type RankMember struct {
	UID  int64
	Val  int64
	Rank int64
}

func New(rds *redis.Client, key string) (*RankingList, error) {
	return &RankingList{
		Redis: rds,
		Key:   key,
	}, nil
}

func (h *RankingList) MyhandleFuncClear(resp http.ResponseWriter, req *http.Request) {
	// contentLenth := req.ContentLength
	// fmt.Printf("Content Length Received : %v\n", contentLenth)

	ctx := context.TODO()
	err := h.ClearRankingList(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Clear the ranking List successfully!")
	result := &rank.ClearRankInofRes{RetCode: 0}
	response, err := proto3.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	resp.Write(response)
}

func (h *RankingList) MyhandlerFuncDelete(resp http.ResponseWriter, req *http.Request) {
	// contentLenth := req.ContentLength
	// fmt.Printf("Content Length Received : %v\n", contentLenth)
	request := &rank.DeletePlayerRankReq{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	proto3.Unmarshal(data, request)

	roleid := request.GetRoleid()
	ctx := context.TODO()
	h.RemoveMemberFromRankingList(ctx, roleid)
	result := &rank.DeletePlayerRankRes{RetCode: 0}
	response, err := proto3.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	fmt.Println("Delete Player Rank Info Successfully!")
	resp.Write(response)
}

func (h *RankingList) MyHanlderFuncQueryTop5(resp http.ResponseWriter, req *http.Request) {
	// contentLenth := req.ContentLength
	// fmt.Printf("Content Length Received : %v\n", contentLenth)

	ctx := context.TODO()
	list, err := h.GetRankingList(ctx, 5, true)
	if err != nil {
		log.Fatalf("Uable to get the top 5 rank list : %v", err)
	}

	result := &rank.QueryTop5RankRes{
		Info:    []*rank.PlayerInfo{}, // 创建空的 PlayerInfo 列表
		RetCode: 0,
	}

	// 循环添加多个 PlayerInfo 对象到重复参数 info
	for _, z := range list {
		playerInfo := &rank.PlayerInfo{
			Roleid: int64(z.UID),
			Score:  int64(z.Val),
			Rank:   int64(z.Rank),
		}
		result.Info = append(result.Info, playerInfo)
	}
	fmt.Printf("Querying for top5 in the ranklist, and the result is %v\n", result)
	response, err := proto3.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	resp.Write(response)
}

func (h *RankingList) MyHanlderFuncQueryTotal(resp http.ResponseWriter, req *http.Request) {
	// contentLenth := req.ContentLength
	// fmt.Printf("Content Length Received : %v\n", contentLenth)

	ctx := context.TODO()
	total := h.GetTotalCount(ctx)
	fmt.Printf("Querying for the total number and the total number is %v\n", total)
	result := &rank.QueryPlayerNumberRes{Total: total}
	response, err := proto3.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	resp.Write(response)
}

func (h *RankingList) MyHanlderFuncQueryScore(resp http.ResponseWriter, req *http.Request) {
	// contentLenth := req.ContentLength
	// fmt.Printf("Content Length Received : %v\n", contentLenth)
	request := &rank.QueryPlayerScoreReq{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	proto3.Unmarshal(data, request)

	roleid := request.GetRoleid()
	ctx := context.TODO()
	ranking, err := h.GetUserRank(ctx, roleid, true)
	score, err := h.GetUserVal(ctx, roleid)
	if err != nil {
		fmt.Printf("get user rank, err:%v\n", err)
		return
	}
	player := &rank.PlayerInfo{
		Roleid: roleid,
		Score:  score,
		Rank:   ranking,
	}

	fmt.Printf("Querying for the player score and the Player %v's score is %v\n", roleid, score)
	result := &rank.QueryPlayerScoreRes{Info: player, RetCode: 0} // 这里要不要返回Player的所有信息，还是只是返回Player的ranking信息
	response, err := proto3.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	resp.Write(response)
}

func (h *RankingList) MyHanlderFuncQueryRank(resp http.ResponseWriter, req *http.Request) {
	// contentLenth := req.ContentLength
	// fmt.Printf("Content Length Received : %v\n", contentLenth)
	request := &rank.QueryPlayerRankReq{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	proto3.Unmarshal(data, request)

	roleid := request.GetRoleid()
	ctx := context.TODO()
	ranking, err := h.GetUserRank(ctx, roleid, true)
	score, err := h.GetUserVal(ctx, roleid)
	if err != nil {
		fmt.Printf("get user rank, err:%v\n", err)
		return
	}
	player := &rank.PlayerInfo{
		Roleid: roleid,
		Score:  score,
		Rank:   ranking,
	}

	fmt.Printf("Querying for the player rank and the Player %v's rank is %v\n", roleid, ranking)
	result := &rank.QueryPlayerRankRes{Info: player, RetCode: 0} // 这里要不要返回Player的所有信息，还是只是返回Player的ranking信息
	response, err := proto3.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	resp.Write(response)
}

func (h *RankingList) MyHanlderFuncUpdate(resp http.ResponseWriter, req *http.Request) {
	// contentLenth := req.ContentLength
	// fmt.Printf("Content Length Received : %v\n", contentLenth)
	request := &rank.UpdatePlayerRankInfoReq{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	proto3.Unmarshal(data, request)

	roleid := request.GetRoleid()
	score := request.GetScore()
	ctx := context.TODO()
	h.Update(ctx, roleid, score)

	result := &rank.UpdatePlayerRankInfoRes{RetCode: 0}
	response, err := proto3.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	fmt.Println("Update Player Info Successfully!")
	resp.Write(response)
}

func (r *RankingList) Update(ctx context.Context, uid, val int64) (int64, error) {
	keys := []string{r.Key}
	args := []interface{}{uid, val}
	zincrby := redis.NewScript(`
-- 排行榜key
local key = KEYS[1]
-- 要更新的用户id
local uid = ARGV[1]
-- 用户本次新增的val（小数位为时间差标识）
local valScore = ARGV[2]

-- 更新用户最新的score信息（累计val.最新时间差）
local newScore = valScore
redis.call("ZADD", key, newScore, uid)

-- 更新成功返回newScore（注意要使用tostring才能返回小数）
return newScore
	`)
	newScore, err := zincrby.Run(ctx, r.Redis, keys, args...).Int64()
	if err != nil {
		panic(err)
	}
	return newScore, nil
}

func (r *RankingList) GetRankingList(ctx context.Context, topN int64, desc bool) ([]RankMember, error) {
	start := int64(0)
	stop := topN - 1
	if topN <= 0 {
		stop -= 1
	}
	total := r.GetTotalCount(ctx)
	if stop >= total {
		stop = total - 1
	}
	zrange := r.Redis.ZRangeWithScores
	if desc {
		zrange = r.Redis.ZRevRangeWithScores
	}
	list, err := zrange(ctx, r.Key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	result := []RankMember{}
	for idx, z := range list {
		val := int64(z.Score)
		member := z.Member.(string)
		uid, err := strconv.ParseInt(member, 10, 64)
		if err != nil {
			panic(err)
		}
		m := RankMember{
			UID:  uid,
			Val:  val,
			Rank: int64(idx + 1),
		}
		result = append(result, m)
	}
	return result, nil
}

func (r *RankingList) GetUserVal(ctx context.Context, uid int64) (int64, error) {
	score, err := r.Redis.ZScore(ctx, r.Key, fmt.Sprint(uid)).Result()
	if err != nil {
		panic(err)
	}
	return int64(score), nil
}

func (r *RankingList) GetUserRank(ctx context.Context, uid int64, desc bool) (int64, error) {
	zrank := r.Redis.ZRank
	if desc {
		zrank = r.Redis.ZRevRank
	}
	idx, err := zrank(ctx, r.Key, fmt.Sprint(uid)).Result()
	if err != nil {
		panic(err)
	}
	rank := idx + 1
	return rank, err
}

func (r *RankingList) RemoveMemberFromRankingList(ctx context.Context, roleID int64) error {
	_, err := r.Redis.ZRem(ctx, r.Key, roleID).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *RankingList) ClearRankingList(ctx context.Context) error {
	_, err := r.Redis.Del(ctx, r.Key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *RankingList) GetTotalCount(ctx context.Context) int64 {
	return r.Redis.ZCard(ctx, r.Key).Val()
}
