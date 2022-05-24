package db

import (
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
	"gorm.io/gorm"

	"context"
)

type User struct {
	gorm.Model
	ID            int64  `gorm:"primarykey" json:"user_id"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"` // md5加密后的密码
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// // MGetUsers multiple get list of user info
// func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
// 	res := make([]*User, 0)
// 	if len(userIDs) == 0 {
// 		return res, nil
// 	}

// 	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
