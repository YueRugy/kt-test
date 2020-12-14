package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/time/rate"
	"kt-test/util"
	"net/http"
	"strconv"
)

type UserRequest struct {
	Uid    int `json:"uid"`
	Method string
}

type UserResponse struct {
	Name string `json:"name"`
}

func UserLogger(l log.Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			uid := request.(UserRequest).Uid
			_ = l.Log("method", "GET", "uid", strconv.Itoa(uid))
			return e(ctx, request)
		}
	}
}

func RateLimit(l *rate.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !l.Allow() {
				//response, _ := http.ResponseWriter.Write([]byte("too mayny requests" + strconv.Itoa(util.ServicePort)))
				return nil, util.NewError(http.StatusTooManyRequests, "too many requests")
			}
			return e(ctx, request)
		}
	}
}

func GetEndpoint(us *UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		uid := r.Uid
		result := "nothing"
		switch r.Method {
		case "GET":
			result = us.GetName(uid) + strconv.Itoa(util.ServicePort)
		case "DELETE":
			err1 := us.DelById(uid)
			if err1 != nil {
				result = err1.Error()
			} else {
				result = fmt.Sprintf("userid位%d的用户删除成功", uid)
			}
		}
		return UserResponse{
			Name: result,
		}, nil
	}
}
func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-Type", contentType)
	if e, ok := err.(*util.MyError); ok {
		w.WriteHeader(e.Code)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, _ = w.Write(body)
}
