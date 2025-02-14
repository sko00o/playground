package conv

import (
	"reflect"
	"unsafe"
)

type noMagic struct{}

func (noMagic) s2b(s string) (b []byte) { return []byte(s) }
func (noMagic) b2s(b []byte) string     { return string(b) }

type v1 struct{}

func (v1) s2b(s string) (b []byte) {
	strh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh.Data = strh.Data
	sh.Len = strh.Len
	sh.Cap = strh.Len
	return b
}

func (v1) b2s(b []byte) string {
	// also see in https://cs.opensource.google/go/go/+/refs/tags/go1.19.9:src/strings/builder.go;l=48
	return *(*string)(unsafe.Pointer(&b))
}

type v2 struct{}

func (v2) s2b(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func (v2) b2s(b []byte) string {
	// also see in https://cs.opensource.google/go/go/+/refs/tags/go1.20:src/strings/builder.go;l=48
	return unsafe.String(unsafe.SliceData(b), len(b))
}
