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

func readSocioCSV(filePath string, csvOut *csv.Writer) {

	f, _ := os.Open(filePath)
	defer f.Close()

	var outputRows [][]string

	r := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(f))
	r.Comma = ';'
	r.LazyQuotes = true

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
		identificadorDeSocio := record[1]
		momeSocio := record[2]
		cnpjcpfDoSocio := extrator.ExtractNumbersFromString(record[3])
		codigoQualificacaoSocio := record[4]
		dataEntradaSociedade := extrator.TextToStringDate(record[5])
		codigoPais := record[6]
		cpfRepresentanteLega := extrator.ExtractNumbersFromString(record[7])
		nomeRepresentante := record[8]
		codigoQualificacaoRepresentanteLegal := record[9]
		faixaEteriaSocio := record[10]

		outputRows = append(outputRows, []string{
			cnpjBasico,
			identificadorDeSocio,
			momeSocio,
			cnpjcpfDoSocio,
			codigoQualificacaoSocio,
			dataEntradaSociedade,
			codigoPais,
			cpfRepresentanteLega,
			nomeRepresentante,
			codigoQualificacaoRepresentanteLegal,
			faixaEteriaSocio,
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

func StartStractSocio(files []string, wg *sync.WaitGroup) {

	outputFile := fmt.Sprintf("%s/socios_out.csv", config.OutputFolder)

	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	csvWriter := csv.NewWriter(outFile)
	csvWriter.Write(config.HeaderSociosInput)

	for _, file := range files {
		readSocioCSV(file, csvWriter)
	}
	csvWriter.Flush()
	wg.Done()
}
