package models

type PageConfig struct{
	PageTypeId			    int64		`json:"Type"`
	IsLoaded				bool
	TemplateUrl			*string
	Layout					int			`json:"-"`
	Style					*string		//css style
	Css						*string		//css class
	//todo add search criteria
	Center					Items
	Top						Items
	Left					Items
	Right					Items
	Bottom					Items
}
