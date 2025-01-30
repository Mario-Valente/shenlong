package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Mario-valente/shenlong/internal/server/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type LogEntry struct {
	Time       string `json:"time"`
	Level      string `json:"level"`
	Message    string `json:"message"`
	Method     string `json:"method"`
	URI        string `json:"uri"`
	StatusCode int    `json:"status_code"`
}

func LogJSON(method, uri, time, level, message string, statusCode int) {

	logEntry := LogEntry{
		Time:       time,
		Level:      level,
		Message:    message,
		Method:     method,
		URI:        uri,
		StatusCode: statusCode,
	}

	logData, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(logData))
}

func Server() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${status} ${method} ${uri} ${error}` + "\n",
		Output: os.Stdout,
	}))
	e.GET("/jobs/:jobName/namespace/:nsName", controller.GetJob)

	e.POST("/jobs/", controller.CreateJob)
	e.GET("/crons/:cronName/namespace/:nsName", controller.GetCron)
	e.POST("/crons/", controller.CreateCron)
	e.DELETE("/jobs/:jobName/namespace/:nsName", controller.DeleteJob)
	e.DELETE("/crons/:cronName/namespace/:nsName", controller.DeleteCron)
	e.GET("/logs/:typeLog", controller.LogsOutputs)
	e.GET("/latency/", controller.LatencyOutputs)
	e.GET("/error/:statusCode", controller.ErrorSimulateCount)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.Logger.Fatal(e.Start(":3001"))

}
