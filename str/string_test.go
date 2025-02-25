package str

import "testing"

var strs string = `Golang is an open source programming language that makes it easy to build simple, reliable, and efficient software.`

func str(str string) {
	_ = str + "golang"
}

func ptr(str *string) {
	_ = *str + "golang"
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str(strs)
	}
}

func BenchmarkStringPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptr(&strs)
	}
}

func TestStringLeftEq(t *testing.T) {
	str := "abcdef"
	substr := "abc"
	if !StringLeftEq(str, substr) {
		t.Errorf("str=%s, substr=%s, ss=%s", str, substr, str[:len(substr)])
	}
}

func TestStringRightEq(t *testing.T) {
	str := "abcdef"
	substr := "def"
	if !StringRightEq(str, substr) {
		t.Errorf("str=%s, substr=%s, ss=%s", str, substr, str[len(str)-len(substr):len(str)])
	}
}

func TestEndsWith(t *testing.T) {
	str := "abc.h"
	comp := ".h"
	if !EndsWith(str, comp) {
		t.Errorf("str=%s, comp=%s", str, comp)
	}
}
