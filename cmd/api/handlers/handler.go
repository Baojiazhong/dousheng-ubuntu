package handlers

import (
	"net/http"
	"time"

	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int32  `json:"status_code"`
	Message string `json:"status_msg"`
}

type Response struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"data"`
}

type LoginResponse struct {
	Code    int32  `json:"status_code"`
	Message string `json:"status_msg"`
	UserId  int64  `json:"user_id"`
	Token   string `json:"token"`
	Expire  string `json:"expire"`
}

type UserInfoResponse struct {
	Code          int32  `json:"status_code"`
	Message       string `json:"status_msg"`
	UserId        int64  `json:"user_id"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserListResponse struct {
	Code     int32            `json:"status_code"`
	Message  string           `json:"status_msg"`
	UserList []*userdemo.User `json:"user_list"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

// SendLoginResponse pack response
func SendLoginResponse(c *gin.Context, err error, userId int64, token string, expire time.Time) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, LoginResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		UserId:  userId,
		Token:   token,
		Expire:  expire.Format(time.RFC3339),
	})
}

// SendUserInfoResponse pack response
func SendUserInfoResponse(c *gin.Context, err error, userId int64, followCount int64, followerCount int64, isFollow bool) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserInfoResponse{
		Code:          Err.ErrCode,
		Message:       Err.ErrMsg,
		UserId:        userId,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	})
}

// SendUserListResponse
func SendUserListResponse(c *gin.Context, err error, userList []*userdemo.User) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserListResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		UserList: userList,
	})
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
