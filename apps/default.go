/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-10-13 11:01:27
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-18 10:10:49
***********************************************/
package apps

import (
	"net/http"
	"projects/cantlost/libs"
	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.HTML(http.StatusOK, "default.html", gin.H{
		"title": libs.Conf.Read("site", "appname"),
	})

}
