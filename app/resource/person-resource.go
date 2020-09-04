package resource

import (
	"net/http"

	"github.com/7dc/gin-structure/app/controller"
	"github.com/7dc/gin-structure/app/repository"
	"github.com/7dc/gin-structure/app/service"
	"github.com/7dc/gin-structure/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db               *mongo.Database             = config.Connect()
	personRepository repository.PersonRepository = repository.New(db)
	personService    service.PersonService       = service.New(personRepository)
	personController controller.PersonController = controller.New(personService)
)

func PersonResource(apiRoutes *gin.RouterGroup) (gin.IRoutes, gin.IRoutes) {
	return apiRoutes.GET("/people", func(ctx *gin.Context) {
			ctx.JSON(200, personController.FindAll())
		}), apiRoutes.POST("/people", func(ctx *gin.Context) {
			err := personController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Person Input is valid!"})
			}
		})

}
