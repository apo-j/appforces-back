package controllers

import (
	"github.com/go-martini/martini"
	"models"
	"net/http"
	"engines"
	"json"
	_ "log"
)

type TestsController struct {
}


//func getPostCriteria(r *http.Request) *models.AppDataSearchCriteria {
//	band, title, yrs := r.FormValue("band"), r.FormValue("title"), r.FormValue("year")
//	yri, err := strconv.Atoi(yrs)
//	if err != nil {
//		yri = 0 // Year is optional, set to 0 if invalid/unspecified
//	}
//	return &Album{
//		Band:  band,
//		Title: title,
//		Year:  yri,
//	}
//}


//User-Agent: Fiddler
//Content-Type: application/json


//{"Criteria":[
//{
//" Key":"Title" ,
//"Value":"铸铁锅",
//"Operator":null
//},
//{
//"Key":"Price" ,
//"Value":100,
//"Operator":null
//}],
//"IsAnd":true
//}

//search
func (controller *TestsController) Get(appEngine *engines.AppEngine, params martini.Params, criteria models.AppDataSearchCriteria) {

//	criteria := models.AppDataSearchCriteria{
//		Criteria: []models.AppDataSearchCriterion{ models.AppDataSearchCriterion{ Key: "Title", Value:"铸铁锅"}, models.AppDataSearchCriterion{Key: "Price", Value:"100"}},
//		IsAnd: true,
//	}

	tt :=  appEngine.DataEngine.SearchAppData(1, "Article", criteria)

	appEngine.Renderer.JSON(http.StatusOK, tt)
}

func (controller *TestsController) New(appEngine *engines.AppEngine, fields json.AppDataFieldsJSON) {
	appEngine.DataEngine.CreateAppData(1, "Article1", fields.Data)
	appEngine.Renderer.JSON(http.StatusOK, fields)
}

func (controller *TestsController) Create(appEngine *engines.AppEngine, app models.App) {
	appEngine.DataEngine.InsertApp(app)

	appEngine.Renderer.Redirect("/")
}
