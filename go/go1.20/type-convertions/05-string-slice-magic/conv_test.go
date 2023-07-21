package conv

import (
	"testing"
)

type s2ber interface {
	s2b(string) []byte
}

type b2ser interface {
	b2s([]byte) string
}

func Test_s2b_b2s_modify(t *testing.T) {
	for _, tt := range []struct {
		name string
		conv interface {
			b2ser
			s2ber
		}
	}{
		{"v1", v1{}},
		{"v2", v2{}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			v := "hello"
			b := []byte(v)
			s := tt.conv.b2s(b)
			if s != v {
				t.Errorf("got %q, want %q", s, v)
			}

			// NOTE: modify b, will also modify s
			// b[0] = 'H'
			// if s != v {
			// 	t.Errorf("after modified, got %q, want %q", s, v)
			// }

			expect := []byte(v)
			b = tt.conv.s2b(v)
			if !equalBytes(b, expect) {
				t.Errorf("got %v, want %v", b, expect)
			}

			// NOTE: modify b will got fatal!
			// oldV := v
			// b[0] = 'H' // fatal error: fault, unexpected fault address 0x10499f345
			// if v != oldV {
			// 	t.Errorf("after modified, got %q, want %q", v, oldV)
			// }
		})
	}
}

func Test_s2b_b2s(t *testing.T) {
	for _, tt := range []struct {
		name string
		conv interface {
			s2ber
			b2ser
		}
	}{
		{"noMagic", noMagic{}},
		{"v1", v1{}},
		{"v2", v2{}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range []string{
				"",
				"hello",
			} {
				expect := []byte(v)
				b := tt.conv.s2b(v)
				if !equalBytes(b, expect) {
					t.Errorf("got %v, want %v", b, expect)
				}

				s := tt.conv.b2s(b)
				if s != v {
					t.Errorf("got %q, want %q", s, v)
				}
			}
		})
	}
}

func Benchmark_s2b(b *testing.B) {
	v := "hello!"
	for _, tt := range []struct {
		name string
		conv s2ber
	}{
		{"noMagic", noMagic{}},
		{"v1", v1{}},
		{"v2", v2{}},
	} {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tt.conv.s2b(v)
			}
		})
	}
}

func Benchmark_b2s(b *testing.B) {
	v := []byte("hello!")
	for _, tt := range []struct {
		name string
		conv b2ser
	}{
		{"noMagic", noMagic{}},
		{"v1", v1{}},
		{"v2", v2{}},
	} {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tt.conv.b2s(v)
			}
		})
	}
}

func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return true
}
