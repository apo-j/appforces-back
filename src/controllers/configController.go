package controllers

import (
	"github.com/go-martini/martini"
	"engines"
	"net/http"
	"fmt"
	_ "log"
)

type ConfigController struct {
}

func (controller *ConfigController) Get(appEngine *engines.AppEngine, params martini.Params) {
	config := appEngine.DataEngine.FetchAppConfig(params["url"])

	if &config == nil{
		appEngine.Renderer.JSON(http.StatusNotFound, NewError(ErrCodeNotExist, fmt.Sprintf("the album with id %s does not exist", params["url"])))
	}
	appEngine.Renderer.JSON(http.StatusOK, config)
}



