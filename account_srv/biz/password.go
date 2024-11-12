package biz

import (
	"crypto/md5"
	"github.com/anaskhan96/go-password-encoder"
)

var option = password.Options{
	SaltLen:      16,
	Iterations:   100,
	KeyLen:       32,
	HashFunction: md5.New,
}

func GetMd5(pwd string) (string, string) {
	salt, p := password.Encode(pwd, &option)
	return salt, p
}

func Verify(pwd, salt, encodePwd string) bool {
	return password.Verify(pwd, salt, encodePwd, &option)
}
