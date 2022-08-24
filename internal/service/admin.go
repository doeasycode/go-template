/**
 * @author   Liang
 * @create   2022/8/19 11:16
 * @version  1.0
 */

package service

import (
	"context"
	"fmt"
	"go-template/api"
)

var (
	_ api.AdminServiceServer = (*AdminService)(nil)
)

type AdminService struct {
}

func (s *AdminService) Add(ctx context.Context, request *api.AddRequest) (*api.AddResponse, error) {

	fmt.Println("user_id:", ctx.Value("user_id"))

	return &api.AddResponse{
		C: request.A + request.B,
	}, nil
}
