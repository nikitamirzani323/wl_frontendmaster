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
	app.Post("/api/companydetail", controller.Companydetail)
	app.Post("/api/companylistadmin", controller.Companylistadmin)
	app.Post("/api/companylistpasaran", controller.Companylistpasaran)
	app.Post("/api/companylistkeluaran", controller.Companylistkeluaran)
	app.Post("/api/companyinvoicemember", controller.Companyinvoicemember)
	app.Post("/api/companyinvoicemembertemp", controller.Companyinvoicemembertemp)
	app.Post("/api/companyinvoicemembersync", controller.Companyinvoicemembersync)
	app.Post("/api/companyinvoicegrouppermainan", controller.Companyinvoicegrouppermainan)
	app.Post("/api/companyinvoicelistpermainan", controller.Companyinvoicelistpermainan)
	app.Post("/api/companyinvoicelistpermainanstatus", controller.Companyinvoicelistpermainanstatus)
	app.Post("/api/companyinvoicelistpermainanmember", controller.Companyinvoicelistpermainanmember)
	app.Post("/api/companypasaranconf", controller.Companypasaranconf)
	app.Post("/api/companypasaranonline", controller.Companypasaranonline)
	app.Post("/api/savecompany", controller.Companysave)
	app.Post("/api/savecompanyadmin", controller.Companysaveadmin)
	app.Post("/api/savecompanyupdatepasaran", controller.Companysaveupdatepasaran)
	app.Post("/api/savecompanyupdatepasaranline", controller.Companysaveupdatepasaranline)
	app.Post("/api/savecompanypasaranonline", controller.Companysavepasaranonline)
	app.Post("/api/deletecompanypasaranonline", controller.Companydeletepasaranonline)
	app.Post("/api/savecompanypasaran", controller.Companysavecompanypasaran)
	app.Post("/api/updatecompanypasaran432", controller.Companyupdatepasaran432)
	app.Post("/api/updatecompanypasarancolokbebas", controller.Companyupdatepasarancolokbebas)
	app.Post("/api/updatecompanypasarancolokmacau", controller.Companyupdatepasarancolokmacau)
	app.Post("/api/updatecompanypasarancoloknaga", controller.Companyupdatepasarancoloknaga)
	app.Post("/api/updatecompanypasarancolokjitu", controller.Companyupdatepasarancolokjitu)
	app.Post("/api/updatecompanypasaran5050umum", controller.Companyupdatepasaran5050umum)
	app.Post("/api/updatecompanypasaran5050special", controller.Companyupdatepasaran5050special)
	app.Post("/api/updatecompanypasaran5050kombinasi", controller.Companyupdatepasaran5050kombinasi)
	app.Post("/api/updatecompanypasarankombinasi", controller.Companyupdatepasarankombinasi)
	app.Post("/api/updatecompanypasarandasar", controller.Companyupdatepasarandasar)
	app.Post("/api/updatecompanypasaranshio", controller.Companyupdatepasaranshio)
	app.Post("/api/fetchpasaranlimitline", controller.Companyfetchpasaranlimitline)
	app.Post("/api/fetchpasaran432", controller.Companyfetchpasaran432)
	app.Post("/api/fetchpasarancbebas", controller.Companyfetchpasarancbebas)
	app.Post("/api/fetchpasarancmacau", controller.Companyfetchpasarancmacau)
	app.Post("/api/fetchpasarancnaga", controller.Companyfetchpasarancnaga)
	app.Post("/api/fetchpasarancjitu", controller.Companyfetchpasarancjitu)
	app.Post("/api/fetchpasaran5050umum", controller.Companyfetchpasaran5050umum)
	app.Post("/api/fetchpasaran5050special", controller.Companyfetchpasaran5050special)
	app.Post("/api/fetchpasaran5050kombinasi", controller.Companyfetchpasaran5050kombinasi)
	app.Post("/api/fetchpasaranmacau", controller.Companyfetchpasaranmacau)
	app.Post("/api/fetchpasarandasar", controller.Companyfetchpasarandasar)
	app.Post("/api/fetchpasaranshio", controller.Companyfetchpasaranshio)

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
	return app
}
