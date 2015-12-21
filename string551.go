package string551

import "unsafe"

func StringToBytes(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&src))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
