/**
 * @author   Liang
 * @create   2022/8/19 10:50
 * @version  1.0
 */

package service

import (
	"context"
	"go-template/api"
)

var (
	_ api.ApiServiceServer = (*ApiService)(nil)
)

type ApiService struct {
}

func (s *ApiService) Echo(ctx context.Context, request *api.EchoRequest) (*api.EchoResponse, error) {
	return &api.EchoResponse{
		Out: request.In,
	}, nil
}
