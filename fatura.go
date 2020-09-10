package webmania

type Fatura struct {
	Numero       string `json:"numero,omitempty"`
	Valor        string `json:"valor,omitempty"`
	Desconto     string `json:"desconto,omitempty"`
	ValorLiquido string `json:"valor_liquido,omitempty"`
}
