package string551_test

import (
	"github.com/go51/string551"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	src := "テスト用文字列"
	ret := string551.StringToBytes(src)
	sample := []byte(src)

	if len(ret) != len(sample) {
		t.Errorf("[]Byte への変換に失敗しました。\nData: %s\nRet: %#v\n", src, ret)
	}

	for i := 0; i < len(ret); i++ {
		if ret[i] != sample[i] {
			t.Errorf("[]Byte への変換に失敗しました。\nData: %s\nIndex: %d\nRet: %#v\n", src, i, ret[i])
		}
	}

}

func BenchmarkStringToBytesNormal(b *testing.B) {
	src := "テスト用文字列"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(src)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	src := "テスト用文字列"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.StringToBytes(src)
	}
}

func TestBytesToString(t *testing.T) {
	src := "テスト用文字列"
	retBytes := string551.StringToBytes(src)
	ret := string551.BytesToString(retBytes)

	if ret != src {
		t.Errorf("string への変換に失敗しました。\nData: %s\nSample: %#v\n", ret, src)
	}
}

func BenchmarkByteToStringNormal(b *testing.B) {
	src := "テスト用文字列"
	retBytes := string551.StringToBytes(src)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(retBytes)
	}
}

func BenchmarkByteToString(b *testing.B) {
	src := "テスト用文字列"
	retBytes := string551.StringToBytes(src)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.BytesToString(retBytes)
	}
}
