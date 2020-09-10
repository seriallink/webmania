package webmania

type Empresa struct {
	TipoTributacao string `json:"tipo_tributacao"`
	CNPJ           string `json:"cnpj,omitempty"`
	RazaoSocial    string `json:"razao_social,omitempty"`
	CPF            string `json:"cpf,omitempty"`
	NomeCompleto   string `json:"nome_completo,omitempty"`
	NomeFantasia   string `json:"nome_fantasia,omitempty"`
	IE             string `json:"ie,omitempty"`
	UnidadeEmpresa string `json:"unidade_empresa,omitempty"`
	Email          string `json:"email,omitempty"`
	Telefone       string `json:"telefone,omitempty"`
	Contabilidade  string `json:"contabilidade,omitempty"`
	Subdominio     string `json:"subdominio,omitempty"`
	UrlNotificacao string `json:"url_notificacao,omitempty"`
	Logomarca      string `json:"logomarca,omitempty"`
}

func (a *Api) AtualizaEmpresa(empresa *Empresa) error {
	message := new(SuccessMessage)
	return a.Post("empresa/", empresa, nil, message)
}
