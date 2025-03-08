package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	resultado := Repetir("a", 5)
	esperado := "aaaaa"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}

func ExampleRepetir() {
	repeticoes := Repetir("a", 5)
	fmt.Println(repeticoes)
	// Output: aaaaa
}
