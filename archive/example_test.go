package archive_test

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/b4fun/battery/archive"
)

func ExampleCreateZipArchive_Compress() {
	c := &archive.CreateZipArchive{
		SourceDir: "testdata",
	}

	var b bytes.Buffer
	if err := c.CompressTo(&b); err != nil {
		log.Fatal(err)
	}

	compressedSource := bytes.NewReader(b.Bytes())
	compressed, err := zip.NewReader(compressedSource, compressedSource.Size())
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range compressed.File {
		fmt.Println(file.Name)
	}
	// Output:
	// /singlefile/test.txt
}

func ExampleCreateZipArchive_BasePath() {
	c := &archive.CreateZipArchive{
		SourceDir: "testdata",
		BasePath:  "testbasepath",
	}

	var b bytes.Buffer
	if err := c.CompressTo(&b); err != nil {
		log.Fatal(err)
	}

	compressedSource := bytes.NewReader(b.Bytes())
	compressed, err := zip.NewReader(compressedSource, compressedSource.Size())
	if err != nil {
		log.Fatal(err)
	}
	if len(compressed.File) != 1 {
		log.Fatalf("should contain test.txt, got: %v", compressed.File)
	}
	for _, file := range compressed.File {
		fmt.Println(file.Name)
	}
	// Output:
	// testbasepath/singlefile/test.txt
}

func ExampleCreateZipArchive_SkipFile() {
	c := &archive.CreateZipArchive{
		SourceDir: "./testdata",
		FilterFile: func(path string, info os.FileInfo) bool {
			return false
		},
	}

	var b bytes.Buffer
	if err := c.CompressTo(&b); err != nil {
		log.Fatal(err)
	}

	compressedSource := bytes.NewReader(b.Bytes())
	compressed, err := zip.NewReader(compressedSource, compressedSource.Size())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(compressed.File))
	// Output:
	// 0
}
