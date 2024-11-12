package handler

import "testing"

func TestGenCaptcha(t *testing.T) {
	err := GenCaptcha()
	if err != nil {
		panic(err)
	}
}
