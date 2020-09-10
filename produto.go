package webmania

type Produto struct {
	ID                    string   `json:"id,omitempty"`
	Item                  string   `json:"item,omitempty"`
	Nome                  string   `json:"nome,omitempty"`
	Codigo                string   `json:"codigo,omitempty"`
	NCM                   string   `json:"ncm,omitempty"`
	Quantidade            string   `json:"quantidade,omitempty"`
	Unidade               string   `json:"unidade,omitempty"`
	Peso                  string   `json:"peso,omitempty"`
	Origem                string   `json:"origem,omitempty"`
	Desconto              string   `json:"desconto,omitempty"`
	Subtotal              string   `json:"subtotal,omitempty"`
	Total                 string   `json:"total,omitempty"`
	ClasseImposto         string   `json:"classe_imposto,omitempty"`
	Impostos              *Imposto `json:"impostos,omitempty"`
	InformacoesAdicionais string   `json:"informacoes_adicionais,omitempty"`
	BeneficioFiscal       string   `json:"beneficio_fiscal,omitempty"`
	IndEscala             string   `json:"ind_escala,omitempty"`
	CNPJFabricante        string   `json:"cnpj_fabricante,omitempty"`
	GTIN                  string   `json:"gtin,omitempty"`
	GTINTributavel        string   `json:"gtin_tributavel,omitempty"`
	CEST                  string   `json:"cest,omitempty"`
	NVE                   string   `json:"nve,omitempty"`
	NRECOPI               string   `json:"nrecopi,omitempty"`
	AtivoPermanente       string   `json:"ativo_permanente,omitempty"`
	VeiculoUsado          string   `json:"veiculo_usado,omitempty"`
	ExIPI                 string   `json:"ex_ipi,omitempty"`
}
