package controller

import (
	"net/http"

	"github.com/Mario-valente/shenlong/internal/k8s"
	"github.com/Mario-valente/shenlong/models"
	"github.com/labstack/echo"
)

func GetCron(c echo.Context) error {
	name := c.Param("cronName")
	namespace := c.Param("nsName")

	result, err := k8s.GetCronsK8s(name, namespace, "")
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, result)
}

func CreateCron(c echo.Context) error {

	u := new(models.Cron)
	if err := c.Bind(u); err != nil {
		return err
	}

	result, err := k8s.CreateCronsK8s(u.Name, u.Namespace, u.Image, u.Command, "", u.TTL, u.Schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, result)
}
