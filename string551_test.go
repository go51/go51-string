package string551_test

import (
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
