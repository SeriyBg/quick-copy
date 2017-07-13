package copy

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const dstDirectory = "../testdata/destination/"
const srcDirectory = "../testdata/source/"

func TestMain(m *testing.M) {
	setUp()
	run := m.Run()
	cleanUp()
	os.Exit(run)
}

func Test_ShouldCopyFile(t *testing.T) {
	const fileName = "test.txt"
	const srcFilename = srcDirectory + fileName

	err := file(srcFilename, dstDirectory)
	assert.NoError(t, err)

	dst, err := os.Open(dstDirectory)
	defer dst.Close()
	assert.False(t, os.IsNotExist(err))
	assert.NoError(t, err)

	assertFileContentsAreEqual(srcFilename, fileName, t)
}

func Test_ShouldGetErrorIfCopyFailed(t *testing.T) {
	const srcFilename = "../testdata/source/notExist"
	err := file(srcFilename, dstDirectory)
	assert.True(t, os.IsNotExist(err))
}

func Test_ShouldCopyDirectory(t *testing.T) {
	const directoryName = "empty_dir"
	const srcDir = srcDirectory + directoryName

	err := directory(srcDir, dstDirectory)
	assert.NoError(t, err)

	stat, err := os.Stat(dstDirectory + directoryName)
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}

func Test_CopyFilesInDirectory(t *testing.T) {
	const directoryName = "file_dir/"
	const srcDir = srcDirectory + directoryName

	err := directory(srcDir, dstDirectory)
	assert.NoError(t, err)

	directory, err := os.Open(srcDir)
	assert.NoError(t, err)

	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		assertFileContentsAreEqual(srcDir+obj.Name(), directoryName+obj.Name(), t)
	}
}

func Test_CopyDirectoryRecursive(t *testing.T) {
	const directoryName = "complex_dir/"
	const nestedDirectoryName = "nested_dir/"
	const srcDir = srcDirectory + directoryName

	err := directory(srcDir, dstDirectory)
	assert.NoError(t, err)

	directory, err := os.Open(srcDir)
	assert.NoError(t, err)

	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		if !obj.IsDir() {
			assertFileContentsAreEqual(srcDir+obj.Name(), directoryName+obj.Name(), t)
		}
	}
	assertFileContentsAreEqual(srcDir+nestedDirectoryName+"test1.txt",
		directoryName+nestedDirectoryName+"test1.txt", t)
}

func assertFileContentsAreEqual(srcFilename string, fileName string, t *testing.T) {
	srcContent, err := ioutil.ReadFile(srcFilename)
	assert.NoError(t, err)
	dstContent, err := ioutil.ReadFile(dstDirectory + fileName)
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(srcContent, dstContent))
}

func setUp() {
	os.MkdirAll(dstDirectory, os.ModePerm)
}

func cleanUp() {
	os.RemoveAll(dstDirectory)
}
