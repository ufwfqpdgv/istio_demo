package routers

import (
	"samh_ticket_rank_0.1/server/apis"
	"samh_ticket_rank_0.1/server/custom_middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(custom_middleware.ServerHeader)
	router.Use(custom_middleware.Log())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/ticket/v1", apis.GetCidTicket)
	return router
}
