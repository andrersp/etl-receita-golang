package config

import (
	"github.com/andrersp/go-etl-receita-federal/src/schemas"
	"github.com/jszwec/csvutil"
)

var (
	HeaderEstabelecimentoInput []string
	HeaderEmpresaInput         []string
	HeaderComplementarInput    []string
	HeaderSociosInput          []string
	HeaderSimplesInput         []string
	HeaderCnaeSecundarua       []string
)

func createCSVHeaders() {
	HeaderEstabelecimentoInput, _ = csvutil.Header(&schemas.SchemaEstabelecimento{}, "csv")

	HeaderCnaeSecundarua, _ = csvutil.Header(&schemas.SchemaCnaeSecundaria{}, "csv")

	HeaderComplementarInput, _ = csvutil.Header(&schemas.SchemaComplementar{}, "csv")

	HeaderSociosInput, _ = csvutil.Header(&schemas.SchemaSocio{}, "csv")

	HeaderEmpresaInput, _ = csvutil.Header(&schemas.SchemaEmpresa{}, "csv")

	HeaderSimplesInput, _ = csvutil.Header(&schemas.SchemaSimples{}, "csv")

}
