package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"shop-api/user-web/forms"
	"shop-api/user-web/global/response"
	"shop-api/user-web/proto/proto"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
				return
			}
		}
	}
}

var (
	conn   *grpc.ClientConn
	client proto.UserClient
)

func init() {
	ip := "127.0.0.1"
	port := 50051
	var err error
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("拨号失败")
	}
	client = proto.NewUserClient(conn)
}

func GetUserList(ctx *gin.Context) {
	page := uint32(1)
	pSize := uint32(10)

	rsp, err := client.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    page,
		PSize: pSize,
	})
	if err != nil {
		zap.S().Errorw(err.Error())
		GrpcErrorToHttp(err, ctx)
		return
	}

	data := make([]response.User, 0)
	for _, item := range rsp.Data {
		tmp := response.User{
			Id:       item.Id,
			Mobile:   item.Mobile,
			NickName: item.NickName,
			BirthDay: response.JsonTime(time.Unix(int64(item.BirthDay), 0)),
			Gender:   item.Gender,
			Role:     item.Role,
		}
		data = append(data, tmp)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total": rsp.Total,
		"data":  data,
		"msg":   "success",
	})
}

func PasswordLogin(c *gin.Context) {
	var passwordLoginForm forms.PasswordLoginForm

	err := c.ShouldBind(&passwordLoginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}
