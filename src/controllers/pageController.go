package controllers

import (
	"github.com/go-martini/martini"
	"engines"
	"net/http"
	"models"
	"utils/layout"
	"fmt"
	_ "log"
	"strconv"
)

type PageController struct {
}

func (controller *PageController) Get(appEngine *engines.AppEngine, params martini.Params) {
	appId, _ := strconv.ParseInt(params["appId"], 10, 64)
	pageId, _ := strconv.ParseInt(params["pageId"], 10, 64)

	page := appEngine.DataEngine.FetchPage(appId,pageId)

	if &page == nil{
		appEngine.Renderer.JSON(http.StatusNotFound, NewError(ErrCodeNotExist, fmt.Sprintf("the album with id %s does not exist", params["url"])))
	}

	items :=  appEngine.DataEngine.FetchPageComponents(appId, pageId, page.Layout)

	page.Center = models.GetComponentsForLayout(items, layout.Center)

	if layout.HasTop(page.Layout){
		page.Top = models.GetComponentsForLayout(items, layout.Top)
	}

	if layout.HasLeft(page.Layout){
		page.Left =  models.GetComponentsForLayout(items, layout.Left)
	}

	if layout.HasRight(page.Layout){
		page.Right =  models.GetComponentsForLayout(items, layout.Right)
	}

	if layout.HasBottom(page.Layout){
		page.Bottom = models.GetComponentsForLayout(items, layout.Bottom)
	}
	appEngine.Renderer.JSON(http.StatusOK, page)
}



