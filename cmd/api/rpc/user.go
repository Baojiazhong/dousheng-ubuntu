package rpc

import (
	"context"
	"time"

	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo/userservice"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

// QueryUser query user info
func QueryUser(ctx context.Context, req *userdemo.CheckUserRequest) (int64, errno.ErrNo) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, errno.ConvertErr(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	// must return success
	return resp.UserId, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
}
