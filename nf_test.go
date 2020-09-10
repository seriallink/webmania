package webmania

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmiteNotaFiscal(t *testing.T) {
	nf := &NotaFiscal{
		Operacao:         1,
		Finalidade:       1,
		Ambiente:         2,
		NaturezaOperacao: "REMESSA",
		Cliente: &Cliente{
			CPF:          "01595480803",
			NomeCompleto: "WILSON JOSE VALENTINI",
			Logradouro:   "RODOVIA UNAI/BONFINÓPOLIS KM 55",
			Bairro:       "ZONA RURAL",
			Cidade:       "BONFINOPOLIS DE MINAS",
			UF:           "MG",
			CEP:          "38650000",
		},
		Produtos: []Produto{{
			Nome:       "MILHO - MILHO EM GRAOS",
			NCM:        "11081200",
			Quantidade: "39.020,00",
			Unidade:    "KG",
			Origem:     "0",
			Total:      "31.863,73",
			Subtotal:   "31.863,73",
			Impostos: &Imposto{
				ICMS: &ICMS{
					CodigoCFOP:         "5905",
					SituacaoTributaria: "101",
					AliquotaCredito:    "0.00",
				},
				IPI: &IPI{
					SituacaoTributaria: "101",
					Aliquota:           "0.00",
				},
				PIS: &PIS{
					SituacaoTributaria: "101",
					Aliquota:           "0.00",
				},
				COFINS: &COFINS{
					SituacaoTributaria: "101",
					Aliquota:           "0.00",
				},
			},
		}},
		Pedido: &Pedido{
			Presenca:        1,
			ModalidadeFrete: 1,
		},
	}
	model, err := api.EmiteNotaFiscal(nf)
	assert.NoError(t, err)
	spew.Dump(model, err)
}

func TestConsultaNotaFiscal(t *testing.T) {
	model, err := api.ConsultaNotaFiscal("d5695d91-a6d6-4c6a-8426-d0573d6042d1")
	assert.NoError(t, err)
	spew.Dump(model, err)
}

func TestCancelarNotaFiscal(t *testing.T) {
	model, err := api.CancelarNotaFiscal("d5695d91-a6d6-4c6a-8426-d0573d6042d1", "Teste de cancelamento")
	assert.NoError(t, err)
	spew.Dump(model, err)
}
