package models

type ModelEmpresa struct {
	ID                        int64
	CnpjBasico                string `gorm:"index"`
	RazaoSocial               string `gorm:"index"`
	NaturezaJuridica          string
	QualificacaoResponsavel   string
	CapitalSocial             string
	PorteEmpresa              string
	EnteFederativoResponsavel string
}
