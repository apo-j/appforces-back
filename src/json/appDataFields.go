package json

type AppDataFieldJSON struct{
	Key string `json:"key"`
	Value int `json:"value"`
}


type AppDataFieldsJSON struct {
	Data []AppDataFieldJSON  `json:"data"`
}

