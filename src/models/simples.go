package models

import (
	"time"
)

type ModelSimples struct {
	ID                  uint64 `gorm:"primaryKey"`
	CnpjBasico          string
	OpcaoSimples        string
	DataOpcaoSimples    time.Time
	DataExclusaoSimples time.Time
	OpcaoMei            time.Time
	DataOpcaoMei        time.Time
	DataExclusaoMei     time.Time
}
