package webmania

type NotaFiscal struct {
	ID               string       `json:"id,omitempty"`
	Operacao         int          `json:"operacao,omitempty"`
	NaturezaOperacao string       `json:"natureza_operacao,omitempty"`
	Modelo           string       `json:"modelo,omitempty"`
	Finalidade       int          `json:"finalidade,omitempty"`
	Ambiente         int          `json:"ambiente,omitempty"`
	UrlNotificacao   string       `json:"url_notificacao,omitempty"`
	Cliente          *Cliente     `json:"cliente,omitempty"`
	Produtos         []Produto    `json:"produtos,omitempty"`
	Pedido           *Pedido      `json:"pedido,omitempty"`
	Transporte       []Transporte `json:"transporte,omitempty"`
	Fatura           []Fatura     `json:"fatura,omitempty"`
	Parcelas         []Parcela    `json:"parcelas,omitempty"`
	Exportacao       []Exportacao `json:"exportacao,omitempty"`
}

type NotaFiscalRetorno struct {
	UUID   string `json:"uuid"`
	Status string `json:"status"`
	NFE    string `json:"nfe"`
	Serie  string `json:"serie"`
	Chave  string `json:"chave"`
	Modelo string `json:"modelo"`
	XML    string `json:"xml"`
	DANFE  string `json:"danfe"`
	Log    *Log   `json:"log"`
}

func (a *Api) EmiteNotaFiscal(nf *NotaFiscal) (*NotaFiscalRetorno, error) {
	model := new(NotaFiscalRetorno)
	err := a.Post("emissao/", nf, nil, model)
	return model, err
}

func (a *Api) ConsultaNotaFiscal(uuid string) (*NotaFiscalRetorno, error) {
	params := map[string]string{
		"uuid": uuid,
	}
	model := new(NotaFiscalRetorno)
	err := a.Get("consulta/", params, nil, model)
	return model, err
}

func (a *Api) CancelarNotaFiscal(uuid, motivo string) (*NotaFiscalRetorno, error) {
	params := map[string]string{
		"uuid":   uuid,
		"motivo": motivo,
	}
	model := new(NotaFiscalRetorno)
	err := a.Put("cancelar/", params, nil, model)
	return model, err
}
