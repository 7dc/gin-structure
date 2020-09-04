package app

import (
	"github.com/7dc/gin-proper-structure/app/resource"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	server := gin.Default()
	apiRoutes := server.Group("/api")
	{
		resource.PersonResource(apiRoutes)
		resource.TestResource(apiRoutes)
	}

	server.Run(":8000")
}
