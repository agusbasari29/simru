package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sip/simru/database"
	"github.com/sip/simru/database/seeders"
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDBConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)
	db.AutoMigrate(&entity.Sections{}, &entity.UserRoles{}, &entity.Persons{}, &entity.Users{})
	g := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Test init project",
		})
	})
	g.GET("/seeder", func(ctx *gin.Context) {
		seeders.Seeders()
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully seed database"})
	})
	g.Run(os.Getenv("SERVER_PORT"))
}
