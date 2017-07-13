package copy

import (
	"os"
	"io"
)

func file(src string, dstDirectory string) (err error) {
	source, err := os.Open(src)
	if err != nil {
		return
	}
	defer source.Close()

	stat, err := source.Stat()
	dstDirectory = dstDirectory + stat.Name()

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
	os.MkdirAll(dstDirectory+ stat.Name(), stat.Mode())
	return
}