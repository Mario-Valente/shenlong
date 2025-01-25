package controller

import "github.com/labstack/echo/v4"

func LogsOutputs(c echo.Context) error {
	typeLog := c.Param("typeLog")

	switch typeLog {
	case "error":
		c.Logger().Error("This is a simulated error log")
		return c.String(500, "Simulated error log")
	case "info":
		c.Logger().Info("This is a simulated info log")
		return c.String(200, "Simulated info log")
	case "warn":
		c.Logger().Warn("This is a simulated warn log")
		return c.String(200, "Simulated warn log")
	case "debug":
		c.Logger().Debug("This is a simulated debug log")
		return c.String(200, "Simulated debug log")
	case "critical":
		c.Logger().Error("This is a simulated critical log")
		return c.String(500, "Simulated critical log")
	default:
		return c.String(400, "Invalid log type")

	}
}
