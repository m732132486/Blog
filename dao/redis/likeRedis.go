package redis

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

// GetLikeCount 获取点赞计数
func GetLikeCount(PostID string) (int64, error) {
	key := "likes:" + PostID + ":users"
	//创建pipeline
	pipe := Client.Pipeline()
	pipe.ZCard(ctx, key)
	//执行pipeline
	cmd, err := pipe.Exec(ctx)
	if err != nil {
		zap.L().Error("pipe.Exec(ctx)", zap.Error(err))
		return 0, err
	}
	//解析pipeline
	for _, cmd := range cmd {
		if countCmd, ok := cmd.(*redis.IntCmd); ok {
			//获取点赞数量并返回
			count, err := countCmd.Result()
			if err != nil {
				zap.L().Error("countCmd.Result() err", zap.Error(err))
				return 0, nil
			}
			return count, nil
		}

	}
	return 0, errors.New("获取点赞数量失败")
}

// RecordLike 记录点赞操作
func RecordLike(PostID, UserID string) error {
	key := "likes:" + PostID + ":users"
	pipe := Client.Pipeline()
	pipe.ZAdd(ctx, key, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: UserID,
	})
	//执行pipeline命令
	_, err := pipe.Exec(ctx)
	if err != nil {
		zap.L().Error("pipe.Exec(ctx)", zap.Error(err))
		return err
	}
	return nil
}

// IsLiked 检查用户是否已点过点赞
func IsLiked(PostID, UserID string) (bool, error) {
	// 构造 Redis 键
	key := "likes:" + PostID + ":users"

	// 获取指定用户在有序集合中的排名
	rankCmd := Client.ZRank(ctx, key, UserID)
	rank, err := rankCmd.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// 如果用户没有在有序集合中，说明用户未点赞过
			return false, nil
		}
		zap.L().Error("ZRank", zap.Error(err))
		return false, err
	}

	// 如果排名存在，则用户已经点赞过
	return rank >= 0, nil
}

// CancelLike 取消点赞操作
func CancelLike(PostID, UserID string) error {
	// 构造 Redis 键
	key := "likes:" + PostID + ":users"
	// 创建 Pipeline
	pipe := Client.TxPipeline()

	// 向 Pipeline 中添加 ZREM 命令
	pipe.ZRem(ctx, key, UserID)

	// 执行 Pipeline 命令
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

//func likeCount(PostID string) (int64, error) {
//	// 获取当前点赞数
//	count, err := GetLikeCount(PostID)
//	if err != nil {
//		zap.L().Error("redis.GetLikeCount(parseInt) ", zap.String("postID", PostID))
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("帖子%s有%d个赞", postID, count)})
//}
