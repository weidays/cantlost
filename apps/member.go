package apps

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"projects/cantlost/models"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp"
)

const (
	appID  = "wx63abe446f9e802d7"
	secret = "bd90beee6bf53fa77a59b8b7c7fd12f4"
)

//MemberLogin 会员登录
func MemberLogin(c *gin.Context) {
	code := c.Param("code")
	// @appID 小程序 appID
	// @secret 小程序的 app secret
	// @code 小程序登录时获取的 code
	res, err := weapp.Login(appID, secret, code)
	if err != nil {
		fmt.Println("调用微信登录失败:", err)
	}

	// res.OpenID
	// res.SessionKey
	// res.UnionID
	fmt.Printf("返回结果: %#v", res)

	//如果数据库中还没有这个账号，则加入

}

//获取列表
func MemberList(c *gin.Context) {
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

//获取一个会员
func MemberGet(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))

	mem, err := models.OneMember(int64(mid))
	fmt.Println(mem, err)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		//log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    mem,
		})
	}

}

//修改会员信息
func MemberEdit(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))
	m := new(models.Member)
	m.ID = mid
	m.LoginName = c.Request.FormValue("login_name")
	m.Password = c.Request.FormValue("password")

	if n, err := m.UpdateMember(mid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}

//新增
func MemberAdd(c *gin.Context) {
	m := new(models.Member)
	m.LoginName = c.Request.FormValue("login_name")
	m.Password = c.Request.FormValue("password")

	if id, err := m.AddMember(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		m.ID = int(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    m,
		})
	}
}

//删除
func MemberDelete(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))
	if n, err := models.DeleteMember(mid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}
