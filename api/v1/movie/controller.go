package movie

import (
	"github.com/gin-gonic/gin"
	"github.com/harisaginting/krdv/common/http/response"
	"github.com/harisaginting/krdv/common/utils/helper"
	"github.com/harisaginting/krdv/model"
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

	p := model.RequestPage{
		Sort:  c.Query("sort"),
		Order: c.Query("order"),
		Page:  helper.ForceInt(c.Query("page")),
	}

	res, err := ctrl.service.List(ctx, p)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// return
	response.Json(c, res)
	return
}

func (ctrl *Controller) Get(c *gin.Context) {

	id := helper.ForceInt(c.Param("id"))

	res, err := ctrl.service.Get(c, id)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// return
	response.Json(c, res)
	return
}
