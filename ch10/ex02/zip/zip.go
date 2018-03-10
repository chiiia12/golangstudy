package zip

import (
	"log"
	"path/filepath"
	"os"
	"io"
	"io/ioutil"
	"archive/zip"
	"../unarchive"
)

type ZipUnArchiver struct {
	inputDir  string
	outputDir string
}

func init() {
	//登録する
	unarchive.Register("zip")
}

func (z *ZipUnArchiver) UnArchive() {
	unZip(z.inputDir, z.outputDir)
}

func unZip(input, output string) error {
	r, err := zip.OpenReader(input)
	if err != nil {
		log.Println(err)
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		if f.FileInfo().IsDir() {
			path := filepath.Join(output, f.Name)
			os.MkdirAll(path, f.Mode())
		} else {
			buf := make([]byte, f.UncompressedSize64)
			_, err := io.ReadFull(rc, buf)
			if err != nil {
				return err
			}
			path := filepath.Join(output, f.Name)
			if err = ioutil.WriteFile(path, buf, f.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}
