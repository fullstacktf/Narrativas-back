package router

import (
	"github.com/fullstacktf/Narrativas-Backend/api/controllers"
	mw "github.com/fullstacktf/Narrativas-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/characters/", mw.IsSignedIn, controllers.GetCharacters)
	router.GET("/characters/:id", mw.IsSignedIn, controllers.GetCharacter)
	router.POST("/characters/", mw.IsSignedIn, controllers.PostCharacter)
	router.DELETE("/characters/:id", mw.IsSignedIn, controllers.DeleteCharacter)
	router.PUT("/characters/", mw.IsSignedIn, controllers.PutCharacter)
	router.POST("/characters/:id/sections", mw.IsSignedIn, controllers.PostSection)

	router.GET("/stories/", controllers.Get)
	router.GET("/stories/:id", controllers.GetStory)
	router.POST("/stories/", controllers.PostStory)
	router.POST("/stories/:id/events/", controllers.PostEvent)
	router.POST("/stories/:id/events/relations", controllers.PostEventRelation)
	router.DELETE("/stories/:id", controllers.DeleteStory)

	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)

	router.POST("/upload/images/character", mw.IsSignedIn, controllers.UploadCharacter)
	router.POST("/upload/images/story", mw.IsSignedIn, controllers.UploadStory)

	router.Static("/static", "./images")

	return router
}
