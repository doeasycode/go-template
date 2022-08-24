/**
 * @author   Liang
 * @create   2022/8/19 19:30
 * @version  1.0
 */

package serve

import (
	"context"
	"go-template/api"
	"go-template/internal/service"

	"github.com/gofiber/fiber/v2"
)

func StartHttp(addr string) error {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Context().SetUserValue("abc", "这是abc")
		ctx := context.WithValue(c.UserContext(), "user_id", "123456")
		c.SetUserContext(ctx)
		return c.Next()
	})
	api.RegisterAdminServiceHttpServer(app, new(service.AdminService))
	return app.Listen(addr)
}
