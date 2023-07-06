package rankList

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRankList(t *testing.T) {
	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.TODO()
	key := "rankList:test"
	defer rds.Del(ctx, key)

	var uid int64 = 21
	rankList, err := New(rds, key)
	require.Nil(t, err)

	score, err := rankList.Update(ctx, uid, 100)
	require.Nil(t, err)

	score1, err := rankList.Update(ctx, uid-1, 200)
	require.Nil(t, err)

	score2, err := rankList.Update(ctx, uid+1, 50)
	require.Nil(t, err)
	t.Logf("score:%v, score1:%v, score2:%v", score, score1, score2)

	total := rankList.GetTotalCount(ctx)
	require.Equal(t, int64(3), total)

	list, err := rankList.GetRankingList(ctx, 3, true)
	require.Nil(t, err)
	require.Len(t, list, 3)
	require.Equal(t, uid-1, list[0].UID)
	require.Equal(t, uid, list[1].UID)
	require.Equal(t, uid+1, list[2].UID)

	list, err = rankList.GetRankingList(ctx, 4, true)
	require.Nil(t, err)
	require.Len(t, list, 3)

	list, err = rankList.GetRankingList(ctx, 2, true)
	require.Nil(t, err)
	require.Len(t, list, 2)

	list, err = rankList.GetRankingList(ctx, 3, false)
	require.Nil(t, err)
	require.Len(t, list, 3)
	require.Equal(t, uid-1, list[2].UID)
	require.Equal(t, uid, list[1].UID)
	require.Equal(t, uid+1, list[0].UID)

	ur, err := rankList.GetUserRank(ctx, uid, true)
	require.Nil(t, err)
	require.Equal(t, int64(2), ur)
	uv, err := rankList.GetUserVal(ctx, uid)
	require.Nil(t, err)
	require.Equal(t, int64(100), uv)

	ur, err = rankList.GetUserRank(ctx, uid+1, true)
	require.Nil(t, err)
	require.Equal(t, int64(3), ur)
	uv, err = rankList.GetUserVal(ctx, uid+1)
	require.Nil(t, err)
	require.Equal(t, int64(50), uv)

	ur, err = rankList.GetUserRank(ctx, uid-1, false)
	require.Nil(t, err)
	require.Equal(t, int64(3), ur)
	uv, err = rankList.GetUserVal(ctx, uid-1)
	require.Nil(t, err)
	require.Equal(t, int64(200), uv)

	err = rankList.RemoveMemberFromRankingList(ctx, uid)
	require.Nil(t, err)
	total = rankList.GetTotalCount(ctx)
	require.Equal(t, int64(2), total)

	err = rankList.ClearRankingList(ctx)
	require.Nil(t, err)
	total = rankList.GetTotalCount(ctx)
	require.Equal(t, int64(0), total)

}
