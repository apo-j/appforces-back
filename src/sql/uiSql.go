package sql

func Page() string{//loaded always true for page
	str := `SELECT page.PageTypeId, true as IsLoaded, page.TemplateUrl, page.Layout, page.Style, page.Css
		 FROM page
		 WHERE page.appId = ?
		 AND page.id = ? `

	return str
}
//IFNULL(componentData.SliderInterval, 0) AS SliderInterval,
func Components() string{
	str := `SELECT
			component.Id,
			component.ComponentTypeId,
			component.TemplateUrl,
			component.IsLoaded,
			component.Style,
			component.Css,
			IFNULL(component.ParentId, -1) AS ParentId,
			IFNULL(componentData.SliderInterval, -1) AS SliderInterval,
			ComponentData.SliderImage,
			ComponentData.SliderDescription,
			ComponentData.SliderThumb,
			ComponentData.SliderThumbDescription,
			ComponentData.SliderUrl,
			ComponentData.listItemLinkUrl,
			ComponentData.listItemLinkTarget,
			ComponentData.listItemLinkLabel,
			ComponentData.searchSearchId,
			ComponentData.searchCriteriaKey,
			ComponentData.searchCriteriaValue,
			ComponentData.searchCriteriaOperator,
			ComponentData.searchCriteriaIsAnd,
			ComponentData.navbarItemUrl,
			ComponentData.navbarItemName,
			ComponentData.navbarItemIcon,
			ComponentData.NavbarChildTemplateUrl,
			ComponentData.htmlContent,
			component.order,
			component.position
		 FROM component
		 INNER JOIN page ON component.pageId = page.Id
		 LEFT JOIN ComponentData ON Component.Id = ComponentData.ComponentId
		 WHERE page.appId = ?
		 AND page.Id = ?
		 AND component.Position & ? > 0 `

	return str
}

