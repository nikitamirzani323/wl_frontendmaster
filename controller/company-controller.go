package controller

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func Company(c *fiber.Ctx) error {
	type company_all struct {
		Company_search string `json:"company_search"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(company_all)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"company_search":  client.Company_search,
		}).
		Post(PATH + "api/allcompany")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companydetail(c *fiber.Ctx) error {
	type company_detail struct {
		Company string `json:"company"`
		Master  string `json:"master" `
		Page    string `json:"page" `
		Sdata   string `json:"sdata" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(company_detail)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"company":         client.Company,
			"page":            client.Page,
			"sdata":           client.Sdata,
		}).
		Post(PATH + "api/detailcompany")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companylistadmin(c *fiber.Ctx) error {
	type payload_listadmin struct {
		Company string `json:"company"`
		Master  string `json:"master" `
		Page    string `json:"page" `
		Sdata   string `json:"sdata" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_listadmin)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"sData":           client.Sdata,
			"company":         client.Company,
		}).
		Post(PATH + "api/companylistadmin")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companylistpasaran(c *fiber.Ctx) error {
	type payload_listadmin struct {
		Company string `json:"company"`
		Master  string `json:"master" `
		Page    string `json:"page" `
		Sdata   string `json:"sdata" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_listadmin)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"sData":           client.Sdata,
			"company":         client.Company,
		}).
		Post(PATH + "api/companylistpasaran")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companylistkeluaran(c *fiber.Ctx) error {
	type payload_listkeluaran struct {
		Master  string `json:"master" `
		Page    string `json:"page" `
		Company string `json:"company"`
		Periode string `json:"periode"`
		Pasaran int    `json:"pasaran"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_listkeluaran)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault_listinvoice{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"company":         client.Company,
			"periode":         client.Periode,
			"pasaran":         client.Pasaran,
		}).
		Post(PATH + "api/companylistkeluaran")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*responsedefault_listinvoice)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"totalwinlose": result.Totalwinlose,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyinvoicemember(c *fiber.Ctx) error {
	type payload_invoicemember struct {
		Master   string `json:"master" `
		Page     string `json:"page" `
		Company  string `json:"company"`
		Username string `json:"username" `
		Invoice  int    `json:"invoice" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_invoicemember)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault_listinvoice{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"company":         client.Company,
			"username":        client.Username,
			"invoice":         client.Invoice,
		}).
		Post(PATH + "api/companyinvoicemember")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault_listinvoice)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"totalwinlose": result.Totalwinlose,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyinvoicemembertemp(c *fiber.Ctx) error {
	type payload_invoicemember struct {
		Master   string `json:"master" `
		Page     string `json:"page" `
		Company  string `json:"company"`
		Username string `json:"username" `
		Invoice  int    `json:"invoice" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_invoicemember)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault_listinvoice{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"company":         client.Company,
			"username":        client.Username,
			"invoice":         client.Invoice,
		}).
		Post(PATH + "api/companyinvoicemembertemp")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault_listinvoice)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"totalwinlose": result.Totalwinlose,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyinvoicemembersync(c *fiber.Ctx) error {
	type payload_invoicemembersync struct {
		Master   string `json:"master" `
		Page     string `json:"page" `
		Company  string `json:"company"`
		Username string `json:"username" `
		Invoice  int    `json:"invoice" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_invoicemembersync)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault_listinvoice{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"company":         client.Company,
			"username":        client.Username,
			"invoice":         client.Invoice,
		}).
		Post(PATH + "api/companyinvoicemembersync")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault_listinvoice)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"totalwinlose": result.Totalwinlose,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyinvoicegrouppermainan(c *fiber.Ctx) error {
	type payload_invoicegrouppermainan struct {
		Master   string `json:"master" `
		Page     string `json:"page" `
		Company  string `json:"company"`
		Username string `json:"username" `
		Invoice  int    `json:"invoice" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_invoicegrouppermainan)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault_listinvoice{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"company":         client.Company,
			"username":        client.Username,
			"invoice":         client.Invoice,
		}).
		Post(PATH + "api/companyinvoicegrouppermainan")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault_listinvoice)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"totalwinlose": result.Totalwinlose,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyinvoicelistpermainan(c *fiber.Ctx) error {
	type payload_invoicelistpermainan struct {
		Master    string `json:"master" `
		Page      string `json:"page" `
		Company   string `json:"company"`
		Invoice   int    `json:"invoice"`
		Permainan string `json:"permainan"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_invoicelistpermainan)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault_listinvoice{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"company":         client.Company,
			"invoice":         client.Invoice,
			"permainan":       client.Permainan,
		}).
		Post(PATH + "api/companyinvoicelistpermainan")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault_listinvoice)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"totalwinlose": result.Totalwinlose,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyinvoicelistpermainanstatus(c *fiber.Ctx) error {
	type payload_invoicelistpermainanstatus struct {
		Master  string `json:"master" `
		Page    string `json:"page" `
		Company string `json:"company"`
		Invoice int    `json:"invoice"`
		Status  string `json:"status"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_invoicelistpermainanstatus)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault_listinvoice{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"company":         client.Company,
			"invoice":         client.Invoice,
			"status":          client.Status,
		}).
		Post(PATH + "api/companyinvoicelistpermainanstatus")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault_listinvoice)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"totalwinlose": result.Totalwinlose,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyinvoicelistpermainanmember(c *fiber.Ctx) error {
	type payload_invoicelistpermainanmember struct {
		Master    string `json:"master" `
		Page      string `json:"page" `
		Company   string `json:"company"`
		Invoice   int    `json:"invoice"`
		Username  string `json:"username" `
		Permainan string `json:"permainan"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_invoicelistpermainanmember)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault_listinvoice{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"company":         client.Company,
			"invoice":         client.Invoice,
			"username":        client.Username,
			"permainan":       client.Permainan,
		}).
		Post(PATH + "api/companyinvoicelistpermainanmember")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault_listinvoice)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"totalwinlose": result.Totalwinlose,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companypasaranconf(c *fiber.Ctx) error {
	type payload_pasaranconf struct {
		Page              string `json:"page" `
		Sdata             string `json:"sData" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaranconf)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"page":              client.Page,
			"sData":             client.Sdata,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
		}).
		Post(PATH + "api/companylistpasaranconf")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companypasaranonline(c *fiber.Ctx) error {
	type payload_pasaranconf struct {
		Page              string `json:"page" `
		Sdata             string `json:"sData" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaranconf)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"page":              client.Page,
			"sData":             client.Sdata,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
		}).
		Post(PATH + "api/companylistpasaranonline")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companysave(c *fiber.Ctx) error {
	type company_save struct {
		Sdata     string `json:"sdata"`
		Company   string `json:"company"`
		Master    string `json:"master"`
		Name      string `json:"name"`
		Urldomain string `json:"urldomain"`
		Status    string `json:"status"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(company_save)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"company":         client.Company,
			"master":          client.Master,
			"sdata":           client.Sdata,
			"name":            client.Name,
			"urldomain":       client.Urldomain,
			"status":          client.Status,
		}).
		Post(PATH + "api/savecompany")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companysaveadmin(c *fiber.Ctx) error {
	type payload_saveadmin struct {
		Sdata          string `json:"sdata"`
		Company        string `json:"company"`
		Master         string `json:"master"`
		Admin_username string `json:"admin_username"`
		Admin_password string `json:"admin_password" `
		Admin_name     string `json:"admin_name"`
		Admin_status   string `json:"admin_status"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_saveadmin)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"sdata":           client.Sdata,
			"master":          client.Master,
			"company":         client.Company,
			"admin_username":  client.Admin_name,
			"admin_password":  client.Admin_password,
			"admin_name":      client.Admin_name,
			"admin_status":    client.Admin_status,
		}).
		Post(PATH + "api/savecompanyadmin")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companysavecompanypasaran(c *fiber.Ctx) error {
	type payload_savecompanypasaran struct {
		Sdata      string `json:"sdata" `
		Company    string `json:"company" `
		Master     string `json:"master" `
		Pasaran_id string `json:"pasaran_id" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_savecompanypasaran)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"sdata":           client.Sdata,
			"master":          client.Master,
			"company":         client.Company,
			"pasaran_id":      client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanypasaran")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companysaveupdatepasaran(c *fiber.Ctx) error {
	type payload_updatepasaran struct {
		Sdata                string `json:"sdata" `
		Company              string `json:"company" `
		Master               string `json:"master" `
		Companypasaran_id    int    `json:"companypasaran_id" `
		Pasaran_diundi       string `json:"pasaran_diundi"`
		Pasaran_url          string `json:"pasaran_url" `
		Pasaran_jamtutup     string `json:"pasaran_jamtutup" `
		Pasaran_jamjadwal    string `json:"pasaran_jamjadwal" `
		Pasaran_jamopen      string `json:"pasaran_jamopen"`
		Pasaran_statusactive string `json:"pasaran_statusactive" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_updatepasaran)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":      hostname,
			"master":               client.Master,
			"company":              client.Company,
			"companypasaran_id":    client.Companypasaran_id,
			"pasaran_diundi":       client.Pasaran_diundi,
			"pasaran_url":          client.Pasaran_url,
			"pasaran_jamtutup":     client.Pasaran_jamtutup,
			"pasaran_jamjadwal":    client.Pasaran_jamjadwal,
			"pasaran_jamopen":      client.Pasaran_jamopen,
			"pasaran_statusactive": client.Pasaran_statusactive,
		}).
		Post(PATH + "api/savecompanyupdatepasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companysaveupdatepasaranline(c *fiber.Ctx) error {
	type payload_updatepasaranline struct {
		Sdata                 string `json:"sdata" `
		Master                string `json:"master" `
		Company               string `json:"company" `
		Companypasaran_id     int    `json:"companypasaran_id" `
		Pasaran_id            string `json:"pasaran_id" `
		Pasaran_limitline_4d  int    `json:"pasaran_limitline_4d" `
		Pasaran_limitline_3d  int    `json:"pasaran_limitline_3d" `
		Pasaran_limitline_3dd int    `json:"pasaran_limitline_3dd" `
		Pasaran_limitline_2d  int    `json:"pasaran_limitline_2d" `
		Pasaran_limitline_2dd int    `json:"pasaran_limitline_2dd" `
		Pasaran_limitline_2dt int    `json:"pasaran_limitline_2dt" `
		Pasaran_bbfs          int    `json:"pasaran_bbfs" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_updatepasaranline)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":       hostname,
			"master":                client.Master,
			"company":               client.Company,
			"companypasaran_id":     client.Companypasaran_id,
			"pasaran_id":            client.Pasaran_id,
			"pasaran_limitline_4d":  client.Pasaran_limitline_4d,
			"pasaran_limitline_3d":  client.Pasaran_limitline_3d,
			"pasaran_limitline_3dd": client.Pasaran_limitline_3dd,
			"pasaran_limitline_2d":  client.Pasaran_limitline_2d,
			"pasaran_limitline_2dd": client.Pasaran_limitline_2dd,
			"pasaran_limitline_2dt": client.Pasaran_limitline_2dt,
			"pasaran_bbfs":          client.Pasaran_bbfs,
		}).
		Post(PATH + "api/savecompanyupdatepasaranline")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companysavepasaranonline(c *fiber.Ctx) error {
	type payload_savepasaranonline struct {
		Sdata             string `json:"sdata" `
		Master            string `json:"master" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id" `
		Pasaran_hari      string `json:"pasaran_hari" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_savepasaranonline)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_hari":      client.Pasaran_hari,
		}).
		Post(PATH + "api/savecompanypasaranonline")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companydeletepasaranonline(c *fiber.Ctx) error {
	type payload_deletepasaranonline struct {
		Company              string `json:"company"`
		Master               string `json:"master"`
		Companypasaran_id    int    `json:"companypasaran_id"`
		Companypasaran_idoff int    `json:"companypasaran_idoffline"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_deletepasaranonline)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	log.Printf("%s-%s-%d-%d", client.Company, client.Master, client.Companypasaran_id, client.Companypasaran_idoff)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":          hostname,
			"master":                   client.Master,
			"company":                  client.Company,
			"companypasaran_id":        client.Companypasaran_id,
			"companypasaran_idoffline": client.Companypasaran_idoff,
		}).
		Post(PATH + "api/deletecompanypasaranonline")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasaran432(c *fiber.Ctx) error {
	type payload_pasaran432 struct {
		Sdata                       string  `json:"sdata" `
		Master                      string  `json:"master" `
		Company                     string  `json:"company" `
		Pasaran_id                  string  `json:"pasaran_id" `
		Companypasaran_id           int     `json:"companypasaran_id" `
		Pasaran_minbet_432d         int     `json:"pasaran_minbet_432d" `
		Pasaran_maxbet4d_432d       int     `json:"pasaran_maxbet4d_432d" `
		Pasaran_maxbet3d_432d       int     `json:"pasaran_maxbet3d_432d" `
		Pasaran_maxbet3dd_432d      int     `json:"pasaran_maxbet3dd_432d" `
		Pasaran_maxbet2d_432d       int     `json:"pasaran_maxbet2d_432d" `
		Pasaran_maxbet2dd_432d      int     `json:"pasaran_maxbet2dd_432d" `
		Pasaran_maxbet2dt_432d      int     `json:"pasaran_maxbet2dt_432d" `
		Pasaran_limitotal4d_432d    int     `json:"pasaran_limitotal4d_432d" `
		Pasaran_limitotal3d_432d    int     `json:"pasaran_limitotal3d_432d" `
		Pasaran_limitotal3dd_432d   int     `json:"pasaran_limitotal3dd_432d" `
		Pasaran_limitotal2d_432d    int     `json:"pasaran_limitotal2d_432d" `
		Pasaran_limitotal2dd_432d   int     `json:"pasaran_limitotal2dd_432d" `
		Pasaran_limitotal2dt_432d   int     `json:"pasaran_limitotal2dt_432d" `
		Pasaran_limitglobal4d_432d  int     `json:"pasaran_limitglobal4d_432d" `
		Pasaran_limitglobal3d_432d  int     `json:"pasaran_limitglobal3d_432d" `
		Pasaran_limitglobal3dd_432d int     `json:"pasaran_limitglobal3dd_432d" `
		Pasaran_limitglobal2d_432d  int     `json:"pasaran_limitglobal2d_432d" `
		Pasaran_limitglobal2dd_432d int     `json:"pasaran_limitglobal2dd_432d" `
		Pasaran_limitglobal2dt_432d int     `json:"pasaran_limitglobal2dt_432d"`
		Pasaran_win4d_432d          int     `json:"pasaran_win4d_432d" `
		Pasaran_win3d_432d          int     `json:"pasaran_win3d_432d" `
		Pasaran_win3dd_432d         int     `json:"pasaran_win3dd_432d" `
		Pasaran_win2d_432d          int     `json:"pasaran_win2d_432d" `
		Pasaran_win2dd_432d         int     `json:"pasaran_win2dd_432d" `
		Pasaran_win2dt_432d         int     `json:"pasaran_win2dt_432d" `
		Pasaran_win4dnodisc_432d    int     `json:"pasaran_win4dnodisc_432d" `
		Pasaran_win3dnodisc_432d    int     `json:"pasaran_win3dnodisc_432d" `
		Pasaran_win3ddnodisc_432d   int     `json:"pasaran_win3ddnodisc_432d" `
		Pasaran_win2dnodisc_432d    int     `json:"pasaran_win2dnodisc_432d" `
		Pasaran_win2ddnodisc_432d   int     `json:"pasaran_win2ddnodisc_432d" `
		Pasaran_win2dtnodisc_432d   int     `json:"pasaran_win2dtnodisc_432d" `
		Pasaran_win4d_bb_432d       int     `json:"pasaran_win4d_bb_432d" `
		Pasaran_win3d_bb_432d       int     `json:"pasaran_win3d_bb_432d" `
		Pasaran_win3dd_bb_432d      int     `json:"pasaran_win3dd_bb_432d" `
		Pasaran_win2d_bb_432d       int     `json:"pasaran_win2d_bb_432d" `
		Pasaran_win2dd_bb_432d      int     `json:"pasaran_win2dd_bb_432d" `
		Pasaran_win2dt_bb_432d      int     `json:"pasaran_win2dt_bb_432d" `
		Pasaran_win4d_bb_kena_432d  int     `json:"pasaran_win4d_bb_kena_432d" `
		Pasaran_win3d_bb_kena_432d  int     `json:"pasaran_win3d_bb_kena_432d" `
		Pasaran_win3dd_bb_kena_432d int     `json:"pasaran_win3dd_bb_kena_432d" `
		Pasaran_win2d_bb_kena_432d  int     `json:"pasaran_win2d_bb_kena_432d" `
		Pasaran_win2dd_bb_kena_432d int     `json:"pasaran_win2dd_bb_kena_432d" `
		Pasaran_win2dt_bb_kena_432d int     `json:"pasaran_win2dt_bb_kena_432d" `
		Pasaran_disc4d_432d         float32 `json:"pasaran_disc4d_432d" `
		Pasaran_disc3d_432d         float32 `json:"pasaran_disc3d_432d" `
		Pasaran_disc3dd_432d        float32 `json:"pasaran_disc3dd_432d" `
		Pasaran_disc2d_432d         float32 `json:"pasaran_disc2d_432d" `
		Pasaran_disc2dd_432d        float32 `json:"pasaran_disc2dd_432d" `
		Pasaran_disc2dt_432d        float32 `json:"pasaran_disc2dt_432d" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaran432)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":             hostname,
			"master":                      client.Master,
			"company":                     client.Company,
			"companypasaran_id":           client.Companypasaran_id,
			"pasaran_id":                  client.Pasaran_id,
			"pasaran_minbet_432d":         client.Pasaran_minbet_432d,
			"pasaran_maxbet4d_432d":       client.Pasaran_maxbet4d_432d,
			"pasaran_maxbet3d_432d":       client.Pasaran_maxbet3d_432d,
			"pasaran_maxbet3dd_432d":      client.Pasaran_maxbet3dd_432d,
			"pasaran_maxbet2d_432d":       client.Pasaran_maxbet2d_432d,
			"pasaran_maxbet2dd_432d":      client.Pasaran_maxbet2dd_432d,
			"pasaran_maxbet2dt_432d":      client.Pasaran_maxbet2dt_432d,
			"pasaran_limitotal4d_432d":    client.Pasaran_limitotal4d_432d,
			"pasaran_limitotal3d_432d":    client.Pasaran_limitotal3d_432d,
			"pasaran_limitotal3dd_432d":   client.Pasaran_limitotal3dd_432d,
			"pasaran_limitotal2d_432d":    client.Pasaran_limitotal2d_432d,
			"pasaran_limitotal2dd_432d":   client.Pasaran_limitotal2dd_432d,
			"pasaran_limitotal2dt_432d":   client.Pasaran_limitotal2dt_432d,
			"pasaran_limitglobal4d_432d":  client.Pasaran_limitglobal4d_432d,
			"pasaran_limitglobal3d_432d":  client.Pasaran_limitglobal3d_432d,
			"pasaran_limitglobal3dd_432d": client.Pasaran_limitglobal3dd_432d,
			"pasaran_limitglobal2d_432d":  client.Pasaran_limitglobal2d_432d,
			"pasaran_limitglobal2dd_432d": client.Pasaran_limitglobal2dd_432d,
			"pasaran_limitglobal2dt_432d": client.Pasaran_limitglobal2dt_432d,
			"pasaran_win4d_432d":          client.Pasaran_win4d_432d,
			"pasaran_win3d_432d":          client.Pasaran_win3d_432d,
			"pasaran_win3dd_432d":         client.Pasaran_win3dd_432d,
			"pasaran_win2d_432d":          client.Pasaran_win2d_432d,
			"pasaran_win2dd_432d":         client.Pasaran_win2dd_432d,
			"pasaran_win2dt_432d":         client.Pasaran_win2dt_432d,
			"pasaran_win4dnodisc_432d":    client.Pasaran_win4dnodisc_432d,
			"pasaran_win3dnodisc_432d":    client.Pasaran_win3dnodisc_432d,
			"pasaran_win3ddnodisc_432d":   client.Pasaran_win3ddnodisc_432d,
			"pasaran_win2dnodisc_432d":    client.Pasaran_win2dnodisc_432d,
			"pasaran_win2ddnodisc_432d":   client.Pasaran_win2ddnodisc_432d,
			"pasaran_win2dtnodisc_432d":   client.Pasaran_win2dtnodisc_432d,
			"pasaran_win4dbb_kena_432d":   client.Pasaran_win4d_bb_kena_432d,
			"pasaran_win3dbb_kena_432d":   client.Pasaran_win3d_bb_kena_432d,
			"pasaran_win3ddbb_kena_432d":  client.Pasaran_win3dd_bb_kena_432d,
			"pasaran_win2dbb_kena_432d":   client.Pasaran_win2d_bb_kena_432d,
			"pasaran_win2ddbb_kena_432d":  client.Pasaran_win2dd_bb_kena_432d,
			"pasaran_win2dtbb_kena_432d":  client.Pasaran_win2dt_bb_kena_432d,
			"pasaran_win4dbb_432d":        client.Pasaran_win4d_bb_432d,
			"pasaran_win3dbb_432d":        client.Pasaran_win3d_bb_432d,
			"pasaran_win3ddbb_432d":       client.Pasaran_win3dd_bb_432d,
			"pasaran_win2dbb_432d":        client.Pasaran_win2d_bb_432d,
			"pasaran_win2ddbb_432d":       client.Pasaran_win2dd_bb_432d,
			"pasaran_win2dtbb_432d":       client.Pasaran_win2dt_bb_432d,
			"pasaran_disc4d_432d":         client.Pasaran_disc4d_432d,
			"pasaran_disc3d_432d":         client.Pasaran_disc3d_432d,
			"pasaran_disc3dd_432d":        client.Pasaran_disc3dd_432d,
			"pasaran_disc2d_432d":         client.Pasaran_disc2d_432d,
			"pasaran_disc2dd_432d":        client.Pasaran_disc2dd_432d,
			"pasaran_disc2dt_432d":        client.Pasaran_disc2dt_432d,
		}).
		Post(PATH + "api/savecompanyupdatepasaran432")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasarancolokbebas(c *fiber.Ctx) error {
	type payload_pasarancolokbebas struct {
		Sdata                      string  `json:"sdata" `
		Master                     string  `json:"master" `
		Company                    string  `json:"company" `
		Pasaran_id                 string  `json:"pasaran_id" `
		Companypasaran_id          int     `json:"companypasaran_id" `
		Pasaran_minbet_cbebas      int     `json:"pasaran_minbet_cbebas" `
		Pasaran_maxbet_cbebas      int     `json:"pasaran_maxbet_cbebas" `
		Pasaran_limitotal_cbebas   int     `json:"pasaran_limitotal_cbebas" `
		Pasaran_limitglobal_cbebas int     `json:"pasaran_limitglobal_cbebas" `
		Pasaran_win_cbebas         float32 `json:"pasaran_win_cbebas" `
		Pasaran_disc_cbebas        float32 `json:"pasaran_disc_cbebas" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarancolokbebas)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":            hostname,
			"master":                     client.Master,
			"company":                    client.Company,
			"companypasaran_id":          client.Companypasaran_id,
			"pasaran_id":                 client.Pasaran_id,
			"pasaran_minbet_cbebas":      client.Pasaran_minbet_cbebas,
			"pasaran_maxbet_cbebas":      client.Pasaran_maxbet_cbebas,
			"pasaran_limitotal_cbebas":   client.Pasaran_limitotal_cbebas,
			"pasaran_limitglobal_cbebas": client.Pasaran_limitglobal_cbebas,
			"pasaran_win_cbebas":         client.Pasaran_win_cbebas,
			"pasaran_disc_cbebas":        client.Pasaran_disc_cbebas,
		}).
		Post(PATH + "api/savecompanyupdatepasarancolokbebas")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasarancolokmacau(c *fiber.Ctx) error {
	type payload_pasarancolokmacau struct {
		Sdata                      string  `json:"sdata" `
		Master                     string  `json:"master" `
		Company                    string  `json:"company" `
		Pasaran_id                 string  `json:"pasaran_id" `
		Companypasaran_id          int     `json:"companypasaran_id" `
		Pasaran_minbet_cmacau      int     `json:"pasaran_minbet_cmacau" `
		Pasaran_maxbet_cmacau      int     `json:"pasaran_maxbet_cmacau" `
		Pasaran_limitotal_cmacau   int     `json:"pasaran_limitotal_cmacau" `
		Pasaran_limitglobal_cmacau int     `json:"pasaran_limitglobal_cmacau" `
		Pasaran_win2_cmacau        float32 `json:"pasaran_win2_cmacau"`
		Pasaran_win3_cmacau        float32 `json:"pasaran_win3_cmacau"`
		Pasaran_win4_cmacau        float32 `json:"pasaran_win4_cmacau"`
		Pasaran_disc_cmacau        float32 `json:"pasaran_disc_cmacau"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarancolokmacau)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":            hostname,
			"master":                     client.Master,
			"company":                    client.Company,
			"companypasaran_id":          client.Companypasaran_id,
			"pasaran_id":                 client.Pasaran_id,
			"pasaran_minbet_cmacau":      client.Pasaran_minbet_cmacau,
			"pasaran_maxbet_cmacau":      client.Pasaran_maxbet_cmacau,
			"pasaran_limitotal_cmacau":   client.Pasaran_limitotal_cmacau,
			"pasaran_limitglobal_cmacau": client.Pasaran_limitglobal_cmacau,
			"pasaran_win2_cmacau":        client.Pasaran_win2_cmacau,
			"pasaran_win3_cmacau":        client.Pasaran_win3_cmacau,
			"pasaran_win4_cmacau":        client.Pasaran_win4_cmacau,
			"pasaran_disc_cmacau":        client.Pasaran_disc_cmacau,
		}).
		Post(PATH + "api/savecompanyupdatepasarancolokmacau")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasarancoloknaga(c *fiber.Ctx) error {
	type payload_pasarancoloknaga struct {
		Sdata                     string  `json:"sdata" `
		Master                    string  `json:"master" `
		Company                   string  `json:"company" `
		Pasaran_id                string  `json:"pasaran_id" `
		Companypasaran_id         int     `json:"companypasaran_id" `
		Pasaran_minbet_cnaga      int     `json:"pasaran_minbet_cnaga" `
		Pasaran_maxbet_cnaga      int     `json:"pasaran_maxbet_cnaga" `
		Pasaran_limittotal_cnaga  int     `json:"pasaran_limittotal_cnaga" `
		Pasaran_limitglobal_cnaga int     `json:"pasaran_limitglobal_cnaga" `
		Pasaran_win3_cnaga        float32 `json:"pasaran_win3_cnaga" `
		Pasaran_win4_cnaga        float32 `json:"pasaran_win4_cnaga" `
		Pasaran_disc_cnaga        float32 `json:"pasaran_disc_cnaga" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarancoloknaga)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":           hostname,
			"master":                    client.Master,
			"company":                   client.Company,
			"companypasaran_id":         client.Companypasaran_id,
			"pasaran_id":                client.Pasaran_id,
			"pasaran_minbet_cnaga":      client.Pasaran_minbet_cnaga,
			"pasaran_maxbet_cnaga":      client.Pasaran_maxbet_cnaga,
			"pasaran_limittotal_cnaga":  client.Pasaran_limittotal_cnaga,
			"pasaran_limitglobal_cnaga": client.Pasaran_limitglobal_cnaga,
			"pasaran_win3_cnaga":        client.Pasaran_win3_cnaga,
			"pasaran_win4_cnaga":        client.Pasaran_win4_cnaga,
			"pasaran_disc_cnaga":        client.Pasaran_disc_cnaga,
		}).
		Post(PATH + "api/savecompanyupdatepasarancoloknaga")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasarancolokjitu(c *fiber.Ctx) error {
	type payload_pasarancolokjitu struct {
		Sdata                     string  `json:"sdata" `
		Master                    string  `json:"master" `
		Company                   string  `json:"company" `
		Pasaran_id                string  `json:"pasaran_id" `
		Companypasaran_id         int     `json:"companypasaran_id" `
		Pasaran_minbet_cjitu      int     `json:"pasaran_minbet_cjitu" `
		Pasaran_maxbet_cjitu      int     `json:"pasaran_maxbet_cjitu" `
		Pasaran_limittotal_cjitu  int     `json:"pasaran_limittotal_cjitu" `
		Pasaran_limitglobal_cjitu int     `json:"pasaran_limitglobal_cjitu" `
		Pasaran_winas_cjitu       float32 `json:"pasaran_winas_cjitu" `
		Pasaran_winkop_cjitu      float32 `json:"pasaran_winkop_cjitu" `
		Pasaran_winkepala_cjitu   float32 `json:"pasaran_winkepala_cjitu" `
		Pasaran_winekor_cjitu     float32 `json:"pasaran_winekor_cjitu"`
		Pasaran_desc_cjitu        float32 `json:"pasaran_desc_cjitu" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarancolokjitu)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":           hostname,
			"master":                    client.Master,
			"company":                   client.Company,
			"companypasaran_id":         client.Companypasaran_id,
			"pasaran_id":                client.Pasaran_id,
			"pasaran_minbet_cjitu":      client.Pasaran_minbet_cjitu,
			"pasaran_maxbet_cjitu":      client.Pasaran_maxbet_cjitu,
			"pasaran_limittotal_cjitu":  client.Pasaran_limittotal_cjitu,
			"pasaran_limitglobal_cjitu": client.Pasaran_limitglobal_cjitu,
			"pasaran_winas_cjitu":       client.Pasaran_winas_cjitu,
			"pasaran_winkop_cjitu":      client.Pasaran_winkop_cjitu,
			"pasaran_winkepala_cjitu":   client.Pasaran_winkepala_cjitu,
			"pasaran_winekor_cjitu":     client.Pasaran_winekor_cjitu,
			"pasaran_desc_cjitu":        client.Pasaran_desc_cjitu,
		}).
		Post(PATH + "api/savecompanyupdatepasarancolokjitu")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasaran5050umum(c *fiber.Ctx) error {
	type payload_pasaran5050umum struct {
		Sdata                        string  `json:"sdata" `
		Master                       string  `json:"master" `
		Company                      string  `json:"company" `
		Pasaran_id                   string  `json:"pasaran_id" `
		Companypasaran_id            int     `json:"companypasaran_id" `
		Pasaran_minbet_5050umum      int     `json:"pasaran_minbet_5050umum" `
		Pasaran_maxbet_5050umum      int     `json:"pasaran_maxbet_5050umum" `
		Pasaran_limittotal_5050umum  int     `json:"pasaran_limittotal_5050umum" `
		Pasaran_limitglobal_5050umum int     `json:"pasaran_limitglobal_5050umum" `
		Pasaran_keibesar_5050umum    float64 `json:"pasaran_keibesar_5050umum" `
		Pasaran_keikecil_5050umum    float64 `json:"pasaran_keikecil_5050umum" `
		Pasaran_keigenap_5050umum    float64 `json:"pasaran_keigenap_5050umum" `
		Pasaran_keiganjil_5050umum   float64 `json:"pasaran_keiganjil_5050umum" `
		Pasaran_keitengah_5050umum   float64 `json:"pasaran_keitengah_5050umum" `
		Pasaran_keitepi_5050umum     float64 `json:"pasaran_keitepi_5050umum" `
		Pasaran_discbesar_5050umum   float64 `json:"pasaran_discbesar_5050umum" `
		Pasaran_disckecil_5050umum   float64 `json:"pasaran_disckecil_5050umum" `
		Pasaran_discgenap_5050umum   float64 `json:"pasaran_discgenap_5050umum" `
		Pasaran_discganjil_5050umum  float64 `json:"pasaran_discganjil_5050umum" `
		Pasaran_disctengah_5050umum  float64 `json:"pasaran_disctengah_5050umum" `
		Pasaran_disctepi_5050umum    float64 `json:"pasaran_disctepi_5050umum" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaran5050umum)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":              hostname,
			"master":                       client.Master,
			"company":                      client.Company,
			"companypasaran_id":            client.Companypasaran_id,
			"pasaran_id":                   client.Pasaran_id,
			"pasaran_minbet_5050umum":      client.Pasaran_minbet_5050umum,
			"pasaran_maxbet_5050umum":      client.Pasaran_maxbet_5050umum,
			"pasaran_limittotal_5050umum":  client.Pasaran_limittotal_5050umum,
			"pasaran_limitglobal_5050umum": client.Pasaran_limitglobal_5050umum,
			"pasaran_keibesar_5050umum":    client.Pasaran_keibesar_5050umum,
			"pasaran_keikecil_5050umum":    client.Pasaran_keikecil_5050umum,
			"pasaran_keigenap_5050umum":    client.Pasaran_keigenap_5050umum,
			"pasaran_keiganjil_5050umum":   client.Pasaran_keiganjil_5050umum,
			"pasaran_keitengah_5050umum":   client.Pasaran_keitengah_5050umum,
			"pasaran_keitepi_5050umum":     client.Pasaran_keitepi_5050umum,
			"pasaran_discbesar_5050umum":   client.Pasaran_discbesar_5050umum,
			"pasaran_disckecil_5050umum":   client.Pasaran_disckecil_5050umum,
			"pasaran_discgenap_5050umum":   client.Pasaran_discgenap_5050umum,
			"pasaran_discganjil_5050umum":  client.Pasaran_discganjil_5050umum,
			"pasaran_disctengah_5050umum":  client.Pasaran_disctengah_5050umum,
			"pasaran_disctepi_5050umum":    client.Pasaran_disctepi_5050umum,
		}).
		Post(PATH + "api/savecompanyupdatepasaran5050umum")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasaran5050special(c *fiber.Ctx) error {
	type payload_pasaran5050special struct {
		Sdata                                string  `json:"sdata" `
		Master                               string  `json:"master" `
		Company                              string  `json:"company" `
		Pasaran_id                           string  `json:"pasaran_id" `
		Companypasaran_id                    int     `json:"companypasaran_id" `
		Pasaran_minbet_5050special           int     `json:"pasaran_minbet_5050special" `
		Pasaran_maxbet_5050special           int     `json:"pasaran_maxbet_5050special" `
		Pasaran_limitglobal_5050special      int     `json:"pasaran_limitglobal_5050special" "`
		Pasaran_limittotal_5050special       int     `json:"pasaran_limittotal_5050special" `
		Pasaran_keiasganjil_5050special      float64 `json:"pasaran_keiasganjil_5050special" `
		Pasaran_keiasgenap_5050special       float64 `json:"pasaran_keiasgenap_5050special" `
		Pasaran_keiasbesar_5050special       float64 `json:"pasaran_keiasbesar_5050special" `
		Pasaran_keiaskecil_5050special       float64 `json:"pasaran_keiaskecil_5050special" `
		Pasaran_keikopganjil_5050special     float64 `json:"pasaran_keikopganjil_5050special" `
		Pasaran_keikopgenap_5050special      float64 `json:"pasaran_keikopgenap_5050special" `
		Pasaran_keikopbesar_5050special      float64 `json:"pasaran_keikopbesar_5050special" `
		Pasaran_keikopkecil_5050special      float64 `json:"pasaran_keikopkecil_5050special" `
		Pasaran_keikepalaganjil_5050special  float64 `json:"pasaran_keikepalaganjil_5050special" `
		Pasaran_keikepalagenap_5050special   float64 `json:"pasaran_keikepalagenap_5050special" `
		Pasaran_keikepalabesar_5050special   float64 `json:"pasaran_keikepalabesar_5050special" `
		Pasaran_keikepalakecil_5050special   float64 `json:"pasaran_keikepalakecil_5050special" `
		Pasaran_keiekorganjil_5050special    float64 `json:"pasaran_keiekorganjil_5050special" `
		Pasaran_keiekorgenap_5050special     float64 `json:"pasaran_keiekorgenap_5050special" `
		Pasaran_keiekorbesar_5050special     float64 `json:"pasaran_keiekorbesar_5050special" `
		Pasaran_keiekorkecil_5050special     float64 `json:"pasaran_keiekorkecil_5050special" `
		Pasaran_discasganjil_5050special     float64 `json:"pasaran_discasganjil_5050special" `
		Pasaran_discasgenap_5050special      float64 `json:"pasaran_discasgenap_5050special" `
		Pasaran_discasbesar_5050special      float64 `json:"pasaran_discasbesar_5050special" `
		Pasaran_discaskecil_5050special      float64 `json:"pasaran_discaskecil_5050special" `
		Pasaran_disckopganjil_5050special    float64 `json:"pasaran_disckopganjil_5050special" `
		Pasaran_disckopgenap_5050special     float64 `json:"pasaran_disckopgenap_5050special" `
		Pasaran_disckopbesar_5050special     float64 `json:"pasaran_disckopbesar_5050special" `
		Pasaran_disckopkecil_5050special     float64 `json:"pasaran_disckopkecil_5050special" `
		Pasaran_disckepalaganjil_5050special float64 `json:"pasaran_disckepalaganjil_5050special" `
		Pasaran_disckepalagenap_5050special  float64 `json:"pasaran_disckepalagenap_5050special" `
		Pasaran_disckepalabesar_5050special  float64 `json:"pasaran_disckepalabesar_5050special" `
		Pasaran_disckepalakecil_5050special  float64 `json:"pasaran_disckepalakecil_5050special" `
		Pasaran_discekorganjil_5050special   float64 `json:"pasaran_discekorganjil_5050special" `
		Pasaran_discekorgenap_5050special    float64 `json:"pasaran_discekorgenap_5050special" `
		Pasaran_discekorbesar_5050special    float64 `json:"pasaran_discekorbesar_5050special" `
		Pasaran_discekorkecil_5050special    float64 `json:"pasaran_discekorkecil_5050special" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaran5050special)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":                      hostname,
			"master":                               client.Master,
			"company":                              client.Company,
			"companypasaran_id":                    client.Companypasaran_id,
			"pasaran_id":                           client.Pasaran_id,
			"pasaran_minbet_5050special":           client.Pasaran_minbet_5050special,
			"pasaran_maxbet_5050special":           client.Pasaran_maxbet_5050special,
			"pasaran_limitglobal_5050special":      client.Pasaran_limitglobal_5050special,
			"pasaran_limittotal_5050special":       client.Pasaran_limittotal_5050special,
			"pasaran_keiasganjil_5050special":      client.Pasaran_keiasganjil_5050special,
			"pasaran_keiasgenap_5050special":       client.Pasaran_keiasgenap_5050special,
			"pasaran_keiasbesar_5050special":       client.Pasaran_keiasbesar_5050special,
			"pasaran_keiaskecil_5050special":       client.Pasaran_keiaskecil_5050special,
			"pasaran_keikopganjil_5050special":     client.Pasaran_keikopganjil_5050special,
			"pasaran_keikopgenap_5050special":      client.Pasaran_keikopgenap_5050special,
			"pasaran_keikopbesar_5050special":      client.Pasaran_keikopbesar_5050special,
			"pasaran_keikopkecil_5050special":      client.Pasaran_keikopkecil_5050special,
			"pasaran_keikepalaganjil_5050special":  client.Pasaran_keikepalaganjil_5050special,
			"pasaran_keikepalagenap_5050special":   client.Pasaran_keikepalagenap_5050special,
			"pasaran_keikepalabesar_5050special":   client.Pasaran_keikepalabesar_5050special,
			"pasaran_keikepalakecil_5050special":   client.Pasaran_keikepalakecil_5050special,
			"pasaran_keiekorganjil_5050special":    client.Pasaran_keiekorganjil_5050special,
			"pasaran_keiekorgenap_5050special":     client.Pasaran_keiekorgenap_5050special,
			"pasaran_keiekorbesar_5050special":     client.Pasaran_keiekorbesar_5050special,
			"pasaran_keiekorkecil_5050special":     client.Pasaran_keiekorkecil_5050special,
			"pasaran_discasganjil_5050special":     client.Pasaran_discasganjil_5050special,
			"pasaran_discasgenap_5050special":      client.Pasaran_discasgenap_5050special,
			"pasaran_discasbesar_5050special":      client.Pasaran_discasbesar_5050special,
			"pasaran_discaskecil_5050special":      client.Pasaran_discaskecil_5050special,
			"pasaran_disckopganjil_5050special":    client.Pasaran_disckopganjil_5050special,
			"pasaran_disckopgenap_5050special":     client.Pasaran_disckopgenap_5050special,
			"pasaran_disckopbesar_5050special":     client.Pasaran_disckopbesar_5050special,
			"pasaran_disckopkecil_5050special":     client.Pasaran_disckopkecil_5050special,
			"pasaran_disckepalaganjil_5050special": client.Pasaran_disckepalaganjil_5050special,
			"pasaran_disckepalagenap_5050special":  client.Pasaran_disckepalagenap_5050special,
			"pasaran_disckepalabesar_5050special":  client.Pasaran_disckepalabesar_5050special,
			"pasaran_disckepalakecil_5050special":  client.Pasaran_disckepalakecil_5050special,
			"pasaran_discekorganjil_5050special":   client.Pasaran_discekorganjil_5050special,
			"pasaran_discekorgenap_5050special":    client.Pasaran_discekorgenap_5050special,
			"pasaran_discekorbesar_5050special":    client.Pasaran_discekorbesar_5050special,
			"pasaran_discekorkecil_5050special":    client.Pasaran_discekorkecil_5050special,
		}).
		Post(PATH + "api/savecompanyupdatepasaran5050special")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasaran5050kombinasi(c *fiber.Ctx) error {
	type payload_pasaran5050kombinasi struct {
		Sdata                                     string  `json:"sdata" `
		Master                                    string  `json:"master" `
		Company                                   string  `json:"company" `
		Pasaran_id                                string  `json:"pasaran_id" `
		Companypasaran_id                         int     `json:"companypasaran_id" `
		Pasaran_minbet_5050kombinasi              int     `json:"pasaran_minbet_5050kombinasi" `
		Pasaran_maxbet_5050kombinasi              int     `json:"pasaran_maxbet_5050kombinasi" `
		Pasaran_limitglobal_5050kombinasi         int     `json:"pasaran_limitglobal_5050kombinasi" `
		Pasaran_limittotal_5050kombinasi          int     `json:"pasaran_limittotal_5050kombinasi" `
		Pasaran_belakangkeimono_5050kombinasi     float64 `json:"pasaran_belakangkeimono_5050kombinasi" `
		Pasaran_belakangkeistereo_5050kombinasi   float64 `json:"pasaran_belakangkeistereo_5050kombinasi" `
		Pasaran_belakangkeikembang_5050kombinasi  float64 `json:"pasaran_belakangkeikembang_5050kombinasi" `
		Pasaran_belakangkeikempis_5050kombinasi   float64 `json:"pasaran_belakangkeikempis_5050kombinasi" `
		Pasaran_belakangkeikembar_5050kombinasi   float64 `json:"pasaran_belakangkeikembar_5050kombinasi" `
		Pasaran_tengahkeimono_5050kombinasi       float64 `json:"pasaran_tengahkeimono_5050kombinasi" `
		Pasaran_tengahkeistereo_5050kombinasi     float64 `json:"pasaran_tengahkeistereo_5050kombinasi" `
		Pasaran_tengahkeikembang_5050kombinasi    float64 `json:"pasaran_tengahkeikembang_5050kombinasi" `
		Pasaran_tengahkeikempis_5050kombinasi     float64 `json:"pasaran_tengahkeikempis_5050kombinasi" `
		Pasaran_tengahkeikembar_5050kombinasi     float64 `json:"pasaran_tengahkeikembar_5050kombinasi" `
		Pasaran_depankeimono_5050kombinasi        float64 `json:"pasaran_depankeimono_5050kombinasi" `
		Pasaran_depankeistereo_5050kombinasi      float64 `json:"pasaran_depankeistereo_5050kombinasi" `
		Pasaran_depankeikembang_5050kombinasi     float64 `json:"pasaran_depankeikembang_5050kombinasi" `
		Pasaran_depankeikempis_5050kombinasi      float64 `json:"pasaran_depankeikempis_5050kombinasi" `
		Pasaran_depankeikembar_5050kombinasi      float64 `json:"pasaran_depankeikembar_5050kombinasi" `
		Pasaran_belakangdiscmono_5050kombinasi    float64 `json:"pasaran_belakangdiscmono_5050kombinasi" `
		Pasaran_belakangdiscstereo_5050kombinasi  float64 `json:"pasaran_belakangdiscstereo_5050kombinasi" `
		Pasaran_belakangdisckembang_5050kombinasi float64 `json:"pasaran_belakangdisckembang_5050kombinasi"`
		Pasaran_belakangdisckempis_5050kombinasi  float64 `json:"pasaran_belakangdisckempis_5050kombinasi" `
		Pasaran_belakangdisckembar_5050kombinasi  float64 `json:"pasaran_belakangdisckembar_5050kombinasi" `
		Pasaran_tengahdiscmono_5050kombinasi      float64 `json:"pasaran_tengahdiscmono_5050kombinasi" `
		Pasaran_tengahdiscstereo_5050kombinasi    float64 `json:"pasaran_tengahdiscstereo_5050kombinasi" `
		Pasaran_tengahdisckembang_5050kombinasi   float64 `json:"pasaran_tengahdisckembang_5050kombinasi" `
		Pasaran_tengahdisckempis_5050kombinasi    float64 `json:"pasaran_tengahdisckempis_5050kombinasi" `
		Pasaran_tengahdisckembar_5050kombinasi    float64 `json:"pasaran_tengahdisckembar_5050kombinasi" `
		Pasaran_depandiscmono_5050kombinasi       float64 `json:"pasaran_depandiscmono_5050kombinasi" `
		Pasaran_depandiscstereo_5050kombinasi     float64 `json:"pasaran_depandiscstereo_5050kombinasi" `
		Pasaran_depandisckembang_5050kombinasi    float64 `json:"pasaran_depandisckembang_5050kombinasi" `
		Pasaran_depandisckempis_5050kombinasi     float64 `json:"pasaran_depandisckempis_5050kombinasi" `
		Pasaran_depandisckembar_5050kombinasi     float64 `json:"pasaran_depandisckembar_5050kombinasi" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaran5050kombinasi)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":                           hostname,
			"master":                                    client.Master,
			"company":                                   client.Company,
			"companypasaran_id":                         client.Companypasaran_id,
			"pasaran_id":                                client.Pasaran_id,
			"pasaran_minbet_5050kombinasi":              client.Pasaran_minbet_5050kombinasi,
			"pasaran_maxbet_5050kombinasi":              client.Pasaran_maxbet_5050kombinasi,
			"pasaran_limitglobal_5050kombinasi":         client.Pasaran_limitglobal_5050kombinasi,
			"pasaran_limittotal_5050kombinasi":          client.Pasaran_limittotal_5050kombinasi,
			"pasaran_belakangkeimono_5050kombinasi":     client.Pasaran_belakangkeimono_5050kombinasi,
			"pasaran_belakangkeistereo_5050kombinasi":   client.Pasaran_belakangkeistereo_5050kombinasi,
			"pasaran_belakangkeikembang_5050kombinasi":  client.Pasaran_belakangkeikembang_5050kombinasi,
			"pasaran_belakangkeikempis_5050kombinasi":   client.Pasaran_belakangkeikempis_5050kombinasi,
			"pasaran_belakangkeikembar_5050kombinasi":   client.Pasaran_belakangkeikembar_5050kombinasi,
			"pasaran_tengahkeimono_5050kombinasi":       client.Pasaran_tengahkeimono_5050kombinasi,
			"pasaran_tengahkeistereo_5050kombinasi":     client.Pasaran_tengahkeistereo_5050kombinasi,
			"pasaran_tengahkeikembang_5050kombinasi":    client.Pasaran_tengahkeikembang_5050kombinasi,
			"pasaran_tengahkeikempis_5050kombinasi":     client.Pasaran_tengahkeikempis_5050kombinasi,
			"pasaran_tengahkeikembar_5050kombinasi":     client.Pasaran_tengahkeikembar_5050kombinasi,
			"pasaran_depankeimono_5050kombinasi":        client.Pasaran_depankeimono_5050kombinasi,
			"pasaran_depankeistereo_5050kombinasi":      client.Pasaran_depankeistereo_5050kombinasi,
			"pasaran_depankeikembang_5050kombinasi":     client.Pasaran_depankeikembang_5050kombinasi,
			"pasaran_depankeikempis_5050kombinasi":      client.Pasaran_depankeikempis_5050kombinasi,
			"pasaran_depankeikembar_5050kombinasi":      client.Pasaran_depankeikembar_5050kombinasi,
			"pasaran_belakangdiscmono_5050kombinasi":    client.Pasaran_belakangdiscmono_5050kombinasi,
			"pasaran_belakangdiscstereo_5050kombinasi":  client.Pasaran_belakangdiscstereo_5050kombinasi,
			"pasaran_belakangdisckembang_5050kombinasi": client.Pasaran_belakangdisckembang_5050kombinasi,
			"pasaran_belakangdisckempis_5050kombinasi":  client.Pasaran_belakangdisckempis_5050kombinasi,
			"pasaran_belakangdisckembar_5050kombinasi":  client.Pasaran_belakangdisckembar_5050kombinasi,
			"pasaran_tengahdiscmono_5050kombinasi":      client.Pasaran_tengahdiscmono_5050kombinasi,
			"pasaran_tengahdiscstereo_5050kombinasi":    client.Pasaran_tengahdiscstereo_5050kombinasi,
			"pasaran_tengahdisckembang_5050kombinasi":   client.Pasaran_tengahdisckembang_5050kombinasi,
			"pasaran_tengahdisckempis_5050kombinasi":    client.Pasaran_tengahdisckempis_5050kombinasi,
			"pasaran_tengahdisckembar_5050kombinasi":    client.Pasaran_tengahdisckembar_5050kombinasi,
			"pasaran_depandiscmono_5050kombinasi":       client.Pasaran_depandiscmono_5050kombinasi,
			"pasaran_depandiscstereo_5050kombinasi":     client.Pasaran_depandiscstereo_5050kombinasi,
			"pasaran_depandisckembang_5050kombinasi":    client.Pasaran_depandisckembang_5050kombinasi,
			"pasaran_depandisckempis_5050kombinasi":     client.Pasaran_depandisckempis_5050kombinasi,
			"pasaran_depandisckembar_5050kombinasi":     client.Pasaran_depandisckembar_5050kombinasi,
		}).
		Post(PATH + "api/savecompanyupdatepasaran5050kombinasi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasarankombinasi(c *fiber.Ctx) error {
	type payload_pasarankombinasi struct {
		Sdata                         string  `json:"sdata" `
		Master                        string  `json:"master" `
		Company                       string  `json:"company" `
		Pasaran_id                    string  `json:"pasaran_id" `
		Companypasaran_id             int     `json:"companypasaran_id" `
		Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi" `
		Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi" `
		Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi" `
		Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi" `
		Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi" `
		Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarankombinasi)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":               hostname,
			"master":                        client.Master,
			"company":                       client.Company,
			"companypasaran_id":             client.Companypasaran_id,
			"pasaran_id":                    client.Pasaran_id,
			"pasaran_minbet_kombinasi":      client.Pasaran_minbet_kombinasi,
			"pasaran_maxbet_kombinasi":      client.Pasaran_maxbet_kombinasi,
			"pasaran_limitglobal_kombinasi": client.Pasaran_limitglobal_kombinasi,
			"pasaran_limittotal_kombinasi":  client.Pasaran_limittotal_kombinasi,
			"pasaran_win_kombinasi":         client.Pasaran_win_kombinasi,
			"pasaran_disc_kombinasi":        client.Pasaran_disc_kombinasi,
		}).
		Post(PATH + "api/savecompanyupdatepasarankombinasi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasarandasar(c *fiber.Ctx) error {
	type payload_pasarandasar struct {
		Sdata                     string  `json:"sdata" `
		Master                    string  `json:"master" `
		Company                   string  `json:"company" `
		Pasaran_id                string  `json:"pasaran_id" `
		Companypasaran_id         int     `json:"companypasaran_id" `
		Pasaran_minbet_dasar      int     `json:"pasaran_minbet_dasar" `
		Pasaran_maxbet_dasar      int     `json:"pasaran_maxbet_dasar" `
		Pasaran_limitglobal_dasar int     `json:"pasaran_limitglobal_dasar" `
		Pasaran_limittotal_dasar  int     `json:"pasaran_limittotal_dasar" `
		Pasaran_keibesar_dasar    float32 `json:"pasaran_keibesar_dasar" `
		Pasaran_keikecil_dasar    float32 `json:"pasaran_keikecil_dasar" `
		Pasaran_keigenap_dasar    float32 `json:"pasaran_keigenap_dasar" `
		Pasaran_keiganjil_dasar   float32 `json:"pasaran_keiganjil_dasar" `
		Pasaran_discbesar_dasar   float32 `json:"pasaran_discbesar_dasar" `
		Pasaran_disckecil_dasar   float32 `json:"pasaran_disckecil_dasar" `
		Pasaran_discgenap_dasar   float32 `json:"pasaran_discgenap_dasar" `
		Pasaran_discganjil_dasar  float32 `json:"pasaran_discganjil_dasar" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarandasar)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":           hostname,
			"master":                    client.Master,
			"company":                   client.Company,
			"companypasaran_id":         client.Companypasaran_id,
			"pasaran_id":                client.Pasaran_id,
			"pasaran_minbet_dasar":      client.Pasaran_minbet_dasar,
			"pasaran_maxbet_dasar":      client.Pasaran_maxbet_dasar,
			"pasaran_limitglobal_dasar": client.Pasaran_limitglobal_dasar,
			"pasaran_limittotal_dasar":  client.Pasaran_limittotal_dasar,
			"pasaran_keibesar_dasar":    client.Pasaran_keibesar_dasar,
			"pasaran_keikecil_dasar":    client.Pasaran_keikecil_dasar,
			"pasaran_keigenap_dasar":    client.Pasaran_keigenap_dasar,
			"pasaran_keiganjil_dasar":   client.Pasaran_keiganjil_dasar,
			"pasaran_discbesar_dasar":   client.Pasaran_discbesar_dasar,
			"pasaran_disckecil_dasar":   client.Pasaran_disckecil_dasar,
			"pasaran_discgenap_dasar":   client.Pasaran_discgenap_dasar,
			"pasaran_discganjil_dasar":  client.Pasaran_discganjil_dasar,
		}).
		Post(PATH + "api/savecompanyupdatepasarandasar")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyupdatepasaranshio(c *fiber.Ctx) error {
	type payload_pasaranshio struct {
		Sdata                    string  `json:"sdata" `
		Master                   string  `json:"master" `
		Company                  string  `json:"company" `
		Pasaran_id               string  `json:"pasaran_id" `
		Companypasaran_id        int     `json:"companypasaran_id" `
		Pasaran_shioyear_shio    string  `json:"pasaran_shioyear_shio" `
		Pasaran_minbet_shio      int     `json:"pasaran_minbet_shio" `
		Pasaran_maxbet_shio      int     `json:"pasaran_maxbet_shio" `
		Pasaran_limitglobal_shio int     `json:"pasaran_limitglobal_shio" `
		Pasaran_limittotal_shio  int     `json:"pasaran_limittotal_shio" `
		Pasaran_disc_shio        float32 `json:"pasaran_disc_shio" `
		Pasaran_win_shio         float32 `json:"pasaran_win_shio" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaranshio)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	log.Printf("%d-%d", client.Pasaran_minbet_shio, client.Pasaran_maxbet_shio)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":          hostname,
			"master":                   client.Master,
			"company":                  client.Company,
			"companypasaran_id":        client.Companypasaran_id,
			"pasaran_id":               client.Pasaran_id,
			"pasaran_shioyear_shio":    client.Pasaran_shioyear_shio,
			"pasaran_minbet_shio":      client.Pasaran_minbet_shio,
			"pasaran_maxbet_shio":      client.Pasaran_maxbet_shio,
			"pasaran_limitglobal_shio": client.Pasaran_limitglobal_shio,
			"pasaran_limittotal_shio":  client.Pasaran_limittotal_shio,
			"pasaran_disc_shio":        client.Pasaran_disc_shio,
			"pasaran_win_shio":         client.Pasaran_win_shio,
		}).
		Post(PATH + "api/savecompanyupdatepasaranshio")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasaranlimitline(c *fiber.Ctx) error {
	type payload_fetchpasaranlimitline struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasaranlimitline)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasaranlimitline")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasaran432(c *fiber.Ctx) error {
	type payload_fetchpasaran432 struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasaran432)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasaran432")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasarancbebas(c *fiber.Ctx) error {
	type payload_fetchpasarancbebas struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasarancbebas)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasarancolokbebas")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasarancmacau(c *fiber.Ctx) error {
	type payload_fetchpasarancmacau struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasarancmacau)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasarancolokmacau")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasarancnaga(c *fiber.Ctx) error {
	type payload_fetchpasarancnaga struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasarancnaga)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasarancoloknaga")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasarancjitu(c *fiber.Ctx) error {
	type payload_fetchpasarancjitu struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasarancjitu)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasarancolokjitu")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasaran5050umum(c *fiber.Ctx) error {
	type payload_fetchpasaran5050umum struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasaran5050umum)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasaran5050umum")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasaran5050special(c *fiber.Ctx) error {
	type payload_fetchpasaran5050special struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasaran5050special)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasaran5050special")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasaran5050kombinasi(c *fiber.Ctx) error {
	type payload_fetchpasaran5050kombinasi struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasaran5050kombinasi)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasaran5050kombinasi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasaranmacau(c *fiber.Ctx) error {
	type payload_fetchpasaranmacau struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasaranmacau)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasaranmacaukombinasi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasarandasar(c *fiber.Ctx) error {
	type payload_fetchpasarandasar struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasarandasar)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasarandasar")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyfetchpasaranshio(c *fiber.Ctx) error {
	type payload_fetchpasaranshio struct {
		Sdata             string `json:"sdata" `
		Company           string `json:"company" `
		Companypasaran_id int    `json:"companypasaran_id"`
		Master            string `json:"master" `
		Pasaran_id        string `json:"pasaran_id" `
	}

	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_fetchpasaranshio)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"sdata":             client.Sdata,
			"master":            client.Master,
			"company":           client.Company,
			"companypasaran_id": client.Companypasaran_id,
			"pasaran_id":        client.Pasaran_id,
		}).
		Post(PATH + "api/savecompanyfetchpasaranshio")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
