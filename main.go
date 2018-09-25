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
	member := &models.Member{}
	db.PGMaster.CreateTable(member)

}
