//go:build !gofuzz

package plist

import (
	"bytes"
	"testing"
)

func FuzzDecode(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		buf := bytes.NewReader(data)
		var obj interface{}
		_ = NewDecoder(buf).Decode(&obj)
	})
}
