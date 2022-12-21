package watchlist

import (
	"github.com/gin-gonic/gin"
	"github.com/harisaginting/krdv/common/http/response"
)

type Controller struct {
	service Service
}

func ProviderController(s Service) Controller {
	return Controller{
		service: s,
	}
}

func (ctrl *Controller) List(c *gin.Context) {
	ctx := c.Request.Context()

	var resData ResponseList
	ctrl.service.List(ctx, &resData)

	// return
	response.Json(c, resData)
	return
}
