package handler

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"regexp"
	"strconv"
	"testProject/microservice/account_srv/proto/pb"
	"testProject/microservice/account_web/req"
	"testProject/microservice/account_web/res"
	"testProject/microservice/custom_error"
	"testProject/microservice/jwt_op"
	"testProject/microservice/log"
	"time"
)

func HandleError(err error) string {
	if err != nil {
		switch err.Error() {
		case custom_error.AccountNotFound:
			return custom_error.AccountNotFound
		case custom_error.AccountPasswordError:
			return custom_error.AccountPasswordError
		case custom_error.AccountIsExist:
			return custom_error.AccountIsExist
		case custom_error.SaltError:
			return custom_error.SaltError
		default:
			return custom_error.InternalError
		}
	}
	return ""
}

func AccountListHandler(ctx *gin.Context) {
	pageNoStr := ctx.DefaultQuery("pageNo", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "3")
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		s := fmt.Sprintf("AccountListHandler-GRPC拨号失败: %v", err)
		log.Logger.Info(s)
		e := HandleError(err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}

	pageNo, _ := strconv.ParseInt(pageNoStr, 10, 32)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 32)

	client := pb.NewAccountServiceClient(conn)
	accountRes, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   int32(pageNo),
		PageSize: int32(pageSize),
	})
	if err != nil {
		s := fmt.Sprintf("GetAccountList调用失败: %v", err.Error())
		log.Logger.Info(s)
		e := HandleError(err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}
	var resList []res.Account4Res
	for _, item := range accountRes.AccountList {
		resList = append(resList, Pb2Res(item))
	}

	log.Logger.Info("AccountListHandler调试通过")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "",
		"total": accountRes.Total,
		"data":  resList,
	})
}

func Pb2Res(accountRes *pb.AccountRes) res.Account4Res {
	return res.Account4Res{
		Mobile:   accountRes.Mobile,
		NickName: accountRes.Nickname,
		Gender:   accountRes.Gender,
	}
}

func LoginByPasswordHandler(ctx *gin.Context) {
	var loginByPassword req.LoginByPassword
	err := ctx.ShouldBind(&loginByPassword)
	if err != nil {
		log.Logger.Error("LoginByPassword出错，" + err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "解析参数错误",
		})
		return
	}
	// TODO 校验手机号码格式
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !re.MatchString(loginByPassword.Mobile) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "号码格式不匹配",
		})
		return
	}

	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		log.Logger.Error("LoginByPassword 请求服务端出错，" + err.Error())
		e := HandleError(err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}

	client := pb.NewAccountServiceClient(conn)
	accountRes, err := client.GetAccountByMobile(ctx, &pb.MobileRequest{
		Mobile: loginByPassword.Mobile,
	})
	if err != nil {
		log.Logger.Error("LoginBP GRPC GetByID出错，" + err.Error())
		e := HandleError(err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}
	result, err := client.CheckPassword(ctx, &pb.CheckPasswordRequest{
		Password:       loginByPassword.Passwd,
		HashedPassword: accountRes.Password,
		AccountId:      accountRes.Id,
	})
	if err != nil {
		log.Logger.Error("LoginBP GRPC CheckPassword出错，" + err.Error())
		e := HandleError(err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}
	if !result.Result {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":    "",
			"result": "登录失败",
			"token":  "",
		})
		return
	}
	j := jwt_op.NewJwt()
	now := time.Now()
	token, err := j.GenerateToken(jwt_op.CustomClaims{
		ID:       accountRes.Id,
		NickName: accountRes.Nickname,
		StandardClaims: jwt.StandardClaims{
			NotBefore: now.Unix(),
			ExpiresAt: now.Add(30 * 24 * time.Hour).Unix(),
		},
		AuthorityId: accountRes.Role,
	})
	ctx.Set("token", token)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "",
		"result": "登录成功",
		"token":  token,
	})
}
