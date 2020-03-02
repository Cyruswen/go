package TestString

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "中,华,人,民,共,和,国"
	//for _, v := range s {
	//	t.Logf("%[1]c, %[1]x", v)
	//}
	parts := strings.Split(s, ",") //parts是切片
	for k, v := range parts {
		t.Log(k, v)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi(s); err == nil {
		t.Log(10 + i)
	}
}
