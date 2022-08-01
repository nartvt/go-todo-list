package middleware

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowedMethods:  []string{"GET", "POST", "PUT", "PATH", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	})
}
