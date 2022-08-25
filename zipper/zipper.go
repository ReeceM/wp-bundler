package zipper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Zipper struct {
	Writer     zip.Writer
	Archive    os.File
	ignoreList []string
}

// Write a file set
func (m *Zipper) Write(dir string) {

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := m.Writer.Create(path)
		if err != nil {
			return err
		}

		fmt.Printf("Adding: %#v\n", path)
		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}

	err := filepath.Walk("dist", walker)

	if err != nil {
		panic(err)
	}
}

func (m Zipper) CloseAll() {
	m.Writer.Close()
	m.Archive.Close()
}

func (m *Zipper) create() {
	fmt.Println("Opening archive")
	archive, err := os.Create("archive.zip")
	m.Archive = *archive

	if err != nil {
		panic(err)
	}
	defer archive.Close()

	m.Writer = *zip.NewWriter(archive)
	defer m.Writer.Close()
}

var ZipWriter Zipper

func Init() {
	ZipWriter.create()
}
