package engines

import (
   	"log"
	"models"
	"sql"
	 "strconv"
	"reflect"
	_"fmt"
	"strings"
	"json"
)

type DataEngine struct{
	DbContext 	models.DbContext
}

func (self DataEngine) Dispose(){
	self.DbContext.Dbmap.Db.Close()
}

func (self DataEngine) FetchAppConfig(url string) models.AppConfig{
	var appConfig models.AppConfig

	err := self.DbContext.Dbmap.SelectOne(&appConfig, sql.Config(), url)
	if err != nil {panic(err)}

	//get scripts of app
	_, err = self.DbContext.Dbmap.Select(&appConfig.Scripts, sql.Scripts(), appConfig.AppId)
	if err != nil {panic(err)}

	//get styles of app
	_, err = self.DbContext.Dbmap.Select(&appConfig.Styles, sql.Styles(), appConfig.AppId)
	if err != nil {panic(err)}

	//get pages of app
	_, err = self.DbContext.Dbmap.Select(&appConfig.Pages, sql.Pages(), appConfig.AppId)
	if err != nil {panic(err)}

	return appConfig
}

/*************************************** AppData ****************************************************/
//Fetch
func (self DataEngine) FetchAppData(appId int64, objName string, id interface {}) interface {}{//[]models.DataResult{//use interface to allow id to be nil
	var dataResults []models.DataResult

	_, err := self.DbContext.Dbmap.Select(&dataResults, sql.AppData(), appId, objName, id, id, id)
	if err != nil {panic(err)}

	return MapRawDataToObject(dataResults)
}

//Search
func (self DataEngine) SearchAppData(appId int64, objName string, criteria models.AppDataSearchCriteria) interface {}{
	if configures := self.fetchAppDataConfig(appId, objName); len(configures) != 0 {
		fields := make([]string, 0)
		predicates := make([]string, 0)
		fields = append(fields, "AppData.ID");
		fields = append(fields, "Res.FieldName");
		fields = append(fields, "Res.FieldIndex");
		fields = append(fields, "Res.DataTypeID");
		fields = append(fields, "Res.HasMultipleValue");
		fields = append(fields, "Res.ValueSeparator");
		fields = append(fields, "Res.MediaUrl");
		fields = append(fields, "Res.MediaTitle");
		fields = append(fields, "Res.MediaWidth");
		fields = append(fields, "Res.MediaHeight");
		fields = append(fields, "Res.MediaTypeID");
		fields = append(fields, "Res.IsReference");

		for _, config := range configures{
			fields = append(fields, "AppData.V" + strconv.Itoa(config.FieldIndex))
			if criterion := criteria.FilterCriterion(*config.FieldName); criterion.Key != ""{
				predicateValue := ""
				dataValue :=""
				switch config.DataTypeID {
				case 1://int
					predicateValue = "convert('" + reflect.ValueOf(criterion.Value).String() +"', decimal)"
					dataValue = "convert(AppData.V" +strconv.Itoa(config.FieldIndex) +", decimal)"
				case 2://float
					predicateValue = "convert('" + reflect.ValueOf(criterion.Value).String() +"', decimal)"
					dataValue = "convert(AppData.V" +strconv.Itoa(config.FieldIndex) +", decimal)"
				case 3://string
					predicateValue = "'" + reflect.ValueOf(criterion.Value).String() + "'"
					dataValue = "AppData.V" +strconv.Itoa(config.FieldIndex)
				case 4://Boolean
					//todo
					predicateValue =  "convert('" + reflect.ValueOf(criterion.Value).String() +"', decimal)"
					dataValue = "convert(AppData.V" +strconv.Itoa(config.FieldIndex) +", decimal)"
				}

				switch strings.TrimSpace(criterion.Operator){
					case "!=", "<>":
						predicates = append(predicates,  dataValue + " != " + predicateValue)
					case "<":
						predicates = append(predicates,  dataValue + " < " + predicateValue)
					case "<=":
						predicates = append(predicates,  dataValue + " <= " + predicateValue)
					case ">":
						predicates = append(predicates,  dataValue + " > " + predicateValue)
					case ">=":
						predicates = append(predicates,  dataValue + " >= " + predicateValue)
					case "=":
						predicates = append(predicates,  dataValue + " = " + predicateValue)
				}
			}
		}

		sql := sql.Search()
		sql = strings.Replace(sql, "{0}", strings.Join(fields, ","), -1)
		if len(predicates) == 0{
			sql = strings.Replace(sql, "{1}", "1=1", -1)
		}else{
			if criteria.IsAnd {
				sql = strings.Replace(sql, "{1}", strings.Join(predicates, " AND "), -1)
			}else{
				sql = strings.Replace(sql, "{1}", strings.Join(predicates, " OR "), -1)
			}
		}
		log.Println(sql)
		var dataResults []models.DataResult
		_, err := self.DbContext.Dbmap.Select(&dataResults, sql, appId, objName)
		if err != nil {panic(err)}


		return  MapRawDataToObject(dataResults)
	}

	return nil
}

//Create
func (self DataEngine) CreateAppData(appId int64, objName string, fields []json.AppDataFieldJSON){
	log.Println(appId)
	if configures := self.fetchAppDataConfig(appId, objName); configures == nil{
		//use transaction

		var metaDataId int64
		err := self.DbContext.Dbmap.SelectOne(&metaDataId, sql.CreateMetaData(), appId, objName)
		if err != nil {panic(err)}
		log.Println(metaDataId)
		fieldIndex := 1
		for _, f := range fields{
			self.DbContext.Dbmap.Exec(sql.CreateMetaDataField(), appId, metaDataId, f.Key, f.Value, fieldIndex, 0)

			fieldIndex++
		}
	}
}

//func (self DataEngine) AddAppData(appId int64, objName, )

func (self DataEngine) fetchAppDataConfig(appId int64, objName string) []models.MetaDataField{
	var config []models.MetaDataField

	_, err := self.DbContext.Dbmap.Select(&config, sql.AppDataConfig(),appId, objName)
	if err != nil {panic(err)}

	return config
}

func MapRawDataToObject(data []models.DataResult) interface {}{
	res := make([]interface {}, 0)

	//value contains all lines about one object
	for key, value := range models.DataResultGroupById(data){
		obj := make(map[string]interface {})

		//group lines in value by field name
		for fieldName, group := range models.GroupByFieldName(value){
			firstLine := group[0];
			if firstLine.IsReference {
				values := make([]interface {}, 0)

				switch firstLine.DataTypeID {
				case 5://Media
					for _, line := range group{
						values = append(values, models.Media{
								MediaUrl : line.MediaUrl,
								MediaTitle : line.MediaTitle,
								MediaWidth : line.MediaWidth,
								MediaHeight : line.MediaWidth,
								MediaTypeID : line.MediaTypeID,
							})
					}
				}
				obj[fieldName] = values
			}else{
				val := reflect.ValueOf(interface {}(firstLine))
				dataField := "V" + strconv.FormatInt(int64(firstLine.FieldIndex), 10)

				if firstLine.HasMultipleValue {
					separator := *firstLine.ValueSeparator
					if len(separator) == 0{
						separator = "|"
					}

					obj[fieldName] = strings.Split(val.FieldByName(dataField).Elem().String(), separator)
				}else{
					switch firstLine.DataTypeID {
					case 1://int
						v, _ :=  strconv.ParseInt(val.FieldByName(dataField).Elem().String(), 10, 64)
						obj[fieldName] = v
					case 2://float
						v, _ :=  strconv.ParseFloat(val.FieldByName(dataField).Elem().String(), 10)
						obj[fieldName] = v
					case 3://string
						obj[fieldName] =  val.FieldByName(dataField).Interface()
					case 4://Boolean
						v, _ :=  strconv.ParseBool(val.FieldByName(dataField).Elem().String())
						obj[fieldName] = v
					}
				}
			}
//			for _, line := range group{
//				val :=  reflect.ValueOf(line)
//				dataField := "V" + strconv.FormatInt(val.FieldByName("FieldIndex").Int(), 10)
//				fieldName := val.FieldByName("FieldName").Elem().String()
//
//				if val.FieldByName("HasMultipleValue").Bool() {
//					separator := val.FieldByName("ValueSeparator").Elem().String()
//					if len(separator) == 0{
//						separator = "|"
//					}
//
//					obj[fieldName] = strings.Split(val.FieldByName(dataField).Elem().String(), separator)
//				}else{
//					switch val.FieldByName("DataTypeID").Int() {
//					case 1://int
//						v, _ :=  strconv.ParseInt(val.FieldByName(dataField).Elem().String(), 10, 64)
//						obj[fieldName] = v
//					case 2://float
//						v, _ :=  strconv.ParseFloat(val.FieldByName(dataField).Elem().String(), 10)
//						obj[fieldName] = v
//					case 3://string
//						obj[fieldName] =  val.FieldByName(dataField).Interface()
//					case 4://Boolean
//						v, _ :=  strconv.ParseBool(val.FieldByName(dataField).Elem().String())
//						obj[fieldName] = v
//					}
//				}
//			}
		}

		obj["ID"] = key

		res = append(res, obj)
	}

	return res
}

func inspect(f interface{}) map[string]interface {} {
	m := make(map[string]interface {})
	val := reflect.ValueOf(f)

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		f := valueField.Interface()
		val := reflect.ValueOf(f)
		m[typeField.Name] = val.String()

		m[typeField.Name] = val.String()
		/*switch val.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				m[typeField.Name] = strconv.FormatInt(val.Int(), 10)
			case reflect.String:
				m[typeField.Name] = val.String()

		}*/
	}

	return m
}

/*func (self DataEngine) FetchAppData(appId int64, objName string, id interface {}) []models.DataResult{//use interface to allow id to be nil
	var dataResults []models.DataResult

	_, err := self.DbContext.Dbmap.Select(&dataResults, sql.AppData(), appId, objName, id, id, id)
	if err != nil {panic(err)}

	return dataResults
}*/


func (self DataEngine) FetchPageComponents(appId int64, pageId int64, layout int) []models.ComponentData {
	var componentData []models.ComponentData

	_, err := self.DbContext.Dbmap.Select(&componentData, sql.Components(), appId, pageId, layout)
	if err != nil {panic(err)}

	return componentData
}

func (self DataEngine) FetchPage(appId int64, pageId int64) models.PageConfig{
	var page models.PageConfig

	err := self.DbContext.Dbmap.SelectOne(&page, sql.Page(), appId, pageId)
	if err != nil {panic(err)}

	return page
}



func (self DataEngine) FetchApps() []models.App{
	var apps []models.App
	_, err := self.DbContext.Dbmap.Select(&apps, "SELECT * FROM APP")

	if err != nil {panic(err)}

	return apps
}

func (self DataEngine) InsertApp(app models.App){

	err := self.DbContext.Dbmap.Insert(&app)
	if err != nil {panic(err)}
}
