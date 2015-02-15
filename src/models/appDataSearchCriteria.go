package models

import (
	"strings"
)

type AppDataSearchCriterion struct{
	Key 		string
	Value		interface {}
	Operator	string
}

type AppDataSearchCriteria struct{
	Criteria 	[]AppDataSearchCriterion
	IsAnd	bool //true: AND, false: or
}

func (self AppDataSearchCriteria) FilterCriterion(key string) AppDataSearchCriterion{
	for _, criterion := range self.Criteria{
		if strings.EqualFold(strings.ToUpper(criterion.Key), strings.ToUpper(key)){
			return criterion
		}
	}
	return AppDataSearchCriterion{}
}
