## Auth-service 用户信息和鉴权
### 软件包
* config 
  * app.yml 存放web服务和数据库等配置文件
  * config.go 存放数据库配置接口等代码
* dao 
  * 数据库操作函数
* router
  * 存放各个服务的路由
  * auth/route.go 用户路由
  * routers.go 公共路由文件
* service
  * 存放grpc服务的接口实现服务
* target
  * 配置Goland 运行配置，存放打包的exe可执行文件
* util
  * 存放相关工具类
* main.go 程序入口
* go.mod 依赖包