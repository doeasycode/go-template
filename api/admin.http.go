/**
 * @author   Liang
 * @create   2022/8/19 19:24
 * @version  1.0
 */

package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	srv AdminServiceServer
)

func RegisterAdminServiceHttpServer(fiber *fiber.App, service AdminServiceServer) {
	srv = service
	fiber.Post("/admin/add", addHandler)
}

func addHandler(ctx *fiber.Ctx) error {
	//ctx := context.WithValue(ctx.Context() , "" , "")

	//get := ctx.Get("token")

	fmt.Println("x-token:", ctx.Get("x-token"))

	fmt.Println("abc:", ctx.Context().UserValue("abc"))

	req := &AddRequest{}

	err := ctx.BodyParser(req)

	if err != nil {
		return ctx.SendString(err.Error())
	}

	rs, err := srv.Add(ctx.UserContext(), req)

	if err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.JSON(rs)
}
