package webmania

type Reboque struct {
	Placa     int    `json:"placa,omitempty"`
	UfVeiculo string `json:"uf_veiculo,omitempty"`
	RNTC      string `json:"rntc,omitempty"`
	Vagao     int    `json:"vagao,omitempty"`
	Balsa     string `json:"balsa,omitempty"`
}
