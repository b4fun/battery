package fs_test

import (
	"log"
	"os"

	"github.com/b4fun/battery/fs"
	"github.com/b4fun/battery/fs/testdata"
)

func ExampleCopyAllToDir() {
	tempDir, err := os.MkdirTemp("", "battery-fs")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	err = fs.CopyAllToDir(testdata.Root, tempDir)
	if err != nil {
		log.Fatal(err)
	}
}
