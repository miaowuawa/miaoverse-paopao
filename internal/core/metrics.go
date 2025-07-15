// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

type TweetMetricServantA interface {
	UpdateTweetMetric(metric *cs.TweetMetric) error
	AddTweetMetric(postId int64) error
	DeleteTweetMetric(postId int64) error
}

type CommentMetricServantA interface {
	UpdateCommentMetric(metric *cs.CommentMetric) error
	AddCommentMetric(commentId int64) error
	DeleteCommentMetric(commentId int64) error
	GetCommentMetric(commentId int64) (*cs.CommentMetric, error)
}

type UserMetricServantA interface {
	UpdateUserMetric(userId int64, action uint8) error
	AddUserMetric(userId int64) error
	DeleteUserMetric(userId int64) error
	UpdateUserExperience(userId int64, Experience int) error
	GetUserMetric(userId int64) (*cs.UserMetric, error)
}
