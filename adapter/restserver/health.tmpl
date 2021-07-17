package restserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type JSON map[string]interface{}

type HeartbeatResponse struct {
	Stage    string `json:"stage"`
	Version  string `json:"app_version"`
}

func response(c echo.Context, status int, message interface{}) error {
	if err, ok := message.(error); ok {
		if err = c.JSON(status, Error{status, err.Error()}); err != nil {
			return err
		}
		return nil
	}

	if err := c.JSON(status, message); err != nil {
		return err
	}
	return nil

}

func (s *Server) health(c echo.Context) error {
	if err := response(c, http.StatusOK, HeartbeatResponse{
		Stage:    s.conf.Env,
		Version:  s.conf.Version,
	}); err != nil {
		return err
	}
	return nil
}
