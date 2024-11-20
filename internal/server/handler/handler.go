package handler

import (
	"github.com/Mario-valente/shenlong/internal/server/controller"
	"github.com/labstack/echo"
)

func Server() {
	e := echo.New()
	e.GET("/jobs/:jobName/namespace/:nsName", controller.GetJob)
	e.POST("/jobs/", controller.CreateJob)
	e.GET("/crons/:cronName/namespace/:nsName", controller.GetCron)
	e.POST("/crons/", controller.CreateCron)
	e.Logger.Fatal(e.Start(":3001"))

}
