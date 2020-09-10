package webmania

type Exportacao struct {
	UFEmbarque    string `json:"uf_embarque,omitempty"`
	LocalEmbarque string `json:"local_embarque,omitempty"`
	LocalDespacho string `json:"local_despacho,omitempty"`
}
