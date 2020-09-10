package webmania

type Log struct {
	BStat    bool        `json:"bStat"`
	Versao   string      `json:"versao"`
	TpAmb    string      `json:"tpAmb"`
	CStat    string      `json:"cStat"`
	VerAplic string      `json:"verAplic"`
	XMotivo  string      `json:"xMotivo"`
	DhRecbto string      `json:"dhRecbto"`
	CUF      string      `json:"cUF"`
	NRec     string      `json:"nRec"`
	AProt    interface{} `json:"aProt"`
}
