package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/andrersp/go-etl-receita-federal/src/config"
	"golang.org/x/text/encoding/charmap"
)

func readComplementosCSV(filePath string, csvOut *csv.Writer) {

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
		codigo := record[0]
		descricao := strings.Replace(record[1], ";", ",", -1)
		outputRows = append(outputRows, []string{codigo, descricao})

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

func StartStractComplementos(files []string, fileNameOutput string, wg *sync.WaitGroup) {

	outputFile := fmt.Sprintf("%s/%s_out.csv", config.OutputFolder, strings.ToLower(fileNameOutput))

	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	csvWriter := csv.NewWriter(outFile)
	csvWriter.Write(config.HeaderComplementarInput)

	for _, file := range files {
		readComplementosCSV(file, csvWriter)
	}
	csvWriter.Flush()
	wg.Done()
}
