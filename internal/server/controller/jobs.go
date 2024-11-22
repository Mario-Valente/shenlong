package controller

import (
	"net/http"

	"github.com/Mario-valente/shenlong/internal/k8s"
	"github.com/Mario-valente/shenlong/models"
	"github.com/labstack/echo"
)

func GetJob(c echo.Context) error {
	name := c.Param("jobName")
	namespace := c.Param("nsName")

	result, err := k8s.GetJobsK8s(name, namespace, "")
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, result)
}

func CreateJob(c echo.Context) error {

	u := new(models.Job)
	if err := c.Bind(u); err != nil {
		return err
	}

	result, err := k8s.CreateJobsK8s(u.Name, u.Namespace, u.Image, u.Command, u.TTL, "")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, result)
}

func DeleteJob(c echo.Context) error {
	name := c.Param("jobName")
	namespace := c.Param("nsName")

	err := k8s.DeleteJobsK8s(name, namespace, "")
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}
