package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/wl_frontendsuper/controller"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/dist", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("api/healthz", controller.HealthCheck)
	app.Post("/api/init", controller.Init)
	app.Post("/api/log", controller.Log)
	app.Post("/api/login", controller.CheckLogin)
	app.Post("/api/agen", controller.Agen)
	app.Post("/api/saveagen", controller.Agensave)
	app.Post("/api/agenlistadmin", controller.Agenlistadmin)
	app.Post("/api/saveagenlistadmin", controller.Agensavelistadmin)

	app.Post("/api/setting", controller.Setting)
	app.Post("/api/savesetting", controller.Settingsave)

	app.Post("/api/admin", controller.Admin)
	app.Post("/api/saveadmin", controller.Saveadmin)
	app.Post("/api/adminrule", controller.Adminrule)
	app.Post("/api/saveadminrule", controller.Saveadminrule)
	app.Post("/api/currency", controller.Currency)
	app.Post("/api/savecurrency", controller.Savecurr)
	return app
}
