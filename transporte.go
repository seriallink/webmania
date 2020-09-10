package webmania

type Transporte struct {
	Volume         string          `json:"volume,omitempty"`
	Especie        string          `json:"especie,omitempty"`
	PesoBruto      string          `json:"peso_bruto,omitempty"`
	PesoLiquido    string          `json:"peso_liquido,omitempty"`
	Marca          string          `json:"marca,omitempty"`
	Numeracao      string          `json:"numeracao,omitempty"`
	Lacres         string          `json:"lacres,omitempty"`
	Transportadora *Transportadora `json:"transportadora,omitempty"`
	Reboque        *Reboque        `json:"reboque,omitempty"`
	Entrega        *Local          `json:"entrega,omitempty"`
	Retirada       *Local          `json:"retirada,omitempty"`
}
