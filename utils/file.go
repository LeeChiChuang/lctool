package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/logrusorgru/aurora"
)

func WriteStringToFile(content, path string) error {
	fp, err := CreateIfNotExist(path)
	if err != nil {
		return err
	}

	defer fp.Close()

	write := bufio.NewWriter(fp)
	_, _ = write.WriteString(content)
	_ = write.Flush()

	return nil
}

const (
	NL = "\n"
)

func CreateIfNotExist(file string) (*os.File, error) {
	fmt.Println(file)
	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s already exist", file)
	}

	return os.Create(file)
}

func RemoveIfExist(filename string) error {
	if !FileExists(filename) {
		return nil
	}

	return os.Remove(filename)
}

func RemoveOrQuit(filename string) error {
	if !FileExists(filename) {
		return nil
	}

	fmt.Printf("%s exists, overwrite it?\nEnter to overwrite or Ctrl-C to cancel...",
		aurora.BgRed(aurora.Bold(filename)))
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	return os.Remove(filename)
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func FileNameWithoutExt(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}

func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}
