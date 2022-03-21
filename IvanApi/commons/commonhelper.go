package commons

import "unsafe"

//byte to string
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
