package webmania

type ICMS struct {
	CodigoCFOP         string `json:"codigo_cfop,omitempty"`
	SituacaoTributaria string `json:"situacao_tributaria,omitempty"`
	AliquotaCredito    string `json:"aliquota_credito,omitempty"`
	AliquotaImportacao string `json:"aliquota_importacao,omitempty"`
	Industria          string `json:"industria,omitempty"`
}
