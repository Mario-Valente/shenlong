package controller

import (
	"github.com/Mario-valente/shenlong/internal/k8s"
	"github.com/Mario-valente/shenlong/models"
	"github.com/labstack/echo"
)

func GetJob(c echo.Context) error {
	name := c.Param("jobName")
	namespace := c.Param("nsName")

	k8s.GetJobsK8s(name, namespace, "")
	return c.String(200, "ok")
}

func CreateJob(c echo.Context) error {

	u := new(models.Job)
	if err := c.Bind(u); err != nil {
		return err
	}

	k8s.CreateJobsK8s(u.Name, u.Namespace, u.Image, u.Command, u.TTL, "")
	return c.String(200, "ok")
}
