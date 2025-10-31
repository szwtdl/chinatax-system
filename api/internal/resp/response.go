package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"net/http"
)

func SUCCESS(c *gin.Context, data interface{}, code ...types.BizCode) {
	respCode := types.Success
	if len(code) > 0 {
		respCode = code[0]
	}
	c.JSON(http.StatusOK, types.BizVo{
		Code: respCode,
		Data: data,
	})
}

func ERROR(c *gin.Context, message string, code ...types.BizCode) {
	respCode := types.Failed
	if len(code) > 0 {
		respCode = code[0]
	}
	c.JSON(http.StatusOK, types.BizVo{
		Code:    respCode,
		Message: message,
	})
}

func HACKER(c *gin.Context) {
	c.JSON(http.StatusOK, types.BizVo{Code: types.Failed, Message: "Hacker attempt!!!"})
}

func NotAuth(c *gin.Context) {
	c.JSON(http.StatusOK, types.BizVo{Code: types.NotAuthorized, Message: "Not Authorized"})
}

func NotAuthorize(c *gin.Context) {
	c.JSON(http.StatusOK, types.BizVo{Code: types.NotAuthorized, Message: "您的授权已过期，请授权后再试"})
}
