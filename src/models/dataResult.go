package models

type DataResult struct{
	ID				int64
	Name			*string
	FieldName		*string
	DataTypeID		int64
	FieldIndex		int16
	HasMultipleValue bool
	ValueSeparator  *string
	IsReference		bool
	MediaUrl		*string
	MediaTitle		*string
	MediaWidth		*int16
	MediaHeight		*int16
	MediaTypeID		*int64
	V1				*string
	V2				*string
	V3				*string
	V4				*string
	V5				*string
	V500			*string
}

func DataResultGroupById(self []DataResult) map[int64][]DataResult{
	res := make(map[int64][]DataResult)

	for i:=0; i<len(self); i++ {
		res[self[i].ID] = append(res[self[i].ID], self[i])
	}
	return res
}

func GroupByFieldName(self []DataResult) map[string][]DataResult{
	res := make(map[string][]DataResult)

	for _, data := range self{
		res[*data.FieldName] = append(res[*data.FieldName], data)
	}

	return res
}

