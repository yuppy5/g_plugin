package g_plugin

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yuppy5/g_healthcheck"
	"github.com/yuppy5/g_log"
)

var onceDo sync.Once

func DemoPlugin(r *gin.Engine) {
	onceDo.Do(func() {
		r.Use(g_log.LoggerMiddleware)

		r.GET("/is_running", g_healthcheck.Process)
		r.HEAD("/is_running", g_healthcheck.Process)
		r.GET("/healthcheck.html", g_healthcheck.HealthCheck)
		r.HEAD("/healthcheck.html", g_healthcheck.HealthCheck)
	})
}
