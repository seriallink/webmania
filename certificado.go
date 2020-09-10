package webmania

import "encoding/base64"

func (a *Api) ValidaCertificado() (int64, error) {

	exp := struct {
		Expiration int64 `json:"expiration"`
	}{}

	err := a.Get("certificado/", nil, nil, &exp)

	return exp.Expiration, err

}

func (a *Api) AtualizaCertificado(certificado []byte, senha string) error {

	params := map[string]interface{}{
		"certificado":       base64.StdEncoding.EncodeToString(certificado),
		"certificado_senha": senha,
	}

	message := new(SuccessMessage)

	return a.Post("empresa/", params, nil, message)

}
