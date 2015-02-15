package sql

func Config() string{
	str := `SELECT appConfig.appID, appConfig.name AS AppName, domain.Url, appConfig.FaviconUrl, appConfig.appTouchFaviconUrl
		 FROM domain
		 INNER JOIN appConfig ON appConfig.domainID = domain.ID
		 WHERE domain.url = ?
		 AND domain.isActive = 1 `
	return str
}

func Scripts() string{
	str := `SELECT script.Url
		 FROM appHasScript
		 INNER JOIN script ON appHasScript.scriptID = Script.Id
		 WHERE appHasScript.AppID = ? `
	return str
}

func Styles() string{
	str := `SELECT style.Url
		 FROM appHasStyle
		 INNER JOIN style ON appHasStyle.styleID = style.Id
		 WHERE appHasStyle.AppID = ? `
	return str
}

func Pages() string{
	str := `SELECT page.Id, page.title, page.code, page.Url, page.isIndexPage, page.layoutUrl, page.pageTypeId, page.ctrl
		 FROM page
		 WHERE page.appID = ?
		 AND page.isActive = 1 `
	return str
}


