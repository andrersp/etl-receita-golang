package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/andrersp/go-etl-receita-federal/src/config"
)

func Unzip(src, destination string) ([]string, error) {

	var fileNames []string

	inputFile := fmt.Sprintf("%s/%s", config.DownloadFolder, src)

	r, err := zip.OpenReader(inputFile)

	if err != nil {
		return fileNames, err
	}

	defer r.Close()

	for _, f := range r.File {
		outName := strings.Replace(src, "zip", "csv", -1)
		fpath := filepath.Join(destination, outName)

		// Check for invalid file paths
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return fileNames, fmt.Errorf("%s is an illegal filepath", fpath)
		}

		fileNames = append(fileNames, fpath)

		if f.FileInfo().IsDir() {
			// Create New Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return fileNames, err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())

		if err != nil {
			return fileNames, err
		}

		rc, err := f.Open()

		if err != nil {
			return fileNames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return fileNames, err
		}

	}
	os.Remove(inputFile)

	return fileNames, nil
}
