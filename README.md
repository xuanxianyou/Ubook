# UBook
&nbsp;&nbsp;UBook是使用go-kit实现的微服务应用，正如其名所示，仅仅实现了  
用户模块和图书模块，算是个入门项目吧。
## 项目目录介绍
* authorization：授权服务器
* book：图书服务，主要实现了图书推荐功能
* gateway：API网关，通过反向代理实现
* kubernetes：基础服务文件，实现服务部署
* user：用户服务，主要包括用户注册、登录
