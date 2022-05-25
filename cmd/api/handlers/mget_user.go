package handlers

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/api/rpc"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// get follow list
func GetFollowList(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	req := &userdemo.MGetUserRequest{UserId: userID, ActionType: constants.QueryFollowList}
	user, err := rpc.MGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserListResponse(c, nil, user)
}

// get follower list
func GetFollowerList(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	req := &userdemo.MGetUserRequest{UserId: userID, ActionType: constants.QueryFollowerList}
	user, err := rpc.MGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserListResponse(c, nil, user)
}
