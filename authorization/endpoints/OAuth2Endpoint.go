package endpoint

import (
	. "GoProject/WebProject/MicroService/UserBook/authorization/model"
	"GoProject/WebProject/MicroService/UserBook/authorization/services"
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"net/http"
)

//
type OAuth2Endpoints struct {
	TokenEndpoint		endpoint.Endpoint
	CheckTokenEndpoint	endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}



func MakeClientAuthorizationMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			if err, ok := ctx.Value(OAuth2ErrorKey).(error); ok{
				return nil, err
			}
			if _, ok := ctx.Value(OAuth2ClientDetailsKey).(Client); !ok{
				return  nil, ErrInvalidClientRequest
			}
			return next(ctx, request)
		}
	}
}

func MakeOAuth2AuthorizationMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			if err, ok := ctx.Value(OAuth2ErrorKey).(error); ok{
				return nil, err
			}
			if _, ok := ctx.Value(OAuth2DetailsKey).(*OAuth2Details); !ok{
				return  nil, ErrInvalidUserRequest
			}
			return next(ctx, request)
		}
	}
}
func MakeAuthorityAuthorizationMiddleware(authority string, logger log.Logger) endpoint.Middleware  {
	return func(next endpoint.Endpoint) endpoint.Endpoint {

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			if err, ok := ctx.Value(OAuth2ErrorKey).(error); ok{
				return nil, err
			}
			if details, ok := ctx.Value(OAuth2DetailsKey).(*OAuth2Details); !ok{
				return  nil, ErrInvalidClientRequest
			}else {
				for _, value := range details.User.Permissions{
					if value == authority{
						return next(ctx, request)
					}
				}
				return nil, ErrNotPermit
			}
		}
	}
}

const (

	OAuth2DetailsKey       = "OAuth2Details"
	OAuth2ClientDetailsKey = "OAuth2ClientDetails"
	OAuth2ErrorKey         = "OAuth2Error"

)


var (
	ErrInvalidClientRequest = errors.New("invalid client message")
	ErrInvalidUserRequest = errors.New("invalid user message")
	ErrNotPermit = errors.New("not permit")
)



type TokenRequest struct {
	GrantType string
	Reader *http.Request
}


type TokenResponse struct {
	AccessToken *Token `json:"access_token"`
	Error string             `json:"error"`
}

//  make endpoint
func MakeTokenEndpoint(svc services.TokenGranter, clientService services.ClientDetailsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// 获取请求
		req := request.(*TokenRequest)
		//
		token, err := svc.Grant(ctx, req.GrantType, ctx.Value(OAuth2ClientDetailsKey).(Client), req.Reader)
		var errString = ""
		if err != nil{
			errString = err.Error()
		}

		return TokenResponse{
			AccessToken:token,
			Error:errString,
		}, nil
	}
}


type CheckTokenRequest struct {
	Token         string
	ClientDetails Client
}

type CheckTokenResponse struct {
	OAuthDetails *OAuth2Details `json:"o_auth_details"`
	Error string                `json:"error"`

}

func MakeCheckTokenEndpoint(svc services.TokenService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CheckTokenRequest)
		tokenDetails, err := svc.GetOAuth2DetailsByAccessToken(req.Token)

		var errString = ""
		if err != nil{
			errString = err.Error()
		}

		return CheckTokenResponse{
			OAuthDetails:tokenDetails,
			Error:errString,
		}, nil
	}
}


// HealthRequest 健康检查请求结构
type HealthRequest struct{}

// HealthResponse 健康检查响应结构
type HealthResponse struct {
	Status bool `json:"status"`
}

// MakeHealthCheckEndpoint 创建健康检查Endpoint
func MakeHealthCheckEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		return HealthResponse{
			Status:status,
		}, nil
	}
}
