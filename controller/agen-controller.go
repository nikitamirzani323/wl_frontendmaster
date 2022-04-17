package controller

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type responseagen struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Listcurr interface{} `json:"listcurrency"`
	Record   interface{} `json:"record"`
}

func Agen(c *fiber.Ctx) error {
	type payload_agen struct {
		Master string `json:"master"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_agen)
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
		SetResult(responseagen{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"master":          client.Master,
		}).
		Post(PATH + "api/agen")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*responseagen)
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
func Agensave(c *fiber.Ctx) error {
	type company_save struct {
		Sdata       string `json:"sdata"`
		Page        string `json:"page"`
		Idagen      string `json:"idagen"`
		Idcurr      string `json:"idcurr"`
		Nmagen      string `json:"nmagen"`
		Ownername   string `json:"ownername"`
		Ownerphone  string `json:"ownerphone"`
		Owneremail  string `json:"owneremail"`
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
			"client_hostname":  hostname,
			"page":             client.Page,
			"sdata":            client.Sdata,
			"agen_idagen":      client.Idagen,
			"agen_idcurr":      client.Idcurr,
			"agen_nmagen":      client.Nmagen,
			"agen_owneragen":   client.Ownername,
			"agen_ownerphone":  client.Ownerphone,
			"agen_owneremail":  client.Owneremail,
			"agen_urlendpoint": client.Urlendpoint,
			"agen_status":      client.Status,
		}).
		Post(PATH + "api/saveagen")
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
func Agenlistadmin(c *fiber.Ctx) error {
	type payload_agenadmin struct {
		Idagen string `json:"agen"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_agenadmin)
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
			"idagen":          client.Idagen,
		}).
		Post(PATH + "api/agenlistadmin")
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
func Agensavelistadmin(c *fiber.Ctx) error {
	type company_save struct {
		Sdata    string `json:"sdata"`
		Page     string `json:"page"`
		Idagen   string `json:"agen"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Status   string `json:"status"`
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
			"client_hostname":    hostname,
			"page":               client.Page,
			"sdata":              client.Sdata,
			"agen_idagen":        client.Idagen,
			"agenadmin_username": client.Username,
			"agenadmin_password": client.Password,
			"agenadmin_name":     client.Name,
			"agenadmin_phone":    client.Phone,
			"agenadmin_email":    client.Email,
			"agenadmin_status":   client.Status,
		}).
		Post(PATH + "api/saveagenlistadmin")
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
