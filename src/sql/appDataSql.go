package sql

//Appid, MetadataName, ID, ID
func AppData() string{
	str := ` SELECT MetaData.Name, MetaDataField.FieldName,MetaDataField.FieldIndex,MetaDataField.HasMultipleValue, MetaDataField.ValueSeparator, MetaDataField.DataTypeID,MetaDataField.IsReference,
	Media.Url AS MediaUrl, Media.Title AS MediaTitle, Media.Width AS MediaWidth, Media.Height AS MediaHeight, Media.MediaTypeID,
	AppData.Id, AppData.V1, AppData.V2, AppData.V3, AppData.V4, AppData.V5, AppData.V500
	FROM MetaData
	INNER JOIN MetaDataField ON MetaDataField.MetaDataID = MetaData.ID
	INNER JOIN AppData ON AppData.MetaDataID = MetaData.ID
	LEFT JOIN LinkMediaAppData ON LinkMediaAppData.appDataId = AppData.ID
	LEFT JOIN Media ON Media.ID = LinkMediaAppData.MediaID
	WHERE MetaData.AppID = ?
	AND MetaData.IsActive = 1
	AND MetaData.Name = ?
	AND AppData.IsActive = 1
	AND (AppData.ID = ? OR ? is NULL OR ? <= 0)
	ORDER BY AppData.ID `

	return str
}

func AppDataConfig() string{
	str := ` SELECT MetaDataField.FieldName, MetaDataField.FieldIndex, MetaDataField.DataTypeId, MetaDataField.HasMultipleValue, MetaDataField.ValueSeparator
			FROM MetaData
			INNER JOIN MetaDataField ON MetaDataField.MetaDataID = MetaData.ID
			WHERE MetaData.AppID = ?
			AND MetaData.IsActive = 1
			AND MetaData.Name = ?
			And MetaData.IsActive = 1 `

	return str
}

func Search() string{
	str := ` SELECT {0}
			 FROM AppData
			 INNER JOIN (
			  	SELECT AppData.ID, MetaDataField.FieldName, MetaDataField.FieldIndex, MetaDataField.DataTypeID, MetaDataField.HasMultipleValue, MetaDataField.ValueSeparator,MetaDataField.IsReference,
			  	Media.Url AS MediaUrl, Media.Title AS MediaTitle, Media.Width AS MediaWidth, Media.Height AS MediaHeight, Media.MediaTypeID
			  	FROM MetaData
			  	INNER JOIN MetaDataField ON MetaDataField.MetaDataId = MetaData.ID
			  	INNER JOIN AppData ON AppData.MetaDataId = MetaData.ID
			  	LEFT JOIN LinkMediaAppData ON LinkMediaAppData.appDataId = AppData.ID
				LEFT JOIN Media ON Media.ID = LinkMediaAppData.MediaID
			  	WHERE MetaData.AppId = ? AND MetaData.IsActive AND MetaData.Name = ? AND AppData.IsActive = 1 AND ({1})
			 ) AS Res ON Res.ID = AppData.ID `

	return str
}

func CreateMetaData() string{
	str := ` INSERT INTO MetaData(AppID, IsActive, Name) VALUES(?, 1, ?); SELECT LAST_INSERT_ID(); `

	return str
}

func CreateMetaDataField() string{
	str := ` INSERT INTO MetaDataFields(AppID, MetaDataID, FieldName, DataTypeID, FieldIndex, IsIndexed) VALUES (?, ?, ?, ?, ?, ?)

	`
	return str
}
