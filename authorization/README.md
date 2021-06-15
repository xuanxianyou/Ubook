# Authorization: (授权服务器)
授权服务器主要实现颁发令牌和验证令牌功能，对此，授权服务器提供了两个接口：
* /oauth/token: 用于客户端携带用户凭证获取访问令牌
* /oauth/check_token: 验证访问令牌的有效性，访问令牌绑定的的客户端和用户信息 
## 授权服务器的组成
* TokenGranter: 令牌生成器，根据客户端授权类型，验证客户端和用户信息
* TokenService: 生成和管理令牌
* TokenStore: 实现令牌的存取
* ClientService: 根据ClientId查询系统注册的客户端信息
* UserService: 用户获取用户信息
