package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/andrersp/go-etl-receita-federal/src/config"
	"github.com/andrersp/go-etl-receita-federal/src/download"
	"github.com/andrersp/go-etl-receita-federal/src/extract"
)

func ListFiles(fileNameBase string) []string {
	folder := fmt.Sprintf("%s/*%s*.csv", config.DownloadFolder, fileNameBase)
	files, err := filepath.Glob(folder)
	if err != nil {
		log.Fatal(err)
	}
	return files

}

func main() {

	err := config.SetInitialConfig()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	download.StartDownload()

	for _, file := range config.FilesToFind {
		files := ListFiles(file)

		if len(files) > 0 {
			switch file {
			case "Cnae":
				wg.Add(1)
				go extract.StartStractComplementos(files, file, &wg)
			case "Estabelecimento":
				wg.Add(1)
				go extract.StartStractEstabelecimento(files, &wg)
			case "Empresa":
				wg.Add(1)
				go extract.StartStractEmpresa(files, &wg)
			case "Motivo":
				wg.Add(1)
				go extract.StartStractComplementos(files, file, &wg)
			case "Pais":
				wg.Add(1)
				go extract.StartStractComplementos(files, file, &wg)
			case "Municipio":
				wg.Add(1)
				go extract.StartStractComplementos(files, file, &wg)
			case "Natureza":
				wg.Add(1)
				go extract.StartStractComplementos(files, file, &wg)
			case "Qualificacoes":
				wg.Add(1)
				go extract.StartStractComplementos(files, file, &wg)
			case "Socio":
				wg.Add(1)
				go extract.StartStractSocio(files, &wg)
			case "Simples":
				wg.Add(1)
				go extract.StartStractSimples(files, &wg)

			}

		}

	}
	wg.Wait()
	log.Print("Extracao Finalizada")

}
