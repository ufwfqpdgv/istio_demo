package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"samh_ticket_rank_0.1/recommend/go/collection_api"
	"samh_ticket_rank_0.1/recommend/go/define"
	"github.com/zhanglanhui/go-utils/utils"
)

func getParam(c *gin.Context) string {
	comicId := c.Query("comicId")
	if "" == comicId {
		comicId = "0"
	}
	return comicId
}

// 用来检测返回数据
func CheckJsonCode(jsonData define.JsonOut) {
	if jsonData.Status != "0" {
		panic("json code != 0")
	}
}

func GetCidTicket(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			c.Header("Content-Type", "text/json; charset=utf-8")
			c.String(http.StatusInternalServerError, "[]")
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			utils.CheckErrSendEmail(fmt.Errorf("recovery:%s\nstack:%s", rec, string(buf[:n])))
		}
	}()
	comicId := getParam(c)
	jsonData := collection_api.CidTicketApi(comicId) // 把结果封装到struct里
	c.JSON(http.StatusOK, jsonData)
	CheckJsonCode(jsonData)
}
