package linkhandle

import (
	"net/http"
	"resource-server/internal/svc"

	"github.com/gin-gonic/gin"
)

func RedirectHandle(c *gin.Context, svc *svc.ServiceContext) {
	//获取唯一code找到对应源url进行重定向
	code := c.Param("code")
	sourceurl, err := svc.UrlUtil.GetSourceUrl(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	//成功就重定向
	c.Redirect(http.StatusFound, sourceurl)
}
