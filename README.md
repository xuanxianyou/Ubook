# micro-service:user
&nbsp;&nbsp;本部分实现了用户模块，主要功能包括用户注册、用户登录。
## 用户注册
&nbsp;&nbsp用户注册应当提交Username(用户名)、Phone(联系电话)、Password(密码) 
## 用户登录
&nbsp;&nbsp;采用策略模式实现用户登录逻辑  
1. 密码登录
用户提交Username(用户名)/Phone(联系电话)和密码(password)
2. 短信登录
用户提交Phone获取Code(验证码)，提交Phone和Code实现登录 
