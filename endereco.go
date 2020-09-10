package webmania

type Endereco struct {
	Logradouro  string `json:"endereco"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	UF          string `json:"uf"`
	CEP         string `json:"cep"`
}

func (a *Api) AtualizaEndereco(endereco *Endereco) error {
	message := new(SuccessMessage)
	return a.Post("empresa/", endereco, nil, message)
}
