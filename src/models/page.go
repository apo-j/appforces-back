package models

type Page struct{
	Id						int64
	Title					*string
	Code					*string
	Url						*string
	IsIndexPage			bool
	LayoutUrl				*string
	PageTypeId				int64  		`json:"Type"`
	Ctrl					*string
	IsActive 				bool
}


