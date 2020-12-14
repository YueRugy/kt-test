package service

import (
	"errors"
)

type IUserService interface {
	GetName(uid int) string
	DelById(uid int) error
}

type UserService struct {
}

func (us *UserService) GetName(uid int) string {
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
