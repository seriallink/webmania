package webmania

type Pedido struct {
	Presenca                  int      `json:"presenca,omitempty"`
	ModalidadeFrete           int      `json:"modalidade_frete,omitempty"`
	Frete                     string   `json:"frete,omitempty"`
	Desconto                  string   `json:"desconto,omitempty"`
	Total                     string   `json:"total,omitempty"`
	DespesasAcessorias        string   `json:"despesas_acessorias,omitempty"`
	DespesasAduaneiras        string   `json:"despesas_aduaneiras,omitempty"`
	InformacoesFisco          string   `json:"informacoes_fisco,omitempty"`
	InformacoesComplementares string   `json:"informacoes_complementares,omitempty"`
	ObservacoesContribuinte   []string `json:"observacoes_contribuinte,omitempty"`
}
