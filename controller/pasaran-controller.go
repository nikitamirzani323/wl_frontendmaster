package controller

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func Pasaran(c *fiber.Ctx) error {
	type payload_pasaranall struct {
		Master string `json:"master" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaranall)
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
			"master":          client.Master,
		}).
		Post(PATH + "api/allpasaran")
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
func Pasarandetail(c *fiber.Ctx) error {
	type payload_pasaranall struct {
		Pasarancode string `json:"pasarancode" `
		Master      string `json:"master" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaranall)
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
			"master":          client.Master,
			"pasarancode":     client.Pasarancode,
		}).
		Post(PATH + "api/detailpasaran")
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
func Pasarandetailconf(c *fiber.Ctx) error {
	type payload_pasaranall struct {
		Pasarancode string `json:"pasarancode" `
		Master      string `json:"master" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaranall)
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
			"master":          client.Master,
			"pasarancode":     client.Pasarancode,
		}).
		Post(PATH + "api/detailconfpasaran")
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
func Savepasaran(c *fiber.Ctx) error {
	type payload_savepasaran struct {
		Sdata     string `json:"sdata" `
		Master    string `json:"master" `
		Idrecord  string `json:"idrecord" `
		Name      string `json:"pasaran_name" `
		Diundi    string `json:"pasaran_diundi" `
		Url       string `json:"pasaran_url" `
		Jamtutup  string `json:"pasaran_jamtutup" `
		Jamjadwal string `json:"pasaran_jamjadwal"`
		Jamopen   string `json:"pasaran_jamopen"`
		Tipe      string `json:"pasaran_tipe" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_savepasaran)
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
			"idrecord":          client.Idrecord,
			"pasaran_name":      client.Name,
			"pasaran_diundi":    client.Diundi,
			"pasaran_url":       client.Url,
			"pasaran_jamtutup":  client.Jamtutup,
			"pasaran_jamjadwal": client.Jamjadwal,
			"pasaran_jamopen":   client.Jamopen,
			"pasaran_tipe":      client.Tipe,
		}).
		Post(PATH + "api/savepasaran")
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
func Savepasaranlimitline(c *fiber.Ctx) error {
	type payload_savepasaranlimitline struct {
		Master               string `json:"master" `
		Idrecord             string `json:"idrecord" `
		Pasaran_bbfs         int    `json:"pasaran_bbfs" `
		Pasaran_limitline4d  int    `json:"pasaran_limitline4d" `
		Pasaran_limitline3d  int    `json:"pasaran_limitline3d" `
		Pasaran_limitline3dd int    `json:"pasaran_limitline3dd" `
		Pasaran_limitline2d  int    `json:"pasaran_limitline2d" `
		Pasaran_limitline2dd int    `json:"pasaran_limitline2dd" `
		Pasaran_limitline2dt int    `json:"pasaran_limitline2dt" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_savepasaranlimitline)
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
			"idrecord":             client.Idrecord,
			"pasaran_bbfs":         client.Pasaran_bbfs,
			"pasaran_limitline4d":  client.Pasaran_limitline4d,
			"pasaran_limitline3d":  client.Pasaran_limitline3d,
			"pasaran_limitline3dd": client.Pasaran_limitline3dd,
			"pasaran_limitline2d":  client.Pasaran_limitline2d,
			"pasaran_limitline2dd": client.Pasaran_limitline2dd,
			"pasaran_limitline2dt": client.Pasaran_limitline2dt,
		}).
		Post(PATH + "api/savepasaranlimitline")
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
func Savepasaran432(c *fiber.Ctx) error {
	type payload_pasaran432 struct {
		Master                      string  `json:"master" `
		Idrecord                    string  `json:"idrecord" `
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
		Pasaran_limitglobal2dt_432d int     `json:"pasaran_limitglobal2dt_432d" `
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
			"idrecord":                    client.Idrecord,
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
		Post(PATH + "api/savepasaranconf432d")
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
func Savepasarancbebas(c *fiber.Ctx) error {
	type payload_pasarancbebas struct {
		Master                     string  `json:"master" `
		Idrecord                   string  `json:"idrecord" `
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
	client := new(payload_pasarancbebas)
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
			"idrecord":                   client.Idrecord,
			"pasaran_minbet_cbebas":      client.Pasaran_minbet_cbebas,
			"pasaran_maxbet_cbebas":      client.Pasaran_maxbet_cbebas,
			"pasaran_limitotal_cbebas":   client.Pasaran_limitotal_cbebas,
			"pasaran_limitglobal_cbebas": client.Pasaran_limitglobal_cbebas,
			"pasaran_win_cbebas":         client.Pasaran_win_cbebas,
			"pasaran_disc_cbebas":        client.Pasaran_disc_cbebas,
		}).
		Post(PATH + "api/savepasaranconfcolokbebas")
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
func Savepasarancmacau(c *fiber.Ctx) error {
	type payload_pasarancbebas struct {
		Master                     string  `json:"master" `
		Idrecord                   string  `json:"idrecord" `
		Pasaran_minbet_cmacau      int     `json:"pasaran_minbet_cmacau" `
		Pasaran_maxbet_cmacau      int     `json:"pasaran_maxbet_cmacau" `
		Pasaran_limitotal_cmacau   int     `json:"pasaran_limitotal_cmacau"`
		Pasaran_limitglobal_cmacau int     `json:"pasaran_limitglobal_cmacau" `
		Pasaran_win2_cmacau        float32 `json:"pasaran_win2_cmacau" `
		Pasaran_win3_cmacau        float32 `json:"pasaran_win3_cmacau" `
		Pasaran_win4_cmacau        float32 `json:"pasaran_win4_cmacau" `
		Pasaran_disc_cmacau        float32 `json:"pasaran_disc_cmacau" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarancbebas)
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
			"idrecord":                   client.Idrecord,
			"pasaran_minbet_cmacau":      client.Pasaran_minbet_cmacau,
			"pasaran_maxbet_cmacau":      client.Pasaran_maxbet_cmacau,
			"pasaran_limitotal_cmacau":   client.Pasaran_limitotal_cmacau,
			"pasaran_limitglobal_cmacau": client.Pasaran_limitglobal_cmacau,
			"pasaran_win2_cmacau":        client.Pasaran_win2_cmacau,
			"pasaran_win3_cmacau":        client.Pasaran_win3_cmacau,
			"pasaran_win4_cmacau":        client.Pasaran_win4_cmacau,
			"pasaran_disc_cmacau":        client.Pasaran_disc_cmacau,
		}).
		Post(PATH + "api/savepasaranconfcolokmacau")
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
func Savepasarancnaga(c *fiber.Ctx) error {
	type payload_pasarancnaga struct {
		Master                    string  `json:"master" `
		Idrecord                  string  `json:"idrecord" `
		Pasaran_minbet_cnaga      int     `json:"pasaran_minbet_cnaga" `
		Pasaran_maxbet_cnaga      int     `json:"pasaran_maxbet_cnaga" `
		Pasaran_limittotal_cnaga  int     `json:"pasaran_limittotal_cnaga" `
		Pasaran_limitglobal_cnaga int     `json:"pasaran_limitglobal_cnaga"`
		Pasaran_win3_cnaga        float32 `json:"pasaran_win3_cnaga"`
		Pasaran_win4_cnaga        float32 `json:"pasaran_win4_cnaga"`
		Pasaran_disc_cnaga        float32 `json:"pasaran_disc_cnaga"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarancnaga)
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
			"idrecord":                  client.Idrecord,
			"pasaran_minbet_cnaga":      client.Pasaran_minbet_cnaga,
			"pasaran_maxbet_cnaga":      client.Pasaran_maxbet_cnaga,
			"pasaran_limittotal_cnaga":  client.Pasaran_limittotal_cnaga,
			"pasaran_limitglobal_cnaga": client.Pasaran_limitglobal_cnaga,
			"pasaran_win3_cnaga":        client.Pasaran_win3_cnaga,
			"pasaran_win4_cnaga":        client.Pasaran_win4_cnaga,
			"pasaran_disc_cnaga":        client.Pasaran_disc_cnaga,
		}).
		Post(PATH + "api/savepasaranconfcoloknaga")
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
func Savepasarancjitu(c *fiber.Ctx) error {
	type payload_pasarancnaga struct {
		Master                    string  `json:"master" `
		Idrecord                  string  `json:"idrecord" `
		Pasaran_minbet_cjitu      int     `json:"pasaran_minbet_cjitu"`
		Pasaran_maxbet_cjitu      int     `json:"pasaran_maxbet_cjitu"`
		Pasaran_limittotal_cjitu  int     `json:"pasaran_limittotal_cjitu"`
		Pasaran_limitglobal_cjitu int     `json:"pasaran_limitglobal_cjitu"`
		Pasaran_winas_cjitu       float32 `json:"pasaran_winas_cjitu"`
		Pasaran_winkop_cjitu      float32 `json:"pasaran_winkop_cjitu"`
		Pasaran_winkepala_cjitu   float32 `json:"pasaran_winkepala_cjitu"`
		Pasaran_winekor_cjitu     float32 `json:"pasaran_winekor_cjitu"`
		Pasaran_desc_cjitu        float32 `json:"pasaran_desc_cjitu"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasarancnaga)
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
			"idrecord":                  client.Idrecord,
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
		Post(PATH + "api/savepasaranconfcolokjitu")
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
func Savepasaran5050umum(c *fiber.Ctx) error {
	type payload_pasaran5050umum struct {
		Master                       string  `json:"master" `
		Idrecord                     string  `json:"idrecord" `
		Pasaran_minbet_5050umum      int     `json:"pasaran_minbet_5050umum"`
		Pasaran_maxbet_5050umum      int     `json:"pasaran_maxbet_5050umum"`
		Pasaran_limittotal_5050umum  int     `json:"pasaran_limittotal_5050umum" `
		Pasaran_limitglobal_5050umum int     `json:"pasaran_limitglobal_5050umum" `
		Pasaran_keibesar_5050umum    float32 `json:"pasaran_keibesar_5050umum" `
		Pasaran_keikecil_5050umum    float32 `json:"pasaran_keikecil_5050umum" `
		Pasaran_keigenap_5050umum    float32 `json:"pasaran_keigenap_5050umum" `
		Pasaran_keiganjil_5050umum   float32 `json:"pasaran_keiganjil_5050umum" `
		Pasaran_keitengah_5050umum   float32 `json:"pasaran_keitengah_5050umum" `
		Pasaran_keitepi_5050umum     float32 `json:"pasaran_keitepi_5050umum" `
		Pasaran_discbesar_5050umum   float32 `json:"pasaran_discbesar_5050umum"`
		Pasaran_disckecil_5050umum   float32 `json:"pasaran_disckecil_5050umum" `
		Pasaran_discgenap_5050umum   float32 `json:"pasaran_discgenap_5050umum" `
		Pasaran_discganjil_5050umum  float32 `json:"pasaran_discganjil_5050umum"`
		Pasaran_disctengah_5050umum  float32 `json:"pasaran_disctengah_5050umum"`
		Pasaran_disctepi_5050umum    float32 `json:"pasaran_disctepi_5050umum"`
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
			"idrecord":                     client.Idrecord,
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
		Post(PATH + "api/savepasaranconf5050umum")
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
func Savepasaran5050special(c *fiber.Ctx) error {
	type payload_pasaran5050special struct {
		Master                               string  `json:"master" `
		Idrecord                             string  `json:"idrecord" `
		Pasaran_minbet_5050special           int     `json:"pasaran_minbet_5050special"`
		Pasaran_maxbet_5050special           int     `json:"pasaran_maxbet_5050special"`
		Pasaran_limitglobal_5050special      int     `json:"pasaran_limitglobal_5050special"`
		Pasaran_limittotal_5050special       int     `json:"pasaran_limittotal_5050special"`
		Pasaran_keiasganjil_5050special      float32 `json:"pasaran_keiasganjil_5050special"`
		Pasaran_keiasgenap_5050special       float32 `json:"pasaran_keiasgenap_5050special"`
		Pasaran_keiasbesar_5050special       float32 `json:"pasaran_keiasbesar_5050special"`
		Pasaran_keiaskecil_5050special       float32 `json:"pasaran_keiaskecil_5050special"`
		Pasaran_keikopganjil_5050special     float32 `json:"pasaran_keikopganjil_5050special"`
		Pasaran_keikopgenap_5050special      float32 `json:"pasaran_keikopgenap_5050special"`
		Pasaran_keikopbesar_5050special      float32 `json:"pasaran_keikopbesar_5050special"`
		Pasaran_keikopkecil_5050special      float32 `json:"pasaran_keikopkecil_5050special"`
		Pasaran_keikepalaganjil_5050special  float32 `json:"pasaran_keikepalaganjil_5050special"`
		Pasaran_keikepalagenap_5050special   float32 `json:"pasaran_keikepalagenap_5050special"`
		Pasaran_keikepalabesar_5050special   float32 `json:"pasaran_keikepalabesar_5050special"`
		Pasaran_keikepalakecil_5050special   float32 `json:"pasaran_keikepalakecil_5050special"`
		Pasaran_keiekorganjil_5050special    float32 `json:"pasaran_keiekorganjil_5050special"`
		Pasaran_keiekorgenap_5050special     float32 `json:"pasaran_keiekorgenap_5050special"`
		Pasaran_keiekorbesar_5050special     float32 `json:"pasaran_keiekorbesar_5050special"`
		Pasaran_keiekorkecil_5050special     float32 `json:"pasaran_keiekorkecil_5050special"`
		Pasaran_discasganjil_5050special     float32 `json:"pasaran_discasganjil_5050special"`
		Pasaran_discasgenap_5050special      float32 `json:"pasaran_discasgenap_5050special"`
		Pasaran_discasbesar_5050special      float32 `json:"pasaran_discasbesar_5050special"`
		Pasaran_discaskecil_5050special      float32 `json:"pasaran_discaskecil_5050special"`
		Pasaran_disckopganjil_5050special    float32 `json:"pasaran_disckopganjil_5050special"`
		Pasaran_disckopgenap_5050special     float32 `json:"pasaran_disckopgenap_5050special"`
		Pasaran_disckopbesar_5050special     float32 `json:"pasaran_disckopbesar_5050special"`
		Pasaran_disckopkecil_5050special     float32 `json:"pasaran_disckopkecil_5050special"`
		Pasaran_disckepalaganjil_5050special float32 `json:"pasaran_disckepalaganjil_5050special"`
		Pasaran_disckepalagenap_5050special  float32 `json:"pasaran_disckepalagenap_5050special"`
		Pasaran_disckepalabesar_5050special  float32 `json:"pasaran_disckepalabesar_5050special"`
		Pasaran_disckepalakecil_5050special  float32 `json:"pasaran_disckepalakecil_5050special"`
		Pasaran_discekorganjil_5050special   float32 `json:"pasaran_discekorganjil_5050special"`
		Pasaran_discekorgenap_5050special    float32 `json:"pasaran_discekorgenap_5050special"`
		Pasaran_discekorbesar_5050special    float32 `json:"pasaran_discekorbesar_5050special"`
		Pasaran_discekorkecil_5050special    float32 `json:"pasaran_discekorkecil_5050special"`
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
			"idrecord":                             client.Idrecord,
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
		Post(PATH + "api/savepasaranconf5050special")
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
func Savepasaran5050kombinasi(c *fiber.Ctx) error {
	type payload_pasaran5050kombinasi struct {
		Master                                    string  `json:"master" `
		Idrecord                                  string  `json:"idrecord" `
		Pasaran_minbet_5050kombinasi              int     `json:"pasaran_minbet_5050kombinasi"`
		Pasaran_maxbet_5050kombinasi              int     `json:"pasaran_maxbet_5050kombinasi"`
		Pasaran_limitglobal_5050kombinasi         int     `json:"pasaran_limitglobal_5050kombinasi"`
		Pasaran_limittotal_5050kombinasi          int     `json:"pasaran_limittotal_5050kombinasi"`
		Pasaran_belakangkeimono_5050kombinasi     float32 `json:"pasaran_belakangkeimono_5050kombinasi"`
		Pasaran_belakangkeistereo_5050kombinasi   float32 `json:"pasaran_belakangkeistereo_5050kombinasi"`
		Pasaran_belakangkeikembang_5050kombinasi  float32 `json:"pasaran_belakangkeikembang_5050kombinasi"`
		Pasaran_belakangkeikempis_5050kombinasi   float32 `json:"pasaran_belakangkeikempis_5050kombinasi"`
		Pasaran_belakangkeikembar_5050kombinasi   float32 `json:"pasaran_belakangkeikembar_5050kombinasi"`
		Pasaran_tengahkeimono_5050kombinasi       float32 `json:"pasaran_tengahkeimono_5050kombinasi"`
		Pasaran_tengahkeistereo_5050kombinasi     float32 `json:"pasaran_tengahkeistereo_5050kombinasi"`
		Pasaran_tengahkeikembang_5050kombinasi    float32 `json:"pasaran_tengahkeikembang_5050kombinasi"`
		Pasaran_tengahkeikempis_5050kombinasi     float32 `json:"pasaran_tengahkeikempis_5050kombinasi"`
		Pasaran_tengahkeikembar_5050kombinasi     float32 `json:"pasaran_tengahkeikembar_5050kombinasi"`
		Pasaran_depankeimono_5050kombinasi        float32 `json:"pasaran_depankeimono_5050kombinasi"`
		Pasaran_depankeistereo_5050kombinasi      float32 `json:"pasaran_depankeistereo_5050kombinasi"`
		Pasaran_depankeikembang_5050kombinasi     float32 `json:"pasaran_depankeikembang_5050kombinasi"`
		Pasaran_depankeikempis_5050kombinasi      float32 `json:"pasaran_depankeikempis_5050kombinasi"`
		Pasaran_depankeikembar_5050kombinasi      float32 `json:"pasaran_depankeikembar_5050kombinasi"`
		Pasaran_belakangdiscmono_5050kombinasi    float32 `json:"pasaran_belakangdiscmono_5050kombinasi"`
		Pasaran_belakangdiscstereo_5050kombinasi  float32 `json:"pasaran_belakangdiscstereo_5050kombinasi"`
		Pasaran_belakangdisckembang_5050kombinasi float32 `json:"pasaran_belakangdisckembang_5050kombinasi"`
		Pasaran_belakangdisckempis_5050kombinasi  float32 `json:"pasaran_belakangdisckempis_5050kombinasi"`
		Pasaran_belakangdisckembar_5050kombinasi  float32 `json:"pasaran_belakangdisckembar_5050kombinasi"`
		Pasaran_tengahdiscmono_5050kombinasi      float32 `json:"pasaran_tengahdiscmono_5050kombinasi"`
		Pasaran_tengahdiscstereo_5050kombinasi    float32 `json:"pasaran_tengahdiscstereo_5050kombinasi"`
		Pasaran_tengahdisckembang_5050kombinasi   float32 `json:"pasaran_tengahdisckembang_5050kombinasi"`
		Pasaran_tengahdisckempis_5050kombinasi    float32 `json:"pasaran_tengahdisckempis_5050kombinasi"`
		Pasaran_tengahdisckembar_5050kombinasi    float32 `json:"pasaran_tengahdisckembar_5050kombinasi"`
		Pasaran_depandiscmono_5050kombinasi       float32 `json:"pasaran_depandiscmono_5050kombinasi"`
		Pasaran_depandiscstereo_5050kombinasi     float32 `json:"pasaran_depandiscstereo_5050kombinasi"`
		Pasaran_depandisckembang_5050kombinasi    float32 `json:"pasaran_depandisckembang_5050kombinasi"`
		Pasaran_depandisckempis_5050kombinasi     float32 `json:"pasaran_depandisckempis_5050kombinasi"`
		Pasaran_depandisckembar_5050kombinasi     float32 `json:"pasaran_depandisckembar_5050kombinasi"`
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
			"idrecord":                                  client.Idrecord,
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
		Post(PATH + "api/savepasaranconf5050kombinasi")
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
func Savepasaranmacau(c *fiber.Ctx) error {
	type payload_pasaranmacau struct {
		Master                        string  `json:"master" `
		Idrecord                      string  `json:"idrecord" `
		Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi"`
		Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi"`
		Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi"`
		Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi"`
		Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi"`
		Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_pasaranmacau)
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
			"idrecord":                      client.Idrecord,
			"pasaran_minbet_kombinasi":      client.Pasaran_minbet_kombinasi,
			"pasaran_maxbet_kombinasi":      client.Pasaran_maxbet_kombinasi,
			"pasaran_limitglobal_kombinasi": client.Pasaran_limitglobal_kombinasi,
			"pasaran_limittotal_kombinasi":  client.Pasaran_limittotal_kombinasi,
			"pasaran_win_kombinasi":         client.Pasaran_win_kombinasi,
			"pasaran_disc_kombinasi":        client.Pasaran_disc_kombinasi,
		}).
		Post(PATH + "api/savepasaranconfmacaukombinasi")
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
func Savepasarandasar(c *fiber.Ctx) error {
	type payload_pasarandasar struct {
		Master                    string  `json:"master" `
		Idrecord                  string  `json:"idrecord" `
		Pasaran_minbet_dasar      int     `json:"pasaran_minbet_dasar"`
		Pasaran_maxbet_dasar      int     `json:"pasaran_maxbet_dasar"`
		Pasaran_limitglobal_dasar int     `json:"pasaran_limitglobal_dasar"`
		Pasaran_limittotal_dasar  int     `json:"pasaran_limittotal_dasar"`
		Pasaran_keibesar_dasar    float32 `json:"pasaran_keibesar_dasar"`
		Pasaran_keikecil_dasar    float32 `json:"pasaran_keikecil_dasar"`
		Pasaran_keigenap_dasar    float32 `json:"pasaran_keigenap_dasar"`
		Pasaran_keiganjil_dasar   float32 `json:"pasaran_keiganjil_dasar"`
		Pasaran_discbesar_dasar   float32 `json:"pasaran_discbesar_dasar"`
		Pasaran_disckecil_dasar   float32 `json:"pasaran_disckecil_dasar"`
		Pasaran_discgenap_dasar   float32 `json:"pasaran_discgenap_dasar"`
		Pasaran_discganjil_dasar  float32 `json:"pasaran_discganjil_dasar"`
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
			"idrecord":                  client.Idrecord,
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
		Post(PATH + "api/savepasaranconfdasar")
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
func Savepasaranshio(c *fiber.Ctx) error {
	type payload_pasaranshio struct {
		Master                   string  `json:"master" `
		Idrecord                 string  `json:"idrecord" `
		Pasaran_shioyear_shio    string  `json:"pasaran_shioyear_shio"`
		Pasaran_minbet_shio      int     `json:"pasaran_minbet_shio"`
		Pasaran_maxbet_shio      int     `json:"pasaran_maxbet_shio"`
		Pasaran_limitglobal_shio int     `json:"pasaran_limitglobal_shio"`
		Pasaran_limittotal_shio  int     `json:"pasaran_limittotal_shio"`
		Pasaran_disc_shio        float32 `json:"pasaran_disc_shio"`
		Pasaran_win_shio         float32 `json:"pasaran_win_shio"`
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
			"idrecord":                 client.Idrecord,
			"pasaran_shioyear_shio":    client.Pasaran_shioyear_shio,
			"pasaran_minbet_shio":      client.Pasaran_minbet_shio,
			"pasaran_maxbet_shio":      client.Pasaran_maxbet_shio,
			"pasaran_limitglobal_shio": client.Pasaran_limitglobal_shio,
			"pasaran_limittotal_shio":  client.Pasaran_limittotal_shio,
			"pasaran_disc_shio":        client.Pasaran_disc_shio,
			"pasaran_win_shio":         client.Pasaran_win_shio,
		}).
		Post(PATH + "api/savepasaranconfshio")
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
