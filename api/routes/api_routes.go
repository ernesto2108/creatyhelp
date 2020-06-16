package api

import (
	app "github.com/ernesto2108/AP_CreatyHelp/pkg/user/application/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllRoutes(handler *app.CreateUsersHandler) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// users

	// Global middleware
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"Message" : "Hello Api!!"})
		})

		v1.POST("/user", handler.SaveUsersEndPoint)
		v1.GET("/user/{id}", handler.GetIdUsersEndPoint)
	}

	return r

}
