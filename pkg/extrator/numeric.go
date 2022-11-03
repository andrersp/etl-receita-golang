package extrator

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/andrersp/go-etl-receita-federal/src/config"
)

type ConvertReflector func(string) (uint64, error)

func RetryConvert(effector ConvertReflector, retries int) ConvertReflector {

	return func(s string) (uint64, error) {

		for r := 0; ; r++ {
			intNum, err := effector(s)

			if err == nil || r == retries {
				return intNum, nil
			}

			log.Printf("Err on extract %s, retrying", s)
			s = ExtractNumbersFromString(s)
		}
	}
}

func ExtractNumbersFromString(v string) string {

	if v != "" {
		return config.RegexNumberCompile.ReplaceAllString(v, "")
	}
	return v
}

func ParseStringToInt64(stringNumber string) (uint64, error) {

	return strconv.ParseUint(stringNumber, 10, 64)

}

func ConvertStringToUint64(v string) uint64 {

	retry := RetryConvert(ParseStringToInt64, 3)

	num, err := retry(v)
	if err != nil {
		return 0
	}
	return num
}

func TextToStringDate(date string) string {

	if len(date) < 8 {
		return ""
	}
	return fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:8])
}

func CnaeSecundaToArray(cnaeSecundaria string) (cnaesArray []string) {

	cnaes := strings.Split(cnaeSecundaria, ",")

	for _, cnae := range cnaes {
		if cnae != "" && cnae != " " {
			cnaesArray = append(cnaesArray, cnae)
		}

	}
	return

}
