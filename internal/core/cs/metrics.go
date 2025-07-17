// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

const (
	MetricActionCreateTweet = iota + 1
	MetricActionDeleteTweet
)

type TweetMetric struct {
	PostId          int64
	IncentiveScore  int
	DecayFactor     int
	CommentCount    int64
	UpvoteCount     int64
	CollectionCount int64
	ShareCount      int64
}

func (m *TweetMetric) RankScore(motivationFactor int) int64 {
	if m.DecayFactor == 0 {
		return 0
	}
	return int64(m.IncentiveScore * motivationFactor / m.DecayFactor)
}

type CommentMetric struct {
	CommentId       int64
	IncentiveScore  int
	DecayFactor     int
	ReplyCount      int32
	ThumbsUpCount   int32
	ThumbsDownCount int32
}

func (m *CommentMetric) RankScore(motivationFactor int) int64 {
    if m.DecayFactor == 0 {
        return 0
    }
    return int64(m.IncentiveScore * motivationFactor / m.DecayFactor)
}

// UserMetric 用户指标结构体
type UserMetric struct {
	UserId         int64
	TweetsCount    int
	LatestTrendsOn int64
	Experience     int
}
