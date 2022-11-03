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

func readEstabelecimentoCSV(filePath string, csvOut *csv.Writer, cnaeOut *csv.Writer) {

	f, _ := os.Open(filePath)
	defer f.Close()

	r := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(f))
	r.Comma = ';'
	r.LazyQuotes = true

	var outputRows [][]string

	var cnaesSecundarias [][]string

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
		cnpjOrdem := record[1]
		cnpjDv := record[2]
		identificadorMatrizFilial := record[3]
		nomeFantasia := record[4]
		situacaoCadastral := record[5]
		dataSituacaoCadastral := extrator.TextToStringDate(record[6])
		motivoSituacaoCadastral := record[7]
		nomeCidadeExterior := record[8]
		codPais := record[9]
		dataInicioAtividade := extrator.TextToStringDate(record[10])
		cnaeFiscal := record[11]
		cnaeSecundaria := record[12]
		tipoLogradouro := record[13]
		logradouro := record[14]
		numero := record[15]
		complemento := record[16]
		bairro := record[17]
		cep := record[18]
		uf := record[19]
		codigo_municipio := record[20]
		ddd_1 := record[21]
		telefone_1 := record[22]
		ddd_2 := record[23]
		telefone_2 := record[24]
		ddd_fax := record[25]
		fax := record[26]
		correioEletronico := record[27]
		situacaoEspecial := record[28]
		dataSituacaoEspecial := extrator.TextToStringDate(record[29])

		outputRows = append(outputRows, []string{
			cnpjBasico,
			cnpjOrdem,
			cnpjDv,
			identificadorMatrizFilial,
			nomeFantasia,
			situacaoCadastral,
			dataSituacaoCadastral,
			motivoSituacaoCadastral,
			nomeCidadeExterior,
			codPais,
			dataInicioAtividade,
			cnaeFiscal,
			tipoLogradouro,
			logradouro,
			numero,
			complemento,
			bairro,
			cep,
			uf,
			codigo_municipio,
			ddd_1,
			telefone_1,
			ddd_2,
			telefone_2,
			ddd_fax,
			fax,
			correioEletronico,
			situacaoEspecial,
			dataSituacaoEspecial,
		})

		cnaesArray := extrator.CnaeSecundaToArray(cnaeSecundaria)

		for _, cnae := range cnaesArray {
			cnaesSecundarias = append(cnaesSecundarias,
				[]string{
					cnpjBasico, cnpjOrdem, cnpjDv, cnae,
				},
			)

		}

		if i == 500 {
			csvOut.WriteAll(outputRows)
			cnaeOut.WriteAll(cnaesSecundarias)
			outputRows = nil
			cnaesSecundarias = nil
			i = 0

		}
	}

	if i > 0 {
		csvOut.WriteAll(outputRows)
		cnaeOut.WriteAll(cnaesSecundarias)

	}
}

func StartStractEstabelecimento(files []string, wg *sync.WaitGroup) {

	outputFile := fmt.Sprintf("%s/estabelecimento_out.csv", config.OutputFolder)
	cnaeSecundariaFile := fmt.Sprintf("%s/cnae_secundaria_out.csv", config.OutputFolder)

	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	outFileCnae, err := os.Create(cnaeSecundariaFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	defer outFileCnae.Close()

	csvWriter := csv.NewWriter(outFile)
	csvWriter.Write(config.HeaderEstabelecimentoInput)

	csvCnaeWriter := csv.NewWriter(outFileCnae)
	csvCnaeWriter.Write(config.HeaderCnaeSecundarua)

	for _, file := range files {
		readEstabelecimentoCSV(file, csvWriter, csvCnaeWriter)
	}
	csvWriter.Flush()
	csvCnaeWriter.Flush()
	wg.Done()
}
