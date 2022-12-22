package auth

import (
	"encoding/json"
	"io/ioutil"

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

func (ctrl *Controller) Register(c *gin.Context) {
	var reqData PayloadUserRegister
	var resData ResponseUserRegister
	request, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.BadRequest(c)
		return
	}
	json.Unmarshal(request, &reqData)
	resData, err = ctrl.service.Register(c.Request.Context(), reqData)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Json(c, resData)
	return
}

func (ctrl *Controller) Login(c *gin.Context) {
	var reqData PayloadUserLogin
	var resData ResponseUserLogin
	request, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.BadRequest(c)
		return
	}
	json.Unmarshal(request, &reqData)
	err, resData = ctrl.service.Login(c.Request.Context(), reqData)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Json(c, resData)
	return
}

func (ctrl *Controller) Me(c *gin.Context) {

	var resData ResponseMe
	username := c.Value("username").(string)
	err, resData := ctrl.service.GetByUsername(c.Request.Context(), username)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Json(c, resData)
	return
}

func (ctrl *Controller) List(c *gin.Context) {
	var resData ResponseList
	ctrl.service.List(c.Request.Context(), &resData)

	// return
	response.Json(c, resData)
	return
}
