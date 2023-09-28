package rpcserver

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthTokeninterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// 获取metadata 进行请求合法拦截, 不合法直接return
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("获取metadata失败")
	}
	token := md.Get("Authorization")
	fmt.Println("data data", token)

	logx.Info("拦截器前...")
	resp, err := handler(ctx, req)
	logx.Info("拦截器后...")
	return resp, err
}

func MetadataAppendToOutgoingContextInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 获取 HTTP 请求的 Header 对象
	r, ok := metadata.FromIncomingContext(ctx)
	fmt.Println("rrrrrrrrrrrrrrrrrrrrrrrrr", r, ok)
	var md metadata.MD
	if ok {
		// 创建 gRPC 请求的 Metadata 对象
		md = metadata.Pairs()

		// 将 Header 值添加到 Metadata
		for key, values := range r {
			for _, value := range values {
				md.Append(key, value)
			}
		}

		// 将新的 Metadata 添加到 gRPC 请求的上下文中
		ctx = metadata.NewOutgoingContext(ctx, r)
	}

	fmt.Println("mdmdmdmdmdmddmdm", md)
	logx.Info("调用rpc服务前")

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	logx.Info("调用rpc服务后")
	return nil
}
