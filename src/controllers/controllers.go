package controllers

type Controllers struct {
	ConfigController		ConfigController
	AppDataController		AppDataController
	TestController			TestsController
	PageController			PageController
}

func InitControllers() Controllers{
	controllers := Controllers{
		ConfigController: ConfigController{},
		AppDataController:  AppDataController{},
		TestController: TestsController{},
		PageController: PageController{},
	}

	return controllers
}

