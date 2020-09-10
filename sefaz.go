package webmania

type Sefaz struct {
	Status string `json:"status"`
}

func (a *Api) CheckStatus() (*Sefaz, error) {
	model := new(Sefaz)
	err := a.Get("sefaz/", nil, nil, model)
	return model, err
}
