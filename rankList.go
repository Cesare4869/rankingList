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

func (h *RankingList) MyHanlderFunc(resp http.ResponseWriter, req *http.Request) {
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
	err = h.Redis.Set(ctx, "score3", 300, 0).Err()
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

func (r *RankingList) GerUserRank(ctx context.Context, uid int64, desc bool) (int64, error) {
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

func (r *RankingList) GetTotalCount(ctx context.Context) int64 {
	return r.Redis.ZCard(ctx, r.Key).Val()
}
