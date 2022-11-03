package scrap

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/andrersp/go-etl-receita-federal/src/config"
)

var (
	ErrIpNotFound = errors.New("ip address not found")
)

func FindDownloadsIpAddress() (address string, err error) {

	resp, err := soup.Get("https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/cadastros/consultas/dados-publicos-cnpj")
	if err != nil {
		return
	}

	doc := soup.HTMLParse(resp)
	links := doc.Find("a", "class", "external-link")

	err = links.Error

	if err != nil {
		return
	}

	link := links.Attrs()["href"]

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	address = re.FindString(link)

	if address == "" {
		err = ErrIpNotFound
		return
	}
	return

}

func FindFileName(baseName, url string) string {
	m := regexp.MustCompile(fmt.Sprintf(`%s.*.zip`, baseName))
	res := m.FindString(url)
	return res
}

func GetLinksAndFileNames() (links [][2]string, err error) {

	resp, err := soup.Get("https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/cadastros/consultas/dados-publicos-cnpj")
	if err != nil {
		return
	}

	doc := soup.HTMLParse(resp)
	pageLinks := doc.FindAll("a", "class", "external-link")

	ipv4Regex := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	for _, link := range pageLinks {

		href := link.Attrs()["href"]
		address := ipv4Regex.MatchString(href)

		if address {

			for _, file := range config.FilesToFind {
				if strings.Contains(href, file) {
					fileName := FindFileName(file, href)
					links = append(links, [2]string{fileName, href})
				}
			}
		}

	}

	return

}
