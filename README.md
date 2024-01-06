# DataPulse 数据中台
### 文件结构&业务功能
* common 存放公共配置
* gateway-service 
  * 由Gin框架组成
  * 负责校验鉴权
  * 转发网络请求到各个服务。
* auth-service
  * Gin + grpc
  * 处理用户注册 -新增用户
  * 用户登录 - 签发JWTtoken
* store-service
  * Gin + gprc
  * 实现所有的数据库操作
    * 新增用户
    * 新增日志
    * 新增统计
    * 创建数据仓库
      * 创建数据表
      * 创建数据表
    * 创建Api服务接口
  * 关联底层的元数据库
* store-service-grpc
    * 存放底层三个grpc服务
* log-servie
  * 实现日志记录
  * 流量统计
  * 用户操作记录
* task-service
  * 用户设置任务
  * 用户设置接口
  * 实现数据清洗服务等
  * API池实现服务[仓库](https://github.com/JDruki/DataPulse-DockerAPI)
### 技术栈
* Gin 搭建web服务和网关
* gRpc 服务间通信
* GORM 数据库操作框架
* Viper 开源项目 获取yaml配置
* ETCD 注册发现grpc服务
* Redis 缓存验证码

### 服务分布

* gateway 网关 :8080
* auth-service :8081
    * gin 鉴权服务 :9000
* log-service :8082
* store-service :8083
* store-service-grpc 数据读取存储服务
    * mysql1 gin :8085
    * mysql2 gin :8086
    * mongo gin :8087
* task-service :8084
