package archive

import (
	"io/ioutil"
	"testing"
)

func TestCreateZipArchive_Validate(t *testing.T) {
	c := &CreateZipArchive{}

	if err := c.CompressTo(ioutil.Discard); err == nil {
		t.Errorf("should validate parameters")
	}
}
