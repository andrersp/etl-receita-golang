package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/andrersp/go-etl-receita-federal/pkg/extrator"
	"github.com/andrersp/go-etl-receita-federal/src/config"
	"golang.org/x/text/encoding/charmap"
)

func readEmpresaCSV(filePath string, csvOut *csv.Writer) {

	f, _ := os.Open(filePath)
	defer f.Close()

	r := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(f))
	r.Comma = ';'
	r.LazyQuotes = true

	var outputRows [][]string
	i := 0

	for {
		i++

		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)

		}
		cnpjBasico := record[0]
		razaoSocial := record[1]
		naturezaJuridica := record[2]
		qualificacaoResponsavel := record[3]
		capitalSocial := extrator.ExtractNumbersFromString(record[4])
		porteEmpresa := record[5]
		enteFederativoResponsavel := record[6]

		outputRows = append(outputRows, []string{
			cnpjBasico,
			razaoSocial,
			naturezaJuridica,
			qualificacaoResponsavel,
			capitalSocial,
			porteEmpresa,
			enteFederativoResponsavel,
		})

		if i == 500 {
			csvOut.WriteAll(outputRows)
			outputRows = nil
			i = 0

		}
	}

	if i > 0 {
		csvOut.WriteAll(outputRows)

	}
}

func StartStractEmpresa(files []string, wg *sync.WaitGroup) {

	outputFile := fmt.Sprintf("%s/empresas_out.csv", config.OutputFolder)

	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	csvWriter := csv.NewWriter(outFile)
	csvWriter.Write(config.HeaderEmpresaInput)

	for _, file := range files {
		readEmpresaCSV(file, csvWriter)
	}
	csvWriter.Flush()
	wg.Done()
}
