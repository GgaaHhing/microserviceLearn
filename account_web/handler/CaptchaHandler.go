package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"os"
	"testProject/microservice/log"
)

func GenCaptcha() error {
	fileName := "captcha.png"
	f, err := os.Create(fileName)
	if err != nil {
		log.Logger.Error("GenCaptcha() create失败")
		return err
	}
	defer f.Close()
	digits := captcha.RandomDigits(captcha.DefaultLen)
	image := captcha.NewImage("", digits, captcha.StdWidth, captcha.StdHeight)
	// image/io的WriteTo需要传入一个Writer接口，所以只要实现了这个接口都可以传入
	_, err = image.WriteTo(f)
	if err != nil {
		log.Logger.Error("GenCaptcha() write失败")
		return err
	}
	b64, err := GetBase64(fileName)
	if err != nil {
		log.Logger.Error("GenCaptcha() encode失败")
		return err
	}
	fmt.Println(b64)
	return nil
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
