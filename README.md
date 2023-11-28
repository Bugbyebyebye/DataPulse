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
* log-servie
  * 实现日志记录
  * 流量统计
  * 用户操作记录
* task-service
  * 用户设置任务
  * 用户设置接口
  * 实现数据清洗服务等
### 技术栈
* Gin 搭建web服务和网关
* gRpc 服务间通信
* GORM 数据库操作框架
* Viper 开源项目 获取yaml配置
* ETCD 注册发现grpc服务
* Redis 缓存验证码