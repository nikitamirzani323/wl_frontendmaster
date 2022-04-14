package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/go_mastertoto/controller"
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
	app.Post("/api/login", controller.CheckLogin)
	app.Post("/api/company", controller.Company)
	app.Post("/api/savecompany", controller.Companysave)

	app.Post("/api/allpasaran", controller.Pasaran)
	app.Post("/api/pasarandetail", controller.Pasarandetail)
	app.Post("/api/pasarandetailconf", controller.Pasarandetailconf)
	app.Post("/api/savepasaran", controller.Savepasaran)
	app.Post("/api/savepasaranlimitline", controller.Savepasaranlimitline)
	app.Post("/api/savepasaran432", controller.Savepasaran432)
	app.Post("/api/savepasarancbebas", controller.Savepasarancbebas)
	app.Post("/api/savepasarancmacau", controller.Savepasarancmacau)
	app.Post("/api/savepasarancnaga", controller.Savepasarancnaga)
	app.Post("/api/savepasarancjitu", controller.Savepasarancjitu)
	app.Post("/api/savepasaran5050umum", controller.Savepasaran5050umum)
	app.Post("/api/savepasaran5050special", controller.Savepasaran5050special)
	app.Post("/api/savepasaran5050kombinasi", controller.Savepasaran5050kombinasi)
	app.Post("/api/savepasaranmacau", controller.Savepasaranmacau)
	app.Post("/api/savepasarandasar", controller.Savepasarandasar)
	app.Post("/api/savepasaranshio", controller.Savepasaranshio)

	app.Post("/api/invoice", controller.Invoice)
	app.Post("/api/saveinvoice", controller.Saveinvoice)
	app.Post("/api/saveinvoicewinlose", controller.Saveinvoicewinlose)

	app.Post("/api/setting", controller.Setting)
	app.Post("/api/savesetting", controller.Settingsave)
	app.Post("/api/domain", controller.Domain)
	app.Post("/api/savedomain", controller.Savedomain)

	app.Post("/api/admin", controller.Admin)
	app.Post("/api/saveadmin", controller.Saveadmin)
	app.Post("/api/adminrule", controller.Adminrule)
	app.Post("/api/saveadminrule", controller.Saveadminrule)
	app.Post("/api/currency", controller.Currency)
	app.Post("/api/savecurrency", controller.Savecurr)
	return app
}
