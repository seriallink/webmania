package webmania

type IPI struct {
	SituacaoTributaria  string `json:"situacao_tributaria,omitempty"`
	CodigoEnquadramento int    `json:"codigo_enquadramento,omitempty"`
	Aliquota            string `json:"aliquota,omitempty"`
	PercentualDevolvido string `json:"percentual_devolvido,omitempty"`
	CodigoSelo          string `json:"codigo_selo,omitempty"`
	QtdSelo             string `json:"qtd_selo,omitempty"`
}
