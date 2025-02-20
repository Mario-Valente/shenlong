package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func TraceOutputs(c echo.Context) error {

	tracer := otel.Tracer("shenlong-trace")
	ctx, span := tracer.Start(c.Request().Context(), "TraceOutputs")
	defer span.End()

	urls := []string{
		"http://localhost:3001/latency/",
		"http://localhost:3001/error/200",
		"http://localhost:3001/latency/",
	}

	for _, url := range urls {
		err := callExternalService(ctx, url)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to call external service")
			c.String(http.StatusInternalServerError, "failed to call external service")
			return err
		}
	}

	return c.String(http.StatusOK, "All services called successfully")

}

func callExternalService(ctx context.Context, url string) error {
	tracer := otel.Tracer("shenlong-trace")
	_, span := tracer.Start(ctx, fmt.Sprintf("Calling %s", url))
	defer span.End()

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to create request")
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "request failed")
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	span.SetAttributes(
		attribute.Int("http.status_code", resp.StatusCode),
		attribute.String("response", string(body)),
	)

	return nil
}
