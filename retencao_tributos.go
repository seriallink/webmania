package webmania

type RetencaoTributos struct {
	ValorPIS               string `json:"valor_pis,omitempty"`
	ValorCOFINS            string `json:"valor_cofins,omitempty"`
	ValorCSLL              string `json:"valor_csll,omitempty"`
	BaseCalculoIRRF        string `json:"bc_irrf,omitempty"`
	ValorIRRF              string `json:"valor_irrf,omitempty"`
	BaseCalculoPrevidencia string `json:"bc_previdencia,omitempty"`
	ValorPrevidencia       string `json:"valor_previdencia,omitempty"`
}
