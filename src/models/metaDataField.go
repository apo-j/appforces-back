package models

type MetaDataField struct{
	FieldName		*string
	DataTypeID		int64
	FieldIndex		int
	HasMultipleValue bool
	ValueSeparator  *string
	IsReference		bool
}
