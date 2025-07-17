package transport

import (
	"main/handler"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(ExtraMiddleware())

	controlador:= handler.NuevoHandle(db)

	api:= router.Group("/api")
	api.GET("/health", controlador.Ping)
	// api.GET("/licencia/:id", controlador.Ping)
	return router

}

func ExtraMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		 methods := []string{"GET", "POST", "DELETE"}
		c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

		if !contains(methods, c.Request.Method) {
			c.AbortWithStatus(403) 
			return
		}

		c.Next()
	}
}


func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}