# hiyou

### Project Description
hiyou CMS系统服务器端

### Project Initialization
- 克隆代码
    ```shell
    git clone https://github.com/Sabrinaapple/APP_backend
    cd APP_backend
    ```

- 创建表 只有新增或者修改表结构时执行,运行项目时不需要执行此命令
    ```shell
    go ./cmd/db/main.go
    ```
- 生成文档 运行项目前执行此命令
    ```shell
    swag init -g ./internal/server/http/create.go  -o ./docs   
    ```
- 启动
    ```bash
    go mod tidy
    go run .
    ```
应用启动后文档地址为：http://localhost:9000/swagger/index.html

### Test process
- Start this project,
    ```bash
    go run .
    ```
- After see
    ```bash
    Server is running on :<SERVER_LISTEN_PORT>...
    ```
- Open a new terminal, in root path, run the command line below
    ```bash
    go clean -testcache
    go test -v ./test/
    ```
