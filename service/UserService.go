package service

import (
	"errors"
	"log"
)

type IUserService interface {
	GetName(uid int) string
	DelById(uid int) error
}

type UserService struct {
}

func (us *UserService) GetName(uid int) string {
	log.Println(uid)
	if uid == 101 {
		return "yue"
	}
	return "guest"
}
func (us *UserService) DelById(uid int) error {
	if uid == 101 {
		return errors.New("无权限")
	}
	return nil
}
