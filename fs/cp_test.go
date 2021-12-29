package fs

import (
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/b4fun/battery/fs/testdata"
)

func TestCopyAllToDir(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "battery-fs")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tempDir)

	err = CopyAllToDir(testdata.Root, tempDir)
	if err != nil {
		t.Error(err)
	}

	tempDirFS := os.DirFS(tempDir)

	// TODO: add compensieve compare checks for all files
	cmpErr := fs.WalkDir(testdata.Root, ".", func(srcPath string, srcEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if srcEntry.IsDir() {
			dest, err := fs.Stat(tempDirFS, srcPath)
			if err != nil {
				return err
			}
			if !dest.IsDir() {
				return fmt.Errorf("expected %s to be a dir", srcPath)
			}
			return nil
		}

		srcStat, err := srcEntry.Info()
		if err != nil {
			return err
		}

		destFile, err := tempDirFS.Open(srcPath)
		if err != nil {
			return err
		}
		destStat, err := destFile.Stat()
		if err != nil {
			return err
		}
		if destStat.Size() != srcStat.Size() {
			return fmt.Errorf("expected %s to have size %d, got %d", srcPath, srcStat.Size(), destStat.Size())
		}

		return nil
	})
	if cmpErr != nil {
		t.Error(cmpErr)
	}
}
