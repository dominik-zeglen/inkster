//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...

package schema

import (
	"bytes"
	"sort"
)

func String() string {
	assets := AssetNames()
	buf := bytes.Buffer{}

	sort.Strings(assets)

	for _, name := range assets {
		b, _ := Asset(name)
		buf.Write(b)

		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
