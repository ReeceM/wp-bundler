package zipper

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"wp-bundler/config"
)

type Zipper struct {
	Writer     zip.Writer
	Archive    os.File
	IgnoreList []string
	Close      func()
}

// Write a file set
func (m *Zipper) Write(dir string) {

	ignores, err := readLines(_config.Bundler.IgnoreFile)
	if err != nil {
		panic(err)
	}

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		fmt.Println("name", info.Name())

		if err != nil {
			fmt.Println(err)
			return err
		}

		if stringInSlice(info.Name(), ignores) {
			fmt.Printf("Ignoring: %#v\n", path)
			return nil
		}

		if stringInSlice(path, ignores) {
			fmt.Printf("Ignoring: %#v\n", path)
			return nil
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

	err = filepath.Walk(dir, walker)

	if err != nil {
		panic(err)
	}
}

func (m *Zipper) create() {
	fmt.Println("Opening archive")
	archive, err := os.Create("archive.zip")
	m.Archive = *archive

	if err != nil {
		panic(err)
	}

	m.Writer = *zip.NewWriter(archive)

	m.Close = func() {
		fmt.Println("Closed archive")
		m.Writer.Close()
		m.Archive.Close()
	}
}

var ZipWriter Zipper
var _config config.TomlConfig

func Init(config config.TomlConfig) {
	ZipWriter.create()
	_config = config
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
// https://stackoverflow.com/a/18479916/3778963
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// https://stackoverflow.com/a/15323988/3778963
func stringInSlice(needle string, list []string) bool {
	for _, b := range list {
		if b == needle {
			return true
		}

		result, err := regexp.MatchString(fmt.Sprintf(`%s`, b), needle)

		if err != nil {
			return false
		}

		if result {
			return true
		}
	}
	return false
}
