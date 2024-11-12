package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"testProject/microservice/account_srv/proto/pb"
	"testProject/microservice/account_web/res"
	"testProject/microservice/custom_error"
	"testProject/microservice/log"
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

	client := pb.NewAccountServiceClient(conn)
	accountRes, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   1,
		PageSize: 3,
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
