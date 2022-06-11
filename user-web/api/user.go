package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"net/http"
	"shop-api/user-web/forms"
	"shop-api/user-web/global"
	"shop-api/user-web/global/response"
	"shop-api/user-web/middlewares"
	"shop-api/user-web/models"
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

func GetUserList(ctx *gin.Context) {
	page := uint32(1)
	pSize := uint32(10)

	rsp, err := global.UserClient.GetUserList(context.Background(), &proto.PageInfo{
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
	var form forms.PasswordLoginForm

	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	if !store.Verify(form.CaptchaId, form.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	userInfo, err := global.UserClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: form.Mobile})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	check, err := global.UserClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
		PassWord:          form.Password,
		EncryptedPassWord: userInfo.PassWord,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	if !check.Status {
		c.JSON(http.StatusOK, gin.H{
			"msg": "密码错误",
		})
	}

	claims := models.CustomClaims{
		ID:             userInfo.Id,
		NickName:       userInfo.NickName,
		AuthorityId:    userInfo.Role,
		StandardClaims: jwt.StandardClaims{},
	}

	j := middlewares.NewJwt()
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userId": userInfo.Id,
	})
	return
}

func Register(c *gin.Context) {
	//用户注册
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		GrpcErrorToHttp(err, c)
		return
	}

	//验证码
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	value, err := rdb.Get(context.Background(), registerForm.Mobile).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "验证码错误",
		})
		return
	} else {
		if value != registerForm.Code {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "验证码错误",
			})
			return
		}
	}

	user, err := global.UserClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		PassWord: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
	})

	if err != nil {
		zap.S().Errorf("[Register] 查询 【新建用户失败】失败: %s", err.Error())
		GrpcErrorToHttp(err, c)
		return
	}

	j := middlewares.NewJwt()
	claims := models.CustomClaims{
		ID:          user.Id,
		NickName:    user.NickName,
		AuthorityId: user.Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "imooc",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.Id,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

func GetUserDetail(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	rsp, err := global.UserClient.GetUserById(context.Background(), &proto.IdRequest{
		Id: currentUser.ID,
	})
	if err != nil {
		GrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":     rsp.NickName,
		"birthday": time.Unix(int64(rsp.BirthDay), 0).Format("2006-01-02"),
		"gender":   rsp.Gender,
		"mobile":   rsp.Mobile,
	})
}

func UpdateUser(ctx *gin.Context) {
	updateUserForm := forms.UpdateUserForm{}
	if err := ctx.ShouldBind(&updateUserForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	//将前端传递过来的日期格式转换成int
	loc, _ := time.LoadLocation("Local") //local的L必须大写
	birthDay, _ := time.ParseInLocation("2006-01-02", updateUserForm.Birthday, loc)
	_, err := global.UserClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       (currentUser.ID),
		NickName: updateUserForm.Name,
		Gender:   updateUserForm.Gender,
		BirthDay: uint32(birthDay.Unix()),
	})
	if err != nil {
		GrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
