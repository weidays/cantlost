package routers

import (
	"net/http"
	"projects/cantlost/apps"
	"projects/cantlost/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.Static("/apidoc", "/apidoc")
	// router.Static("/static", "/static")
	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLGlob("templates/*")
	router.GET("/", apps.IndexApi)
	v0 := router.Group("/v0")
	v0.Use(middlewares.Auth())
	{

		v0.POST("/member", apps.MemberAdd)

		// curl -X GET http://127.0.0.1:8000/v0/member
		v0.GET("/member", apps.MemberList)
		// curl -X GET http://127.0.0.1:8000/v0/member/1
		v0.GET("/member/:id", apps.MemberGet)
		//curl -X PUT http://127.0.0.1:8000/v0/member/1 -d "login_name=haodaquan&password=1234"
		v0.PUT("/member/:id", apps.MemberEdit)
		// curl -X DELETE http://127.0.0.1:8000/v0/member/2
		v0.DELETE("/member/:id", apps.MemberDelete)

		/**
		* @api GET /lostinfo
		* @apiGroup lostinfo
		* @apiParam page int 第几页
		* @apiParam page_size int 每页条数
		* @apiParam category_id  分类id
		* @apiExample json
		* {"status":  200,
			"message":   "SUCCESS",
			"data":      [],
			"count":     100,
			"page_size": 20,
			"current:":  2}
		*/
		v0.POST("/lostinfo", apps.LostInfoList)

	}

	return router
}
