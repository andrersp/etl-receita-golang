package models

type Cnae struct {
	ID        int64  `gorm:"primaryKey"`
	Codigo    string `gorm:"index"`
	Descricao string
}

func (Cnae) TableName() string {
	return "cnae"
}

type NaturezaJuridica struct {
	ID        int64  `gorm:"primaryKey"`
	Codigo    string `gorm:"index"`
	Descricao string
}

func (NaturezaJuridica) TableName() string {
	return "natureza_juridica"
}

type Pais struct {
	ID        int64  `gorm:"primaryKey"`
	Codigo    string `gorm:"index"`
	Descricao int
}

func (Pais) TableName() string {
	return "pais"
}

type Municipio struct {
	ID        int64  `gorm:"primaryKey"`
	Codigo    string `gorm:"index"`
	Descricao int
}

func (Municipio) TableName() string {
	return "municipio"
}

type QualificacaoSocio struct {
	ID        int64  `gorm:"primaryKey"`
	Codigo    string `gorm:"index"`
	Descricao int
}

func (QualificacaoSocio) TableName() string {
	return "qualificacao_socio"
}

type Motivo struct {
	ID        int64  `gorm:"primaryKey"`
	Codigo    string `gorm:"index"`
	Descricao int
}

func (Motivo) TableName() string {
	return "natureza_juridica"
}
