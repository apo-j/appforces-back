package controllers

import (
	"github.com/go-martini/martini"
	"engines"
	"net/http"
	"fmt"
	_ "log"
	"strconv"
	"models"
)

type AppDataController struct {
}

func (controller *AppDataController) Get(appEngine *engines.AppEngine, params martini.Params) {
	appId, _ := strconv.ParseInt(params["appId"], 10, 64)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	appData := appEngine.DataEngine.FetchAppData(appId, params["dataName"], id)
	if &appData == nil{
		appEngine.Renderer.JSON(http.StatusNotFound, NewError(ErrCodeNotExist, fmt.Sprintf("the data with id %s does not exist", params["id"])))
	}

	appEngine.Renderer.JSON(http.StatusOK, appData)
}

func (controller *AppDataController) Search(appEngine *engines.AppEngine, params martini.Params, criteria models.AppDataSearchCriteria) {

	//	criteria := models.AppDataSearchCriteria{
	//		Criteria: []models.AppDataSearchCriterion{ models.AppDataSearchCriterion{ Key: "Title", Value:"铸铁锅"}, models.AppDataSearchCriterion{Key: "Price", Value:"100"}},
	//		IsAnd: true,
	//	}

	appId, _ := strconv.ParseInt(params["appId"], 10, 64)
	appEngine.Renderer.JSON(http.StatusOK, appEngine.DataEngine.SearchAppData(appId, params["dataName"], criteria))
}
