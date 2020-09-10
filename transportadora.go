package webmania

type Transportadora struct {
	Cnpj         int    `json:"cnpj,omitempty"`
	RazaoSocial  string `json:"razao_social,omitempty"`
	Ie           string `json:"ie,omitempty"`
	Cpf          int    `json:"cpf,omitempty"`
	NomeCompleto string `json:"nome_completo,omitempty"`
	Endereco     string `json:"endereco,omitempty"`
	Uf           string `json:"uf,omitempty"`
	Cidade       string `json:"cidade,omitempty"`
	Cep          string `json:"cep,omitempty"`
	Placa        string `json:"placa,omitempty"`
	UFVeiculo    string `json:"uf_veiculo,omitempty"`
	RNTC         string `json:"rntc,omitempty"`
	Seguro       string `json:"seguro,omitempty"`
}
