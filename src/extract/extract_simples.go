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

func readSimplesCSV(filePath string, csvOut *csv.Writer) {

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
		opcaoSimples := record[1]
		dataOpcaoSimples := extrator.TextToStringDate(record[2])
		dataExclusaoSimples := extrator.TextToStringDate(record[3])
		opcaoMei := record[4]
		dataOpcaoMei := extrator.TextToStringDate(record[5])
		dataExclusaoMei := extrator.TextToStringDate(record[6])

		outputRows = append(outputRows, []string{
			cnpjBasico,
			opcaoSimples,
			dataOpcaoSimples,
			dataExclusaoSimples,
			opcaoMei,
			dataOpcaoMei,
			dataExclusaoMei})

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

func StartStractSimples(files []string, wg *sync.WaitGroup) {

	outputFile := fmt.Sprintf("%s/simples_out.csv", config.OutputFolder)

	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	csvWriter := csv.NewWriter(outFile)
	csvWriter.Write(config.HeaderSimplesInput)

	for _, file := range files {
		readSimplesCSV(file, csvWriter)
	}
	csvWriter.Flush()
	wg.Done()
}
