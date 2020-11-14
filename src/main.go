package main

import (
	controllers "github.com/fullstacktf/Narrativas-Backend/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/character/:id", controllers.GetCharacter)
	r.POST("/character/", controllers.PostCharacter)
	r.DELETE("/character/:id", controllers.DeleteCharacter)
	r.PATCH("/character/:id", controllers.PatchCharacter)

	r.GET("/story/:id", controllers.GetStory)
	r.DELETE("/story/:id", controllers.DeleteStory)

	r.Run(":10000")
}
