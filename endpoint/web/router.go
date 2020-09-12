package web

import (
	"net/http"

	"github.com/farwydi/shorturl/domain"
	"github.com/gin-gonic/gin"
)

func NewGinEngineForTesting() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewRouter(r *gin.Engine, data domain.DataGateway) http.Handler {
	r.GET("/", func(c *gin.Context) {

	})

	r.GET("/:prefix", func(c *gin.Context) {
		ctx := c.Request.Context()
		next, err := data.FindUrl(ctx, c.Param("prefix"))
		if err != nil {
			return
		}

		c.Redirect(http.StatusFound, next.String())
	})

	return r
}
