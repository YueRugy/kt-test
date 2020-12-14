package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	m:=mux.Vars(r)
	log.Println(m)
	uid, err := strconv.Atoi(mux.Vars(r)["uid"])
	if err != nil {
		return nil, errors.New("参数错误")
	}
	return UserRequest{
		Uid: uid,
		Method: r.Method,
	}, nil

}

func EncodeUserResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
