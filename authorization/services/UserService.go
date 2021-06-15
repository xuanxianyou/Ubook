package services

import (
	"GoProject/WebProject/MicroService/UserBook/authorization/model"
	"context"
	"errors"
)

var (
	ErrUserNotExist = errors.New("username is not exist")
	ErrPassword = errors.New("invalid password")
)
// Service Define a service interface
type UserService interface {
	// Get User By username
	GetUserDetailByUsername(ctx context.Context, username, password string) (model.User, error)
}

//UserService implement Service interface
type InMemoryUserService struct {
	userDetailsDict map[string]*model.User

}

func (service *InMemoryUserService) GetUserDetailByUsername(ctx context.Context, username, password string) (model.User, error) {


	// 根据 username 获取用户信息
	userDetails, ok := service.userDetailsDict[username]; if ok{
		// 比较 password 是否匹配
		if userDetails.Password == password{
			return *userDetails, nil
		}else {
			return model.User{}, ErrPassword
		}
	}else {
		return model.User{}, ErrUserNotExist
	}


}

func NewInMemoryUserService(userDetailsList []*model.User) *InMemoryUserService {
	userDetailsDict := make(map[string]*model.User)

	if userDetailsList != nil {
		for _, value := range userDetailsList {
			userDetailsDict[value.Username] = value
		}
	}

	return &InMemoryUserService{
		userDetailsDict:userDetailsDict,
	}
}

