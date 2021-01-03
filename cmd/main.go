package main

import (
	"encoding/base32"
	"fmt"
	"github.com/theykk/2fa-go"
	"io/ioutil"
)

func main() {
	// ? Read the secret token from file system
	data, err := ioutil.ReadFile("secret.pem")
	go2fa.Check(err)

	// ? Generate base32 string from secret
	key := base32.StdEncoding.EncodeToString(data)
	otp := go2fa.GetTOTPToken(key)

	// ? Print otp url and otp code
	fmt.Println("otpauth://totp/TheYkkGO:ykk@theykk.net?secret=" + key + "&issuer=TheYkkGO")
	fmt.Println(otp)
}
