package kali

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerkalian(t *testing.T) {
	bil1 := 5
	bil2 := 4
	expected := 20
	actual := Perkalian(bil1, bil2)
	// if actual != expected {
	// 	t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
	// }
	assert.Equal(t, expected, actual, "hasil tidak sesuai.")
}
