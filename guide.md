```
├── Dockerfile
├── README.md
├── cmd 启动脚本
│   ├── db 更新表
│   └── server
├── docs 文档（不能手动修改）
├── env.toml 环境配置文件
├── go.mod
├── guide.md 使用指南
├── internal
│   ├── app 应用入口
│   ├── controller 接口入口层
│   ├── infra 基础设施
│   ├── model 业务模型
│   │   ├── req 请求参数
│   │   ├── res 返回参数
│   │   ├── schema 数据库表
│   │   ├── xerror 错误码
│   │   └── env 环境变量对应 env.toml 里的字段
│   ├── pkg 三方包
│   ├── repo 数据库操作层
│   ├── server http服务等入口
│   ├── service 业务逻辑层
│   └── util 工具包
└── main.go 应用启动入口
```

三方包

- uber-fx：依赖注入框架，用于解耦应用初始化
- validator：参数校验框架，验证请求参数及返回值标注
- viper：配置文件读取
- swagger：根据注释生成文档
- zap：日志框架

## 项目目标
信息交流

## 分解任务
1.第一期
1）共享图片：发帖，消息，个人
2）创建5张数据表：user，post，postlike，postfile，postreply
2）完成11个接口：（1）登录，（2）注册，（3）createpost，（4）post detail，（5）postdelete，（6）home，（7）like，（8）replay，（9）replaydelete，（10）mymessage（11）fileput，
2.第二期
1）共享视频
3.第三期
1）远程语音匹配

## 时间表
1.3月9日之前：1）学习DB
2）模版代码理解
2.3月16日之前：完成接口（1，2，11，3，4）
3.3月23日之前：完成接口（5.7.8.9.6）
4.3月31日之前：完成接口（10）

## 开发流程
1. 根据业务划分在 controller 里创建对应的文件，如 user nostr。这里的文件名应该和接口中的Tags一致，此层只处理入参校验和发送结果
2. 在 servic 层创建和controller里对应的文件，此层处理业务逻辑，返回封装后的报错信息。错误码通过枚举定义，方便日后做监测和错误归因分析
3. 在 repo 层创建和service里对应的文件，此层处理数据库操作，返回原始错误信息
4. 在 model 层里的对应出入参文件夹创建和controller里对应的文件，所有结构体都应该有注释和校验规则，方便swagger生成文档和参数校验


## 注意事项
1. env.toml 里的配置项和 model.env 里的结构体对应，修改时需要同步
2. 0module 是用来做初始化配置，为了在目录里排第一所以开头是0
3. controller 里的方法都需要有注释，方便 swagger 生成文档，对应返回错误码如果需要前端有业务逻辑时也应该标注，所有的query和body都应该通过结构体定义，方便参数校验
4. 所有的三方服务封装逻辑都应该在 pkg 里，方便以后替换
5. server.http.middlewares 里的中间件都是全局的，如果需要局部的应该在对应的 controller 里定义





