package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andrersp/go-etl-receita-federal/src/config"
	"github.com/andrersp/go-etl-receita-federal/src/scrap"
	"github.com/andrersp/go-etl-receita-federal/src/unzip"
)

type Effector func(string, string, chan string) error

var count int

func Retry(effector Effector, retries int, delay time.Duration) Effector {

	return func(url, fileName string, ch chan string) error {

		for r := 0; ; r++ {
			err := effector(url, fileName, ch)

			if err == nil {
				return err
			}

			log.Printf("%s, retrying in %v", fileName, delay)

			time.Sleep(delay)

		}

	}
}

func downloadFile(url, filename string, ch chan string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("%s failed.", filename)

	}

	// cretate a empty file
	file, err := os.Create(fmt.Sprintf("%s/%s", config.DownloadFolder, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	// Write bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	ch <- filename
	return nil
}

func StartDownload() error {

	ch := make(chan string)

	urls, err := scrap.GetLinksAndFileNames()

	if err != nil {
		log.Fatal()
	}

	for _, url := range urls {
		link := url[1]
		fileName := url[0]
		retry := Retry(downloadFile, 5, time.Duration(2*time.Second))
		go retry(link, fileName, ch)
	}

	for range urls {
		fileName := <-ch

		_, err := unzip.Unzip(fileName, config.DownloadFolder)
		if err != nil {
			log.Fatal(err)
		}

	}

	return nil
}
