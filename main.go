package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sip/simru/database"
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDBConnection("")
)

func main() {
	defer database.CloseDatabaseConnection(db)
	db.AutoMigrate(&entity.UsersPAK3{}, &entity.UsersP2AK3{})
	g := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Test init project",
		})
	})
	g.Run(os.Getenv("SERVER_PORT"))
}
