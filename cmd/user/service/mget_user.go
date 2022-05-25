package service

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/dal/db"
	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/pack"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *userdemo.MGetUserRequest) ([]*userdemo.User, error) {
	Users, err := db.MGetUsers(s.ctx, req)
	if err != nil {
		return nil, err
	}
	if req.ActionType == constants.QueryFollowList {
		return pack.Users(Users, true), nil
	} else if req.ActionType == constants.QueryFollowerList {
		return pack.Users(Users, false), nil
	} else {
		return pack.Users(Users, false), nil // default
	}
}
