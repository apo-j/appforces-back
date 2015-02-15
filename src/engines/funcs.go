package engines

import (
	"models"
)

func CreateDataEngine() DataEngine{
	return DataEngine{DbContext: models.NewDbContext()}
}
