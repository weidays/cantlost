package routers

import (
	"net/http"
	"projects/cantlost/apps"
	"projects/cantlost/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	// router.Static("/static", "/static")
	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLGlob("templates/*")
	router.GET("/", apps.IndexApi)
	v0 := router.Group("/v0")
	v0.Use(middlewares.Auth())
	{

		/**
		* @api GET /member
		* @apiExample json
		* { "id":1, "name":"admin" }
		 */
		v0.POST("/member", apps.MemberAdd)
		/**
		 * @api GET /users/:id 获取指定用户的相关信息
		 * @apiGroup users
		 * @apiParam id int 表示用户 id 的唯一值
		 *
		 * @apiSuccess 200 json ok
		 * @apiExample
		 * {"id":1, "name": "n1"}
		 */
		// curl -X GET http://127.0.0.1:8000/v0/member
		v0.GET("/member", apps.MemberList)
		// curl -X GET http://127.0.0.1:8000/v0/member/1
		v0.GET("/member/:id", apps.MemberGet)
		//curl -X PUT http://127.0.0.1:8000/v0/member/1 -d "login_name=haodaquan&password=1234"
		v0.PUT("/member/:id", apps.MemberEdit)
		// curl -X DELETE http://127.0.0.1:8000/v0/member/2
		v0.DELETE("/member/:id", apps.MemberDelete)
	}

	return router
}
