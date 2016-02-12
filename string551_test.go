package string551_test

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go51/secure551"
	"github.com/go51/string551"
	"strings"
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

func TestHexBytesToString(t *testing.T) {

	prefix := "string551"
	sid := secure551.Hash()
	name := "user"
	key := prefix + ":" + sid + ":" + name

	hash := md5.New()
	bytes := string551.StringToBytes(key)
	hash.Write(bytes)
	hexBytes := hash.Sum(nil)
	result := hex.EncodeToString(hexBytes)
	ret := string551.HexBytesToString(hexBytes)

	if result != ret {
		t.Errorf("string への変換に失敗しました。\nKey: %s\nResult 1: %#v\nResult 2: %#v\n", key, result, ret)
	}

}

func BenchmarkHexBytesToStringNomal(b *testing.B) {
	prefix := "string551"
	sid := secure551.Hash()
	name := "user"

	hash := md5.New()
	bytes := string551.StringToBytes(prefix + ":" + sid + ":" + name)
	hash.Write(bytes)
	hexBytes := hash.Sum(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hex.EncodeToString(hexBytes)
	}

}

func BenchmarkHexBytesToString(b *testing.B) {
	prefix := "string551"
	sid := secure551.Hash()
	name := "user"

	hash := md5.New()
	bytes := string551.StringToBytes(prefix + ":" + sid + ":" + name)
	hash.Write(bytes)
	hexBytes := hash.Sum(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.HexBytesToString(hexBytes)
	}

}

func TestJoin(t *testing.T) {
	src := []string{"Test", ",", "String", ",", "Join", ",", "Comma"}
	ret := string551.Join(src...)
	ret = "Test" + "," + "String" + "," + "Join" + "," + "Comma"

	if ret != "Test,String,Join,Comma" {
		t.Errorf("文字列の結合に失敗しました。\nData: %s\nSample: %#v\n", ret, src)
	}

}

func BenchmarkJoinNormal(b *testing.B) {
	src := []string{"Test", "String", "Join", "Comma"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join(src, ",")
	}
}

func BenchmarkJoinPlus(b *testing.B) {
	src := []string{"Test", "String", "Join", "Comma"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = src[0] + "," + src[1] + "," + src[2] + "," + src[3]
	}
}

func BenchmarkJoin(b *testing.B) {
	src := []string{"Test", ",", "String", ",", "Join", ",", "Comma"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.Join(src...)
	}
}

func TestCamelCase(t *testing.T) {
	src := "test_camel_case_t__e__s__t__c__a__m__e__l__c__a__s__e__"
	camel := string551.CamelCase(src)

	if camel != "TestCamelCaseTESTCAMELCASE" {
		t.Errorf("キャメルケースへの変換に失敗しました。\nData: %s\nCamel: %s\n", src, camel)
	}
}

func BenchmarkCamelCase(b *testing.B) {
	src := "test_camel_case_t__e__s__t__c__a__m__e__l__c__a__s__e__"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.CamelCase(src)
	}
}

func TestSnakeCase(t *testing.T) {
	src := "TestSnakeCaseTESTSNAKECASE"
	snake := string551.SnakeCase(src)

	if snake != "test_snake_case_t_e_s_t_s_n_a_k_e_c_a_s_e" {
		t.Errorf("キャメルケースへの変換に失敗しました。\nData: %s\nSnake: %s\n", src, snake)
	}
}

func BenchmarkSnakeCase(b *testing.B) {
	src := "TestSnakeCaseTESTSNAKECASE"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.SnakeCase(src)
	}
}

func TestReplace(t *testing.T) {
	src := "TestSplitString"
	old := "Split"
	new := "Replace"

	ret := string551.Replace(src, old, new)
	if ret != "TestReplaceString" {
		t.Errorf("文字列の置換に失敗しました。\nData: %s\nOld: %s\nNew: %s\n", src, old, new)
	}

}

func BenchmarkReplaceNormal(b *testing.B) {
	src := "TestSplitString"
	old := "Split"
	new := "Replace"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Replace(src, old, new, -1)
	}

}

func BenchmarkReplace(b *testing.B) {
	src := "TestSplitString"
	old := "Split"
	new := "Replace"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.Replace(src, old, new)
	}

}

func TestSplit(t *testing.T) {
	src := "test_split_string"
	separate := "_"

	ret := string551.Split(src, separate)

	if ret[0] != "test" || ret[1] != "split" || ret[2] != "string" {
		t.Errorf("文字列の分割に失敗しました。\nData: %s\nSeparate: %s\nRet: %#v\n", src, separate, ret)

	}

}

func BenchmarkSplit(b *testing.B) {
	src := "test_split_string"
	separate := "_"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.Split(src, separate)
	}
}

func TestRight(t *testing.T) {
	src := "test_split_string"
	length1 := 3
	length2 := len(src)
	length3 := len(src) + 1

	ret1 := string551.Right(src, length1)

	if ret1 != "ing" {
		t.Errorf("文字列の取得（右）に失敗しました。\nData: %s\nLen: %d\nRet: %#v\n", src, length1, ret1)
	}

	ret2 := string551.Right(src, length2)

	if ret2 != src {
		t.Errorf("文字列の取得（右）に失敗しました。\nData: %s\nLen: %d\nRet: %#v\n", src, length2, ret2)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("指定文字数が超えているのに Panic が発生していません。\nData: %s\nLen: %d\n", src, length3)
		}
	}()

	_ = string551.Right(src, length3)

}

func BenchmarkRight(b *testing.B) {
	src := "test_split_string"
	length := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.Right(src, length)
	}
}

func TestLower(t *testing.T) {
	src := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := string551.Lower(src)

	if lower != "abcdefghijklmnopqrstuvwxyz" {
		t.Errorf("小文字への変換に失敗しました。\nData: %s\nRet: %#v\n", src, lower)
	}
}

func BenchmarkLower(b *testing.B) {
	src := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.Lower(src)
	}
}

func TestRightRightRune(t *testing.T) {
	src := "あいうえおかきくけこさしすせそたちつてと"
	length := 10
	suffix := "..."

	ret := string551.RightRune(src, length, suffix)

	if ret != "あいうえおかきくけこ..." {
		t.Errorf("文字列のカットに失敗しました。\nRet: %s                                     ", ret)
	}

	ret = string551.RightRune(src, length)

	if ret != "あいうえおかきくけこ" {
		t.Errorf("文字列のカットに失敗しました。\nRet: %s                                     ", ret)
	}
}

func BenchmarkLRightRune(b *testing.B) {
	src := "あいうえおかきくけこさしすせそたちつてと"
	length := 10
	suffix := "..."

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string551.RightRune(src, length, suffix)
	}
}
