package apps

import (
	"log"
	"net/http"
	"strconv"

	"projects/cantlost/models"

	"github.com/gin-gonic/gin"
)

//LostInfoList 获取列表
func LostInfoList(c *gin.Context) {
	filters := make([]interface{}, 0)
	filters = append(filters, "id", "<>", "0")

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("page_size"))

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	list, n, err := models.ListMember(page, pageSize, filters...)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "123",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "SUCCESS",
			"data":      list,
			"count":     n,
			"page_size": pageSize,
			"current:":  page,
		})
	}
}
