package handler

import "github.com/labstack/echo/v4"

func GetBody[T any](c echo.Context, s T) (*T, error) {
	req := new(T)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}
