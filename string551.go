package string551

import "unsafe"

func StringToBytes(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&src))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func Join(src ...string) string {
	length := 0
	point := 0
	for i := 0; i < len(src); i++ {
		length += len(StringToBytes(src[i]))
	}
	ret := make([]byte, length)
	for i := 0; i < len(src); i++ {
		byt := StringToBytes(src[i])
		for j := 0; j < len(byt); j++ {
			ret[point] = byt[j]
			point++
		}
	}
	return BytesToString(ret)
}

func CamelCase(src string) string {
	bytes := StringToBytes(src)
	ret := make([]byte, 0, len(bytes))
	lowLine := true
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 0x5F {
			lowLine = true
			continue
		}
		if lowLine {
			ret = append(ret, bytes[i]-0x20)
			lowLine = false
		} else {
			ret = append(ret, bytes[i])
		}
	}

	return BytesToString(ret)
}

func SnakeCase(src string) string {
	bytes := StringToBytes(src)

	ret := make([]byte, 0, len(bytes)+10)
	for i := 0; i < len(bytes); i++ {
		if 0x41 <= bytes[i] && bytes[i] <= 0x5A { // 0x41 - 0x5A => "A" - "Z"
			if i != 0 {
				ret = append(ret, 0x5F) // 0x5f => _
			}
			ret = append(ret, bytes[i]+0x20) // "A" => "a", "B" => "b"
		} else {
			ret = append(ret, bytes[i])
		}
	}

	return BytesToString(ret)
}
