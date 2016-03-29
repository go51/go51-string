package string551

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unsafe"
)

var byteList []byte = StringToBytes("0123456789abcdef")

func StringToBytes(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&src))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func HexBytesToString(hexBytes []byte) string {
	ret := make([]byte, 0, len(hexBytes)*2)
	for i := 0; i < len(hexBytes); i++ {
		ret = append(ret, byteList[hexBytes[i]/16])
		ret = append(ret, byteList[hexBytes[i]%16])
	}

	return BytesToString(ret)
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
		if 0x41 <= bytes[i] && bytes[i] <= 0x5A {
			// 0x41 - 0x5A => "A" - "Z"
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

func Replace(src, old, new string) string {
	if src == old {
		return new
	}
	if new == "" {
		src = strings.Trim(src, old)
	}

	srcBytes := StringToBytes(src)
	oldBytes := StringToBytes(old)
	newBytes := StringToBytes(new)

	if len(srcBytes) == 0 {
		return src
	}
	if len(srcBytes) < len(oldBytes) {
		return src
	}

	retByte := make([]byte, 0, len(srcBytes))

	i := 0
	for i = 0; i <= len(srcBytes)-len(oldBytes); i++ {
		if srcBytes[i] == oldBytes[0] {
			cut := len(oldBytes)
			if BytesToString(oldBytes) == BytesToString(((srcBytes[i:])[:cut])) {
				retByte = append(retByte, newBytes...)
				i += len(oldBytes) - 1
			} else {
				retByte = append(retByte, srcBytes[i])
			}
		} else {
			retByte = append(retByte, srcBytes[i])
		}
	}
	retByte = append(retByte, srcBytes[i:]...)

	return BytesToString(retByte)

}

func Split(src, separate string) []string {
	return strings.Split(src, separate)
}

func Right(src string, length int) string {
	if len(src) < length {
		panic(errors.New("切り取り指定のバイト数が文字列長を超えました。"))
		return ""
	}

	return src[len(src)-length : len(src)]
}

func Lower(src string) string {
	bytes := StringToBytes(src)
	ret := make([]byte, 0, len(bytes))
	for i := 0; i < len(bytes); i++ {
		if 0x41 <= bytes[i] && bytes[i] <= 0x5A {
			// 0x41 - 0x5A => "A" - "Z"
			ret = append(ret, bytes[i]+0x20) // 0x61 - 0x7A => "a" - "z"
		} else {
			ret = append(ret, bytes[i])
		}
	}
	return BytesToString(ret)
}

func RightRune(src string, length int, suffix ...string) string {

	rSrc := []rune(src)
	if len(rSrc) <= length {
		return src
	}

	ret := ""
	for i := 0; i < length; i++ {
		ret = Join(ret, string(rSrc[i]))
	}

	if len(suffix) == 0 {
		return ret
	} else {
		return ret + suffix[0]
	}
}

func UrlEncode(src string) string {
	result := ""

	for _, c := range src {
		if c <= 0x7f { // single byte
			result += fmt.Sprintf("%%%X", c)
		} else if c > 0x1fffff { // quaternary byte
			result += fmt.Sprintf("%%%X%%%X%%%X%%%X",
				0xf0+((c&0x1c0000)>>18),
				0x80+((c&0x3f000)>>12),
				0x80+((c&0xfc0)>>6),
				0x80+(c&0x3f),
			)
		} else if c > 0x7ff { // triple byte
			result += fmt.Sprintf("%%%X%%%X%%%X",
				0xe0+((c&0xf000)>>12),
				0x80+((c&0xfc0)>>6),
				0x80+(c&0x3f),
			)
		} else { // double byte
			result += fmt.Sprintf("%%%X%%%X",
				0xc0+((c&0x7c0)>>6),
				0x80+(c&0x3f),
			)
		}
	}

	return result
}

func Canonical(src string) string {
	// ToLower
	src = Lower(src)

	return src
}

func CanonicalEmail(email string) string {
	// Canonical
	email = Canonical(email)

	// googlemail.com => gmail.com
	email = Replace(email, "googlemail.com", "gmail.com")
	// Remove Space and "
	email = Replace(email, " ", "")
	email = Replace(email, "\"", "")

	b := StringToBytes(email)
	ret := make([]byte, 0, len(b))

	plus := false
	atmark := false
	for i := 0; i < len(b); i++ {
		if b[i] == 0x2B { // 0x2B => "+"
			plus = true
		}
		if b[i] == 0x40 { // 0x40 => "@"
			plus = false
			atmark = true
		}
		if atmark {
			ret = append(ret, b[i])
		} else {
			if !plus && b[i] != 0x5F && b[i] != 0x2E { // 0x5F => "_" / 0x2E => "."
				ret = append(ret, b[i])
			}
		}
	}

	return BytesToString(ret)
}

var kanaConv unicode.SpecialCase = unicode.SpecialCase{
	// ひらがなをカタカナに変換
	unicode.CaseRange{
		0x3041, // Lo: ぁ
		0x3093, // Hi: ん
		[unicode.MaxCase]rune{
			0x30a1 - 0x3041, // UpperCase でカタカナに変換
			0,               // LowerCase では変換しない
			0x30a1 - 0x3041, // TitleCase でカタカナに変換
		},
	},
	// カタカナをひらがなに変換
	unicode.CaseRange{
		0x30a1, // Lo: ァ
		0x30f3, // Hi: ン
		[unicode.MaxCase]rune{
			0,               // UpperCase では変換しない
			0x3041 - 0x30a1, // LowerCase でひらがなに変換
			0,               // TitleCase では変換しない
		},
	},
}

func Katakana2Hiragana(src string) string {
	return strings.ToLowerSpecial(kanaConv, src)
}
