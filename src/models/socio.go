package models

type ModelSocio struct {
	ID                                   int64  `gorm:"primaryKey" csv:"-"`
	CnpjBasico                           string `gorm:"index"`
	IdentificadorDeSocio                 string
	NomeSocio                            string `gorm:"index"`
	CnpjcpfDoSocio                       string `gorm:"index"`
	CodigoQualificacaoSocio              string
	DataEntradaSociedade                 string
	CodigoPais                           string
	CpfRepresentanteLegal                string
	NomeRepresentante                    string
	CodigoQualificacaoRepresentanteLegal string
	FaixaEteriaSocio                     string
}
