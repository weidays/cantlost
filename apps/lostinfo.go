package apps

import (
	"log"
	"net/http"
	"projects/cantlost/libs"

	"projects/cantlost/models"
	"projects/cantlost/models/vo"

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
	listvo := []vo.LostInfoVo{}
	for _, info := range list {
		vo := vo.LostInfoVo{}
		libs.CopyStruct(info, vo)
		member, err := models.OneMember(info.UserID)
		if err != nil {

		}
		vo.MemberInfo = &member
		listvo = append(listvo, vo)
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
			"data":      listvo,
			"count":     n,
			"page_size": requestInfo.PageSize,
			"current:":  requestInfo.Page,
		})
	}
}
