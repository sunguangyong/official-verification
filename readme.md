### telegram bot 

### 非常重要

```api
goctl 暂时没有实现map[string]interface{} 在.api 不要使用map参数类型
```

### 环境安装

```
  golang 1.17.8版本
  安装包下载地址
  http://mirrors.nju.edu.cn/golang/
```

```
  goctl 版本1.3.5
  安装命令  GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@v1.3.5
```

```
  goctl-go-compact
  安装命令  GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-go-compact@latest
``` 

```
  goctl-swagger
  git clone https://github.com/zeromicro/goctl-swagger.git
  修改 go.mod 与 goctl 版本一致
  github.com/zeromicro/go-zero/tools/goctl v1.3.5
  go build -o goctl-swagger main.go
  mv goctl-swagger $GOPATH/bin
```

```
  项目tpl配置
  goctl template clean 清除
  cd $HOME/.goctl
  goctl template init 初始化``
```

### 项目构建

```api
1. 在 app/ 下面微服务目录 例如 coin目录执行，生成api目录结构
goctl api go -api ./api/*.api -dir ./api
```

### model生成

```api
goctl model mysql ddl -src ./model/official_verify.sql -dir ./model
goctl model mysql ddl -src ./model/verify_enum.sql -dir ./model
goctl model mysql ddl -src ./model/social_enum.sql -dir ./model
```

### 文档生成

```
goctl api plugin -plugin goctl-swagger="swagger -filename verify.json -basepath /api" -api ./app/verify/api/verify.api -dir ./doc

```

### 文档迁移

``` cp ./doc/verify.json ../cointiger-golang-doc/doc 
    提交 更改
```

### 文档部署

``` 
    登陆 jenkins
    http://172.16.10.20:8080/jenkins/job/dev/job/golang-doc/
    账号 域控账号    预控密码
    查找 develop golang-doc
    Build Now
    查看 控制台输出是否完成    
    文档地址  http://172.16.10.10:10000/#/Operate 
    输入 ./operate.json 查看文档
```

### 文档运行

```api
3.运行docker 查看swgger文档,在解决方案exchange-asset-services下面执行
docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/doc.json -v $PWD/doc/:/foo swaggerapi/swagger-ui 
```

### jenkins pipline notes:

```api
再往k8s或者eks中部署的时候pipline中需要移动项目下 Dockerfile 到根目录下,各个项目生成不同的Dockerfile放在不同的项目例如api/rpc下面。
本地测试生成image命令**在根目录下执行**：docker build -t statistics-api:v2 .
```
