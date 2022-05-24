package handlers

import (
	"net/http"

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
func SendLoginResponse(c *gin.Context, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, LoginResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		UserId:  userId,
		Token:   token,
	})
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
