package controller

import (
	"log"
	"net/http"
	"time"

	"bitbucket.org/isbtotogroup/japan-pools/config"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

var PATH string = config.GetApiPath()

// type clientlistresult struct {
// 	Token string `json:"token"`
// }

type response struct {
	Status int         `json:"status"`
	Record interface{} `json:"record"`
}

func ListResultHome(c *fiber.Ctx) error {

	render_page := time.Now()

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": "",
		}).
		Post(PATH + "api/resultjapanday")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
func ListResultHomeNight(c *fiber.Ctx) error {
	render_page := time.Now()

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": "",
		}).
		Post(PATH + "api/resultjapannight")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
