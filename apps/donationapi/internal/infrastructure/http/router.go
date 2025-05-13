package http

import (
	"donationapi/internal/infrastructure/http/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() http.Handler {
	router := gin.Default()
	handler.PahpaRouterGroup(router)
	handler.FileRouter(router)

	return router
}
