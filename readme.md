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

```api 生成
1. 在 app/ 下面微服务目录 例如 verify 目录执行，生成api目录结构
goctl api go -api ./api/*.api -dir ./api
```

```rpc 生成
1. 在 app/ 下面微服务目录 例如 verify/grpc 目录执行，生成api目录结构
goctl rpc protoc verify.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

``` gateway 添加 rpc 
1. 在 official-verification 目录下
protoc --descriptor_set_out=gateway/pb/verify.pb app/verify/grpc/verify.proto
在 gateway/etc/gateway.yaml 添加配置
  - Grpc:
      Target: 0.0.0.0:8080
    # protoset mode
    ProtoSets:
      - ./pb/verify.pb
    # Mappings can also be written in proto options
    Mappings:
      - Method: post
        Path: /rpc/offical/verify/listinform
        RpcPath: verify.Verify/ListInform 
```

### model生成

```api
goctl model mysql ddl -src ./model/official_verify.sql -dir ./model
goctl model mysql ddl -src ./model/verify_enum.sql -dir ./model
goctl model mysql ddl -src ./model/social_enum.sql -dir ./model
goctl model mysql ddl -src ./model/report_record.sql -dir ./model
```

### 文档生成

```
goctl api plugin -plugin goctl-swagger="swagger -filename verify.json -basepath /api" -api ./app/verify/api/verify.api -dir ./doc
```


### 文档运行

```api
3.运行docker 查看swgger文档,在解决方案exchange-asset-services下面执行
docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/doc.json -v $PWD/doc/:/foo swaggerapi/swagger-ui 
```
