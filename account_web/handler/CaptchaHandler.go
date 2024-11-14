package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"testProject/microservice/internal"
	"testProject/microservice/log"
	"time"
)

func CaptchaHandler(ctx *gin.Context) {
	mobile, ok := ctx.GetQuery("mobile")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	fileName := "microservice/account_web/handler/png/captcha.png"
	f, err := os.Create(fileName)
	if err != nil {
		log.Logger.Error("GenCaptcha() create失败")
		return
	}
	defer f.Close()

	digits := captcha.RandomDigits(captcha.DefaultLen)
	image := captcha.NewImage("", digits, captcha.StdWidth, captcha.StdHeight)
	// image/io的WriteTo需要传入一个Writer接口，所以只要实现了这个接口都可以传入
	_, err = image.WriteTo(f)
	if err != nil {
		log.Logger.Error("GenCaptcha() write失败")
		return
	}

	capt := ""
	for _, item := range digits {
		capt += string(item)
	}
	fmt.Println("capt: " + capt)

	internal.RedisClient.Set(ctx, mobile, capt, 2*time.Minute)
	//记得注释或删除掉！
	fmt.Println("验证码测试：" + internal.RedisClient.Get(ctx, mobile).String())

	b64, err := GetBase64(fileName)
	if err != nil {
		log.Logger.Error("GenCaptcha() encode失败")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		//发送给前端
		"captcha": b64,
	})
}

func GetBase64(fileName string) (string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	b := make([]byte, 10240)
	// StdEncoding 初始化base64，Encode 将参数2内容加密给参数1
	base64.StdEncoding.Encode(b, file)
	s := string(b)
	return s, nil

}
