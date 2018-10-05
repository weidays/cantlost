package apps

import (
	"log"
	"net/http"
	"projects/cantlost/libs"

	"projects/cantlost/models"

	"github.com/gin-gonic/gin"
)

type RequestInfo struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	CategoryID int64 `json:"category_id"`
}

//LostInfoList 获取列表
func LostInfoList(c *gin.Context) {
	filters := make([]interface{}, 0)
	filters = append(filters, "id", "<>", "0")
	requestInfo := &RequestInfo{}
	c.BindJSON(requestInfo)

	if requestInfo.Page == 0 {
		requestInfo.Page = 1
	}

	if requestInfo.PageSize == 0 {
		requestInfo.PageSize = 10
	}
	list, n, err := models.ListLostInfo(requestInfo.Page, requestInfo.PageSize, filters...)
	for index, info := range list {
		member, err := models.OneMember(info.MemberID)
		if err != nil {

		}
		list[index].MemberInfo = &member
	}

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
			"page_size": requestInfo.PageSize,
			"current:":  requestInfo.Page,
		})
	}
}

//新增
func LostInfoAdd(c *gin.Context) {
	form := &models.LostInfoForm{}
	c.BindJSON(form)

	lostInfo := models.LostInfo{}
	libs.CopyStruct(form, lostInfo)

	if id, err := lostInfo.AddLostInfo(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		lostInfo.ID = id
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    lostInfo,
		})
	}
}
