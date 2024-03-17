package gonatsort

import (
	"reflect"
	"sort"
	"testing"
)

func TestGoNatSort(t *testing.T) {
	strs := []string{"bar2", "foo", "foo123", "foo2", "bar001", "123", "15"}
	expected_strs := []string{"15", "123", "bar001", "bar2", "foo", "foo2", "foo123"}
	sort.Stable(NatSorter(strs))
	if !reflect.DeepEqual(strs, expected_strs) {
		t.Errorf("Failed. Expected sorted order to be %v, but was %v", expected_strs, strs)
	}
}

// goos: linux
// goarch: amd64
// pkg: github.com/wasim-nihal/gonatsort
// cpu: 11th Gen Intel(R) Core(TM) i5-1145G7 @ 2.60GHz
// BenchmarkGoNatSort
// BenchmarkGoNatSort-8     5921215               194.3 ns/op            24 B/op          1 allocs/op
// PASS
// ok      github.com/wasim-nihal/gonatsort        1.367s
func BenchmarkGoNatSort(b *testing.B) {
	strs := []string{"bar2", "foo", "foo123", "foo2", "bar001", "123", "15"}
	for i := 0; i < b.N; i++ {
		sort.Stable(NatSorter(strs))
	}
	b.ReportAllocs()
}
