package handler

import (
	"net/http"
	"project/models"

	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
}

func InitNews() NewsHandler {
	return NewsHandler{}
}

func (h NewsHandler) FetchNews(c echo.Context) (err error) {
	datas := []models.News{
		models.News{
			ID:    "1",
			Title: "Hello World!",
			Body:  "No! Hi World!",
		},
	}

	return c.JSON(http.StatusOK, datas)
}