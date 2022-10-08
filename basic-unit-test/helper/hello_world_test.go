package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Fatih")
	}
}

func BenchmarkHelloWorldKhoiri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Khoiri")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Fatih", func(b *testing.B) {
		SayHello("Fatih")
	})

	b.Run("Khoiri", func(b *testing.B) {
		SayHello("Khoiri")
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{

		{
			name:    "Fatih",
			request: "Fatih",
		},
		{
			name:    "Khoiri",
			request: "Khoiri",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)
			}
		})
	}
}

func TestSayHello(t *testing.T) {
	result := SayHello("Fatih")
	if result != "Hello Fatih" {
		// t.Fail()
		t.Error("Result must be 'Hello Fatih'")
	}
	fmt.Println("Masih diekseskusi")
}

func TestSayHelloKhoiri(t *testing.T) {
	result := SayHello("Khoiri")
	if result != "Hello Khoiri" {
		// t.FailNow()
		t.Fatal("Result must be 'Hello Khoiri'")
	}

	fmt.Println("Tidak dieksekusi")
}

func TestHelloWorldAssert(t *testing.T) {
	result := SayHello("Fatih")
	assert.Equal(t, "Hello Fatih", result, "Result must be 'Hello Fatih'")
	fmt.Println("Test with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := SayHello("Fatih")
	require.Equal(t, "Hello Fatih", result, "Result must be 'Hello Fatih'")
	fmt.Println("Test with Require Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Used on mac os")
	}
	result := SayHello("Fatih")
	require.Equal(t, "Hello Fatih", result, "Result must be 'Hello Fatih'")
	fmt.Println("Test with Require Done")
}

func TestSubTest(t *testing.T) {
	t.Run("Fatih", func(t *testing.T) {
		result := SayHello("Fatih")
		assert.Equal(t, "Hello Fatih", result, "Result must be 'Hello Fatih'")
	})
	t.Run("Khoiri", func(t *testing.T) {
		result := SayHello("Khoiri")
		assert.Equal(t, "Hello Khoiri", result, "Result must be 'Hello Khoiri'")
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fatih",
			request:  "Fatih",
			expected: "Hello Fatih",
		},
		{
			name:     "Khoiri",
			request:  "Khoiri",
			expected: "Hello Khoiri",
		},
		{
			name:     "Andi",
			request:  "Andi",
			expected: "Hello Andi",
		},
		{
			name:     "Joham",
			request:  "Johan",
			expected: "Hello Johans",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum run test main")

	m.Run()

	fmt.Println("Setelah menjalankan unit test")
}
