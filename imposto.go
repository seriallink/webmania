package webmania

type Imposto struct {
	ICMS             *ICMS             `json:"icms,omitempty"`
	IPI              *IPI              `json:"ipi,omitempty"`
	PIS              *PIS              `json:"pis,omitempty"`
	COFINS           *COFINS           `json:"cofins,omitempty"`
	ISSQN            *ISSQN            `json:"issqn,omitempty"`
	RetencaoTributos *RetencaoTributos `json:"retencao_tributos,omitempty"`
	Importacao       *Importacao       `json:"importacao,omitempty"`
}
