package zattr

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/gobwas/glob"
)

type TransformFunc func(*zip.File) error
type FileHeaderMapper func(*zip.FileHeader) error

func OpenZip(name string) *zip.ReadCloser {
	reader, err := zip.OpenReader(name)
	if err != nil {
		log.Fatal(err)
	}
	return reader
}

func PrintFile(f *zip.File) error {
	fmt.Printf("%s\t%#o\n", f.Name, f.ExternalAttrs)
	return nil
}

func AttrChanger(g glob.Glob, mode string) FileHeaderMapper {
	modeNumber, _ := strconv.ParseUint(mode, 0, 32)

	return FileHeaderMatcher(g, func(f *zip.FileHeader) error {
		newModeNumber := (uint32)(modeNumber | syscall.S_IFREG)
		fmt.Printf("%s\t%#o -> %#o\n", f.Name, f.ExternalAttrs, newModeNumber)
		f.ExternalAttrs = newModeNumber
		return nil
	})
}

func CopyZipPath(from string, dest string, mapper FileHeaderMapper) error {
	reader := OpenZip(from)
	defer reader.Close()

	tmp, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.Close()

	writer := zip.NewWriter(tmp)
	defer writer.Close()

	return CopyZip(reader, writer, mapper)
}

func CopyZip(file *zip.ReadCloser, dest *zip.Writer, mapper FileHeaderMapper) error {
	transform := func(f *zip.File) error {
		return AddFile(f, dest, mapper)
	}

	return TransformZip(file, transform)
}

func OpenTemp(prefix string) *os.File {
	tmp, err := ioutil.TempFile("", prefix)
	if err != nil {
		log.Fatal(err)
	}
	return tmp
}

func AddFile(file *zip.File, archive *zip.Writer, mapper FileHeaderMapper) error {
	// duplicate struct
	var header zip.FileHeader
	header = file.FileHeader

	mapper(&header)

	writer, err := archive.CreateHeader(&header)
	if err != nil {
		log.Fatal(err)
	}

	reader, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	_, err = io.Copy(writer, reader)
	return err
}

func TransformZip(file *zip.ReadCloser, transform TransformFunc) error {
	for _, f := range file.File {
		transform(f)
	}

	return nil
}

func FileHeaderMatcher(g glob.Glob, mapper FileHeaderMapper) FileHeaderMapper {
	return func(f *zip.FileHeader) error {
		if g.Match(f.Name) {
			return mapper(f)
		} else {
			return nil
		}
	}
}
