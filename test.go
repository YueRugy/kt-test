package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func main() {
	//_ = util.GenRsaKey(1024)
	//Hs256()
}

func Hs256() {
	sec := []byte("123abc")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{UserName: "yue"})
	ts, _ := t.SignedString(sec)
	fmt.Println(ts)
	uc := UserClaim{}
	token, _ := jwt.ParseWithClaims(ts, &uc, func(token *jwt.Token) (interface{}, error) {
		return sec, nil
	})
	if token.Valid {
		//fmt.Println(token.Claims)
		fmt.Println(token.Claims.(*UserClaim).UserName)
		fmt.Println(token.Claims.(*UserClaim).ExpiresAt)
	}
}
func Rs256()  {
	
}
