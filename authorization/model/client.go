package model

type Client struct {
	ClientId                    string   // client 的标识
	ClientSecret                string   // client 的密钥
	AccessTokenValiditySeconds  int      // 访问令牌有效时间，秒
	RefreshTokenValiditySeconds int      // 刷新令牌有效时间，秒
	RegisteredRedirectUri       string   // 重定向地址，授权码类型中使用
	AuthorizedGrantTypes        []string // 授权类型
}
