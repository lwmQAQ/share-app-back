package resourcehandle

import (
	"net/http"
	"resource-server/internal/constants"
	"resource-server/internal/ecode"
	"resource-server/internal/models"
	"resource-server/internal/svc"
	"resource-server/internal/types"

	"github.com/gin-gonic/gin"
)

var contents = []string{"title", "tags"}

func SearchByEsHandle(c *gin.Context, svc *svc.ServiceContext) {
	input := c.Query("input")
	var resp []models.Resource
	err := svc.ESClient.MultiMatchSearch(contents, input, constants.ResourceIndex, &resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error(ecode.ErrSystemError))
	}
	c.JSON(http.StatusOK, types.Success(resp))
}
