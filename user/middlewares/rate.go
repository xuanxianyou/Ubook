package middlewares

import (
	"GoProject/WebProject/MicroService/UserBook/user/model"
	"context"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

// 限流中间件
func RateLimit(limit *rate.Limiter) endpoint.Middleware{
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow(){
				return nil,model.NewServerError(429,"request frequently")
			}
			return next(ctx,request)
		}
	}
}

