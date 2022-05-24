package db

import (
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
	"gorm.io/gorm"
)

type Follower struct {
	gorm.Model
	UserID     int64 `json:"user_id"`
	FollowerID int64 `json:"follower_id"`
}

func (u *Follower) TableName() string {
	return constants.FollowerTableName
}
