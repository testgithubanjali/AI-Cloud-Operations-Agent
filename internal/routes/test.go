package routes

import (
	"ai-sre-agent/internal/api"

	"github.com/gin-gonic/gin"
)

func TestRoutes(r *gin.Engine) {
	r.GET("/test", api.TestConnection)

	r.GET("/test/pods", api.GetPods)
	r.GET("/test/pods/:name", api.DescribePod)
	r.GET("/test/logs/:name", api.GetLogs)
	r.GET("/test/events", api.GetEvents)
	r.GET("/test/metrics", api.GetMetrics)
	r.GET("/test/planner", api.TestPlanner)
	r.GET("/test/executor", api.TestExecutor)
	r.GET("/test/analyzer", api.TestAnalyzer)
	r.GET("/prometheus", api.PrometheusTest)
}
