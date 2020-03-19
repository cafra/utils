package LRU

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMemCache(t *testing.T) {

	t.Log(1 << 10)
	tests := []struct {
		name string
		want ICache
	}{
		// TODO: Add test cases.
		{
			name: "test",
			want: GetMemCache(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMemCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemCache_Get(t *testing.T) {
	c := GetMemCache(WithCapability(3))
	c.Set("k1", 7)
	c.Set("k2", 0)
	c.Set("k3", 1)
	c.Set("k4", 2)
	c.Get("k2")
	c.Set("k5", 3)
	c.Iterate(func(key string, value interface{}) {
		t.Logf("Key:%v	Value:%v", key, value)
	})
}

//BenchmarkSet-4   	 1217126	       949 ns/op	     208 B/op	       9 allocs/op
func BenchmarkSet(b *testing.B) {
	c := GetMemCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprint(i), i)
	}
}

//BenchmarkGet/set-4         	 1000000	      1053 ns/op
//BenchmarkGet/get-4         	 6941776	       168 ns/op
func BenchmarkGet(b *testing.B) {
	c := GetMemCache()
	b.ReportAllocs()

	b.Run("set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(fmt.Sprint(i), i)
		}
	})
	b.Run("get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Get(fmt.Sprint(i))
		}
	})
}
