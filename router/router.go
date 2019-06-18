package router

import (
	"apiserver/handler/product"
	"apiserver/handler/story"
	"apiserver/handler/wechat"
	"net/http"

	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	p := g.Group("/v1/product")
	{
		p.GET("", product.List)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	s := g.Group("/story")
	{
		s.GET("", story.Get)
	}
	w := g.Group("/procSignature")
	{
		w.GET("", wechat.Get)
		w.POST("", wechat.Post)
	}

	h := g.Group("/")
	{
		h.GET("/main", func(context *gin.Context) {
			context.HTML(http.StatusOK, "main.tmpl", gin.H{
				"title": "Main website",
			})
		})
		h.GET("/nav", func(context *gin.Context) {
			context.HTML(http.StatusOK, "nav.html", gin.H{})
		})
		h.GET("/user", func(context *gin.Context) {
			context.HTML(http.StatusOK, "user.tmpl", gin.H{})
		})
		h.GET("/product", func(context *gin.Context) {
			context.HTML(http.StatusOK, "product.tmpl", gin.H{})
		})
	}

	return g
}
