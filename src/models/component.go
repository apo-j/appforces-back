package models

import (
	"sort"
)

//component
type Items []Item

type Item struct{
	Id			    				int64
	ComponentTypeId			 	int64		`json:"Type"`
	TemplateUrl					*string
	IsLoaded						bool
	//todo add search criteria
	Style							*string    //css style
	Css								*string 	//css classes
	Items							Items
    ParentId						int64    `json:"-"`
	Data							interface {}
	Order							int		 `json:"-"`
}

type ComponentData struct{
	Id			    				int64
	ComponentTypeId			 	int64
	TemplateUrl					*string
	IsLoaded						bool
	Style							*string //css style
	Css								*string //css classes
	ParentId						*int64 //sql.NullInt64  //for coopernurse mapping nil value to int64
    SliderInterval				*int
	SliderImage					*string
	SliderDescription				*string
	SliderThumb					*string
	SliderThumbDescription		*string
	SliderUrl					    *string
	ListItemLinkUrl					*string
	ListItemLinkTarget				*string
	ListItemLinkLabel					*string
	SearchSearchId					*string
	SearchCriteriaKey				*string
	SearchCriteriaValue					*string
	SearchCriteriaOperator			*string
	SearchCriteriaIsAnd					*bool
	NavbarItemUrl						*string
	NavbarItemName					*string
	NavbarItemIcon					*string
	NavbarChildTemplateUrl				*string
	HtmlContent						*string
	Order							int
	Position						int
}

type ComponentDataSlider struct{
	Interval		 int
	Sliders         []Slider
}

type Slider struct{
	Image					*string
	Description			*string
	Thumb					*string
	ThumbDescription		*string
	Url						*string
}

type  ListItemLinkLabel struct{
	Url					*string
	Target				*string
	Label				*string
}

type  Search struct{
	Id					*string
	Criteria			[]SearchCriteria
	IsAnd				bool
}

type SearchCriteria struct{
	Key					*string
	Value				*string
	Operator			*string
}

type NavbarItemData struct{
	Icon					*string
	Name					*string
	Url						*string
}

type NavbarData struct{
	ChildTemplateUrl					*string
}

type HtmlContentData struct{
	Description					*string
}

type ContainerArticlesData struct{
	Criteria			[]SearchCriteria
	IsAnd				bool
}

func GroupById(self []ComponentData, layout int) map[int64][]ComponentData{
	res := make(map[int64][]ComponentData)

	for _, value := range self {
		if value.Position & layout > 0{
			res[value.Id] = append(res[value.Id], value)
		}
	}
	return res
}

func MakeHierarchy(currentNode *Item, items []Item) []Item{
	if currentNode != nil{
		for _, value := range items{
			if value.ParentId == currentNode.Id{
				//remove the element from the slice ??
				value.Items = MakeHierarchy(&value, items)
				currentNode.Items = append(currentNode.Items, value)
			}
		}
		sort.Sort(currentNode.Items)
		return currentNode.Items
	}else{
		//find root items
		roots := make(Items, 0)
		for _, value := range items{
			if value.ParentId <= 0 {
				//remove the element from the slice ??
				value.Items = MakeHierarchy(&value, items)
				roots = append(roots, value)
			}
		}
		sort.Sort(roots)
		return roots
	}
}



func (self Items) Len() int {return len(self)}
func (self Items) Less(i,j int) bool {return self[i].Order < self[j].Order}
func (self Items) Swap(i,j int) {self[i],self[j] = self[j],self[i]}

func GetComponentsForLayout(componentData []ComponentData, layout int) Items{
	components := GroupById(componentData, layout)

	items := make([]Item, 0)

	for k, _ := range components{
		component := Item{
			Id : components[k][0].Id,
			ComponentTypeId : components[k][0].ComponentTypeId,
			TemplateUrl : components[k][0].TemplateUrl,
			IsLoaded : components[k][0].IsLoaded,
			Style : components[k][0].Style,
			Css : components[k][0].Css,
			ParentId : *components[k][0].ParentId,
			Order : components[k][0].Order,
		}

		switch components[k][0].ComponentTypeId{
			case 2://carousel
				sliders := make([]Slider, 0)
				for _, value := range components[k] {
					slide := Slider{
						Image : value.SliderImage,
						Thumb : value.SliderThumb,
						Description : value.SliderDescription,
						ThumbDescription : value.SliderThumbDescription,
						Url : value.SliderUrl,
					}
					sliders = append(sliders, slide)
				}

				component.Data = ComponentDataSlider{
					Interval :*components[k][0].SliderInterval,
					Sliders : sliders,
				}

				break
			case 5://ListItemLink
				component.Data = ListItemLinkLabel{
					Url :components[k][0].ListItemLinkUrl,
					Target : components[k][0].ListItemLinkTarget,
					Label : components[k][0].ListItemLinkLabel,
				}
				break
			case 7://search
				criteria := make([]SearchCriteria, 0)
				for _, value := range components[k] {
					criterion := SearchCriteria{
						Key : value.SearchCriteriaKey,
						Value : value.SearchCriteriaValue,
						Operator: value.SearchCriteriaOperator,
					}
					criteria = append(criteria, criterion)
				}

				searchData := Search{
					Id :components[k][0].SearchSearchId,
					Criteria : criteria,
				}

				if components[k][0].SearchCriteriaIsAnd != nil{
					searchData.IsAnd = *components[k][0].SearchCriteriaIsAnd
				}else{
					searchData.IsAnd = false
				}
				component.Data = searchData
				break
			case 11://navbar
				component.Data = NavbarData{
					ChildTemplateUrl :components[k][0].NavbarChildTemplateUrl,
				}
				break
			case 12://navbarItem
				component.Data = NavbarItemData{
					Url :components[k][0].NavbarItemUrl,
					Icon : components[k][0].NavbarItemIcon,
					Name : components[k][0].NavbarItemName,
				}
				break
			case 13://htmlContent
				component.Data = HtmlContentData{
					Description :components[k][0].HtmlContent,
				}
				break
			case 14://ContainerArticles
				criteria := make([]SearchCriteria, 0)
				for _, value := range components[k] {
					criterion := SearchCriteria{
						Key : value.SearchCriteriaKey,
						Value : value.SearchCriteriaValue,
						Operator: value.SearchCriteriaOperator,
					}
					criteria = append(criteria, criterion)
				}

				containerArticlesData := ContainerArticlesData{
					Criteria : criteria,
				}

				if components[k][0].SearchCriteriaIsAnd != nil{
					containerArticlesData.IsAnd = *components[k][0].SearchCriteriaIsAnd
				}else{
					containerArticlesData.IsAnd = false
				}

				component.Data = containerArticlesData

				break
		}

		items = append(items, component)
	}
	return MakeHierarchy(nil, items)
}
