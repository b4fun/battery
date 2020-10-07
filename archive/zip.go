package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FileFilterFunc is the type of the function called for each file
// visited by archiver. The path argument contains the source dir argument to compress.
// Returns value indicates the archiver to compress the file or not.
type FileFilterFunc func(path string, info os.FileInfo) bool

// CreateZipArchive creates a zip archive from a path.
type CreateZipArchive struct {
	// SourceDir sets the source dir to compress. Required.
	SourceDir string

	// BasePath sets the base path for the archive files. Defaults to no base path.
	BasePath string

	// FilterFile provides an optional callback to filter a file.
	FilterFile FileFilterFunc
}

// CompressTo compresses to a writer.
func (c *CreateZipArchive) CompressTo(dest io.Writer) error {
	zipWriter := zip.NewWriter(dest)

	walkErr := filepath.Walk(c.SourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// continue
			return nil
		}

		if c.FilterFile != nil {
			if !c.FilterFile(path, info) {
				// skip this file
				return nil
			}
		}

		compressPath := filepath.Join(
			c.BasePath,
			strings.TrimPrefix(path, c.SourceDir),
		)
		zipFile, err := zipWriter.Create(compressPath)
		if err != nil {
			return fmt.Errorf("create zip file %s (%s): %w", compressPath, path, err)
		}

		sourceFile, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("open source file %s: %w", path, err)
		}

		_, err = io.Copy(zipFile, sourceFile)
		if err != nil {
			return fmt.Errorf("compress %s to zip file %s: %w", path, compressPath, err)
		}

		if err := sourceFile.Close(); err != nil {
			return fmt.Errorf("close source file %s: %w", path, err)
		}

		return nil
	})
	if walkErr != nil {
		return fmt.Errorf("walk %s: %w", c.SourceDir, walkErr)
	}

	if err := zipWriter.Close(); err != nil {
		return fmt.Errorf("compress %s: %w", c.SourceDir, err)
	}

	return nil
}
