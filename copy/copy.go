package copy

import (
	"io"
	"os"
	"strings"
)

func file(src string, dstDirectory string) (err error) {
	source, err := os.Open(src)
	if err != nil {
		return
	}
	defer source.Close()

	stat, err := source.Stat()
	if strings.HasSuffix(dstDirectory, "/") {
		dstDirectory = dstDirectory + stat.Name()
	} else {
		dstDirectory = dstDirectory + "/" + stat.Name()
	}

	if err != nil {
		return
	}
	destination, err := os.Create(dstDirectory)

	if err != nil {
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)

	if err == nil {
		stat, err := os.Stat(src)
		if err != nil {
			os.Chmod(dstDirectory, stat.Mode())
		}
	}
	return
}

func directory(src string, dstDirectory string) (err error) {
	stat, err := os.Stat(src)
	if err != nil {
		return
	}
	copiedDirectory := dstDirectory + stat.Name()
	os.MkdirAll(copiedDirectory, stat.Mode())

	directory, err := os.Open(src)

	objects, err := directory.Readdir(-1)
	if err != nil {
		return
	}
	for _, obj := range objects {
		err = file(src + obj.Name(), copiedDirectory)
		if err != nil {
			return
		}
	}
	return
}
