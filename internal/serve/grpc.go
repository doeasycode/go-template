/**
 * @author   Liang
 * @create   2022/8/19 11:03
 * @version  1.0
 */

package serve

import (
	"context"
	"errors"
	"go-template/api"
	"go-template/internal/service"
	"net"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

func StartGrpc(addr string) error {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	// 注册interceptor
	opts = append(opts, grpc.UnaryInterceptor(auth))

	// 实例化grpc Server
	s := grpc.NewServer(opts...)

	// 注册Service
	api.RegisterApiServiceServer(s, new(service.ApiService))
	api.RegisterAdminServiceServer(s, new(service.AdminService))

	defer func() {
		s.Stop()
		listen.Close()
	}()

	return s.Serve(listen)

}

func auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("无Token认证信息")
	}

	token := md.Get("token")

	if token == nil || len(token) == 0 {
		return nil, errors.New("无Token认证信息")
	}

	ctx = context.WithValue(ctx, "token", token[0])

	// 继续处理请求
	return handler(ctx, req)
}
