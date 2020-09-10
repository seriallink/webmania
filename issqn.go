package webmania

type ISSQN struct {
	CodigoCFOP             string `json:"codigo_cfop,omitempty"`
	Exigibilidade          int    `json:"exigibilidade,omitempty"`
	ItemServico            string `json:"item_servico,omitempty"`
	IncentivoFiscal        string `json:"incentivo_fiscal,omitempty"`
	Aliquota               string `json:"aliquota,omitempty"`
	Deducao                string `json:"deducao,omitempty"`
	Retencoes              string `json:"retencoes,omitempty"`
	DescontoIncondicionado string `json:"desconto_incondicionado,omitempty"`
	DescontoCondicionado   string `json:"desconto_condicionado,omitempty"`
	IssRetido              string `json:"iss_retido,omitempty"`
	Municipio              string `json:"municipio,omitempty"`
	CodigoServico          string `json:"codigo_servico,omitempty"`
	Processo               string `json:"processo,omitempty"`
}
