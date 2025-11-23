package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/Mario-valente/shenlong/internal/server/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
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

func initTracer() (*sdktrace.TracerProvider, error) {

	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("shenlong-server"),
		),
	)

	exporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint("otel-collector-collector:4318"))

	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func Server() {

	tp, err := initTracer()
	if err != nil {
		fmt.Println(err)
	}
	defer tp.Shutdown(context.Background())

	e := echo.New()

	e.Use(otelecho.Middleware("shenlong-server", otelecho.WithTracerProvider(tp)))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${status} ${method} ${uri} ${error}` + "\n",
		Output: os.Stdout,
	}))

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "healthy", "service": "shenlong-server"})
	})

	e.GET("/jobs/:jobName/namespace/:nsName", controller.GetJob)
	e.POST("/jobs/", controller.CreateJob)
	e.GET("/crons/:cronName/namespace/:nsName", controller.GetCron)
	e.POST("/crons/", controller.CreateCron)
	e.DELETE("/jobs/:jobName/namespace/:nsName", controller.DeleteJob)
	e.DELETE("/crons/:cronName/namespace/:nsName", controller.DeleteCron)

	// Custom routes for testing logs
	e.GET("/logs/:typeLog", controller.LogsOutputs)

	// Custom routes for testing metrics
	e.GET("/latency/", controller.LatencyOutputs)
	e.GET("/error/:statusCode", controller.ErrorSimulateCount)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Custom route for testing traces
	e.GET("/trace", controller.TraceOutputs)

	e.Logger.Fatal(e.Start(":3001"))

}
