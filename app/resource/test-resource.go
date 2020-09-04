package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestResource(apiRoutes *gin.RouterGroup) (gin.IRoutes, gin.IRoutes) {
	return apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, personController.FindAll())
		}), apiRoutes.POST("/test", func(ctx *gin.Context) {
			err := personController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Person Input is valid!"})
			}
		})

}
