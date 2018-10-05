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
			* @api POST http://aip.fixme.vip/v0/lostinfo
			* @apiGroup lostinfo
			* @apiParam page int 第几页
			* @apiParam page_size int 每页条数
			* @apiParam category_id  int 分类id
			* @apiExample 请求示例
			* {
			  "page": 1,
			  "page_size": 15,
			  "category_id": 18
			}
			* @apiExample 返回数据示例
			* {
	    "count": 1,
	    "current:": 1,
	    "data": [
	        {
	            "id": 1,
	            "publish_at": "2018-09-30T08:14:56.118+07:00",
	            "updated_at": "2018-09-30T08:14:58.463+07:00",
	            "member_id": 1,
	            "member_info": {
	                "id": 1,
	                "registed_at": "2018-09-30T08:09:57.018+07:00",
	                "updated_at": "2018-09-30T08:10:01.942+07:00",
	                "pic_url": "https://p1.4499.cn/touxiang/UploadPic/2013-4/27/201304271944118052.jpg",
	                "nick_name": "老妖怪",
	                "real_name": "l老妖",
	                "phone_number": "13808998988",
	                "gender": "2",
	                "birthDay": "2018-09-30T08:14:34.085+07:00",
	                "login_name": "weidays",
	                "password": "goodboy"
	            },
	            "category_id": 1,
	            "category_info": null,
	            "title": "非常牛逼的一个故事",
	            "cover_pics": "",
	            "lost_time": null,
	            "lost_place_str": "",
	            "lost_place_lon": "",
	            "lost_place_lat": "",
	            "target_id": 0,
	            "target_info": null,
	            "lost_event": "",
	            "target_stroy": "",
	            "the_end": "",
	            "like_num": 199,
	            "dislike_num": 299,
	            "comment_num": 99,
	            "share_num": 98
	        }
	    ],
	    "message": "SUCCESS",
	    "page_size": 15,
	    "status": 200
	}
	*/
	v0.POST("/lostinfo", apps.LostInfoList)

	v0.POST("/publish_lostinfo", apps.LostInfoAdd)

	return router
}
