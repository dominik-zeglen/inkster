package utils

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/inkster/core"
)

func TestPrintJson(t *testing.T) {
	t.Run("Print pretty JSON", func(t *testing.T) {
		directory := core.Directory{
			Name:        "Test",
			IsPublished: true,
		}

		result, err := PrintJSON(&directory)
		if err != nil {
			t.Fatal(err)
		}
		cupaloy.SnapshotT(t, result)
	})
}
