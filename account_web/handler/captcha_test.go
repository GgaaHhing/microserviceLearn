package handler

import "testing"

func TestGenCaptcha(t *testing.T) {
	err := CaptchaHandler()
	if err != nil {
		panic(err)
	}
}
