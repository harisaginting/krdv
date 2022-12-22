package watchlist

import (
	"encoding/json"
	"io/ioutil"

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

func (ctrl *Controller) Create(c *gin.Context) {

	var requestBody RequestWatchlist
	request, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	err = json.Unmarshal([]byte(request), &requestBody)
	if err != nil {
		response.BadRequest(c)
		return
	}

	res, err := ctrl.service.Create(c, requestBody)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// return
	response.Json(c, res)
	return
}

func (ctrl *Controller) Update(c *gin.Context) {
	id := helper.ForceInt(c.Param("id"))
	var requestBody RequestWatchlist
	request, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	err = json.Unmarshal([]byte(request), &requestBody)
	if err != nil {
		response.BadRequest(c)
		return
	}

	res, err := ctrl.service.Update(c, id, requestBody)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// return
	response.Json(c, res)
	return
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id := helper.ForceInt(c.Param("id"))

	res, err := ctrl.service.Delete(c, id)
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

	res, err := ctrl.service.GetByUser(c, id)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// return
	response.Json(c, res)
	return
}

func (ctrl *Controller) List(c *gin.Context) {
	p := model.RequestPage{
		Sort:  c.Query("sort"),
		Order: c.Query("order"),
		Page:  helper.ForceInt(c.Query("page")),
		Limit: helper.ForceInt(c.Query("limit")),
	}

	res, err := ctrl.service.List(c, p)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// return
	response.Json(c, res)
	return
}

func (ctrl *Controller) AddMovie(c *gin.Context) {

	id := helper.ForceInt(c.Param("id"))
	movieId := helper.ForceInt(c.Param("movie_id"))

	res, err := ctrl.service.AddMovie(c, id, movieId)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// return
	response.Json(c, res)
	return
}

func (ctrl *Controller) DeleteMovie(c *gin.Context) {

	id := helper.ForceInt(c.Param("id"))
	movieId := helper.ForceInt(c.Param("movie_id"))

	res, err := ctrl.service.DeleteMovie(c, id, movieId)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// return
	response.Json(c, res)
	return
}
