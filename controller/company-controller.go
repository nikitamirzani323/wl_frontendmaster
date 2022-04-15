package controller

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type responsecompany struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Listcurr interface{} `json:"listcurrency"`
	Record   interface{} `json:"record"`
}

func Company(c *fiber.Ctx) error {
	type payload_company struct {
		Master string `json:"master"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_company)
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
		SetResult(responsecompany{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"master":          client.Master,
		}).
		Post(PATH + "api/company")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*responsecompany)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"listcurrency": result.Listcurr,
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
func Companysave(c *fiber.Ctx) error {
	type company_save struct {
		Sdata       string `json:"sdata"`
		Page        string `json:"page"`
		Idcomp      string `json:"idcomp"`
		Idcurr      string `json:"idcurr"`
		Nmcompany   string `json:"nmcompany"`
		Nmowner     string `json:"nmowner"`
		Phoneowner  string `json:"phoneowner"`
		Emailowner  string `json:"emailowner"`
		Urlendpoint string `json:"urlendpoint"`
		Status      string `json:"status"`
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
			"client_hostname":     hostname,
			"page":                client.Page,
			"sdata":               client.Sdata,
			"company_idcompany":   client.Idcomp,
			"company_idcurr":      client.Idcurr,
			"company_nmcompany":   client.Nmcompany,
			"company_nmowner":     client.Nmowner,
			"company_phoneowner":  client.Phoneowner,
			"company_emailowner":  client.Emailowner,
			"company_urlendpoint": client.Urlendpoint,
			"company_status":      client.Status,
		}).
		Post(PATH + "api/savecompany")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsedefault)
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
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
	type payload_company struct {
		Idcompany string `json:"company"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_company)
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
			"idcompany":       client.Idcompany,
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
