package string551

import "unsafe"

func StringToBytes(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&src))
}
