package services

import (
	"GoProject/WebProject/MicroService/UserBook/authorization/model"
	"context"
	"errors"
)

var (

	ErrClientNotExist = errors.New("clientId is not exist")
	ErrClientSecret = errors.New("invalid clientSecret")

)

// Service Define a service interface
type ClientDetailsService interface {
	GetClientDetailsByClientId(ctx context.Context, clientId string, clientSecret string)( model.Client, error)
}

type InMemoryClientDetailsService struct {
	clientDetailsDict map[string]*model.Client

}

func NewInMemoryClientDetailService(clientDetailsList []*model.Client ) *InMemoryClientDetailsService{
	clientDetailsDict := make(map[string]*model.Client)

	if len(clientDetailsList) > 0  {
		for _, value := range clientDetailsList {
			clientDetailsDict[value.ClientId] = value
		}
	}

	return &InMemoryClientDetailsService{
		clientDetailsDict:clientDetailsDict,
	}
}


func (service *InMemoryClientDetailsService)GetClientDetailsByClientId(ctx context.Context, clientId string, clientSecret string)(model.Client, error) {
	// 根据 clientId 获取 clientDetails
	clientDetails, ok := service.clientDetailsDict[clientId]; if ok{
		// 比较 clientSecret 是否正确
		if clientDetails.ClientSecret == clientSecret{
			return *clientDetails, nil
		}else {
			return model.Client{}, ErrClientSecret
		}
	}else {
		return model.Client{}, ErrClientNotExist
	}
}


