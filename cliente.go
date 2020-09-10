package webmania

type Cliente struct {
	CPF                  string `json:"cpf,omitempty"`
	NomeCompleto         string `json:"nome_completo,omitempty"`
	CNPJ                 string `json:"cnpj,omitempty"`
	RazaoSocial          string `json:"razao_social,omitempty"`
	IE                   string `json:"ie,omitempty"`
	Suframa              string `json:"suframa,omitempty"`
	SubstitutoTributario string `json:"substituto_tributario,omitempty"`
	ConsumidorFinal      int    `json:"consumidor_final,omitempty"`
	Contribuinte         int    `json:"contribuinte,omitempty"`
	Microcervejaria      bool   `json:"microcervejaria,omitempty"`
	Logradouro           string `json:"endereco,omitempty"`
	Complemento          string `json:"complemento,omitempty"`
	Numero               string `json:"numero,omitempty"`
	Bairro               string `json:"bairro,omitempty"`
	Cidade               string `json:"cidade,omitempty"`
	UF                   string `json:"uf,omitempty"`
	CEP                  string `json:"cep,omitempty"`
	Telefone             string `json:"telefone,omitempty"`
	Email                string `json:"email,omitempty"`
}
