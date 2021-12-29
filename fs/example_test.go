package fs_test

import (
	"log"
	"os"
	"path/filepath"

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

	walkErr := filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		log.Printf("%s %s", path, info.Mode())
		return nil
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
