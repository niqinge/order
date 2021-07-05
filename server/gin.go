package server

import (
	"github.com/gin-gonic/gin"
	"github.com/niqinge/order/server/middleware"
)

func NewEngine(h *Server) *gin.Engine {
	e := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r := e.Group("/api/v1/")
	r.Use(middleware.Logger())
	{
		r.GET("order/:order_no", h.QueryOrder())
		r.POST("order", h.AddOrder())
		r.PUT("order", h.UpdateOrder())
		r.GET("order/list", h.QueryOrderList())
		r.POST("upload", h.Upload())
	}

	return e
}
