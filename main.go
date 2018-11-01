package main

import (
	db "projects/cantlost/dbs"
	"projects/cantlost/libs"
	"projects/cantlost/models"
	"projects/cantlost/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	defer db.PGMaster.Close()
	initTabel()
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	router.Run(":" + libs.Conf.Read("site", "httpport"))
}

func initTabel() {
	db.PGMaster.AutoMigrate(&models.Member{})
	db.PGMaster.AutoMigrate(&models.LostInfo{})
	db.PGMaster.AutoMigrate(&models.Category{})
	db.PGMaster.AutoMigrate(&models.Target{})
	db.PGMaster.AutoMigrate(&models.TargetOtherInfo{})
	db.PGMaster.AutoMigrate(&models.Area{})
}
