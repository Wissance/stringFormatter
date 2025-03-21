package stringFormatter_test

import (
	"fmt"
	"github.com/wissance/stringFormatter"
	"testing"
)

func BenchmarkSliceToStringAdvancedWith8IntItems(b *testing.B) {
	slice := []any{100, 200, 300, 400, 500, 600, 700, 800}
	separator := ","
	for i := 0; i < b.N; i++ {
		_ = stringFormatter.SliceToString(&slice, &separator)
	}
}

func BenchmarkSliceStandard8IntItems(b *testing.B) {
	slice := []any{100, 200, 300, 400, 500, 600, 700, 800}
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%+q", slice)
	}
}

func BenchmarkSliceToStringAdvanced10MixedItems(b *testing.B) {
	slice := []any{100, "200", 300, "400", 500, 600, "700", 800, 1.09, "hello"}
	separator := ","
	for i := 0; i < b.N; i++ {
		_ = stringFormatter.SliceToString(&slice, &separator)
	}
}

func BenchmarkSliceStandard10MixedItems(b *testing.B) {
	slice := []any{100, "200", 300, "400", 500, 600, "700", 800, 1.09, "hello"}
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%+q", slice)
	}
}

func BenchmarkSliceToStringAdvanced20StrItems(b *testing.B) {
	slice := []any{"str1", "str2", "str3", "str4", "str5", "str6", "str7", "str8", "str9", "str10",
		"str11", "str12", "str13", "str14", "str15", "str16", "str17", "str18", "str19", "str20"}
	for i := 0; i < b.N; i++ {
		_ = stringFormatter.Format("{0:L,}", []any{slice})
	}
}

func BenchmarkSliceToStringAdvanced20TypedStrItems(b *testing.B) {
	slice := []string{"str1", "str2", "str3", "str4", "str5", "str6", "str7", "str8", "str9", "str10",
		"str11", "str12", "str13", "str14", "str15", "str16", "str17", "str18", "str19", "str20"}
	separator := ","
	for i := 0; i < b.N; i++ {
		_ = stringFormatter.SliceSameTypeToString(&slice, &separator)
	}
}

func BenchmarkSliceStandard20StrItems(b *testing.B) {
	slice := []any{"str1", "str2", "str3", "str4", "str5", "str6", "str7", "str8", "str9", "str10",
		"str11", "str12", "str13", "str14", "str15", "str16", "str17", "str18", "str19", "str20"}
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%+q", slice)
	}
}
