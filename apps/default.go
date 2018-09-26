package apps

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "欢迎来到Cantlost的世界，现在时间:%s", time.Now().Format("2006-01-02 15:04:05"))
	// c.HTML(http.StatusOK, "default.html", gin.H{
	// 	"title": libs.Conf.Read("site", "appname"),
	// })

}
