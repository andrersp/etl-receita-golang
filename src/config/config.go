package config

import (
	"os"
	"regexp"
)

var (
	DownloadFolder = "data/Downloads"
	OutputFolder   = "data/Outputs"
	DownloadAdrees = ""
	FilesToFind    = []string{"Socio", "Estabelecimento", "Empresa", "Cnae", "Simples", "Pais", "Motivo", "Municipio", "Natureza", "Qualificacoes"}

	RegexNumberCompile regexp.Regexp
)

func createFolders() error {

	err := os.MkdirAll(DownloadFolder, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(OutputFolder, os.ModePerm)
	if err != nil {
		return err
	}
	return nil

}

func SetInitialConfig() error {

	RegexNumberCompile = *regexp.MustCompile("[^0-9]+")

	if err := createFolders(); err != nil {
		return err
	}
	createCSVHeaders()

	return nil

}
