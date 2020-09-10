package webmania

type Local struct {
	CNPJ         int    `json:"cnpj,omitempty"`
	RazaoSocial  string `json:"razao_social,omitempty"`
	IE           int    `json:"ie,omitempty"`
	CPF          string `json:"cpf,omitempty"`
	NomeCompleto string `json:"nome_completo,omitempty"`
	Endereco     string `json:"endereco,omitempty"`
	Numero       string `json:"numero,omitempty"`
	Complemento  string `json:"complemento,omitempty"`
	Bairro       string `json:"bairro,omitempty"`
	Cidade       string `json:"cidade,omitempty"`
	UF           string `json:"uf,omitempty"`
	Pais         string `json:"pais,omitempty"`
	CEP          int    `json:"cep,omitempty"`
	Telefone     int    `json:"telefone,omitempty"`
	Email        string `json:"email,omitempty"`
}
