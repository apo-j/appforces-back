package models

type AppConfig struct{
	AppId								int64
	AppName 							string 		//a string cannot be nil
	FaviconUrl   						*string 	//a *string can be nil
	AppTouchFaviconUrl    			*string
	Url									string
	Scripts							[]*string
	Styles								[]*string
	Pages								[]*Page
}

