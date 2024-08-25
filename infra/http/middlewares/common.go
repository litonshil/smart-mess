package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"smart-mess/utils"
)

// BindBody binds request body contents to bindable object
func BindBody(c echo.Context, i interface{}) error {
	// read origin body bytes
	var bodyBytes []byte
	if !utils.IsEmpty(c.Request().Body) {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		// write back to request body
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// parse json data
		err := json.Unmarshal(bodyBytes, &i)
		if err != nil {
			return err
		}
	}
	return nil
}
