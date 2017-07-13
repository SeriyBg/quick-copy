package copy

import (
	"io"
	"os"
	"strings"
)

const pathSeparator = string(os.PathSeparator)

func file(src string, dstDirectory string) (err error) {
	source, err := os.Open(src)
	if err != nil {
		return
	}
	defer source.Close()

	stat, err := source.Stat()
	dstDirectory = dstName(dstDirectory, stat)

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
	dstDirectory = dstName(dstDirectory, stat)
	err = os.MkdirAll(dstDirectory, stat.Mode())
	dir, err := os.Open(src)

	objects, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, obj := range objects {
		if obj.IsDir() {
			err = directory(src+pathSeparator+obj.Name()+pathSeparator, dstDirectory)
		} else {
			err = file(src+obj.Name(), dstDirectory)
		}
		if err != nil {
			return
		}
	}
	return
}

func dstName(dstDirectory string, stat os.FileInfo) string {
	if strings.HasSuffix(dstDirectory, pathSeparator) {
		dstDirectory = dstDirectory + stat.Name()
	} else {
		dstDirectory = dstDirectory + pathSeparator + stat.Name()
	}
	return dstDirectory
}
