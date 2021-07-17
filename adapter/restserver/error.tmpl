package restserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/logger"
)

type Error struct {
	Code    interface{} `json:"code,omitempty"`
	Message string      `json:"message"`
}

func errorHandler(log *logger.Logger) func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		if code == http.StatusNotFound {
			log.Errorf("Route [%s] not mapped", c.Path())
		} else {
			log.Errorf("Echo return an error with status code [%d] for the path [%s]", code, c.Path())
		}

		if err = c.JSON(code, Error{code, err.Error()}); err != nil {
			log.Errorf("Error creating response - %+v", err)
		}
	}
}
