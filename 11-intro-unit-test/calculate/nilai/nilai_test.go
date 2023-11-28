package nilai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNilaiSiswa(t *testing.T) {
	t.Run("test nilai 85", func(t *testing.T) {
		nilai := 85
		expected := "A"
		actual := CheckNilaiSiswa(nilai)
		assert.Equal(t, expected, actual, "error. hasil tidak sesuai")
	})

	t.Run("test nilai 70", func(t *testing.T) {
		nilai := 70
		expected := "B+"
		actual := CheckNilaiSiswa(nilai)
		assert.Equal(t, expected, actual, "error. hasil tidak sesuai")
	})
	t.Run("test nilai 50", func(t *testing.T) {
		nilai := 50
		expected := "B"
		actual := CheckNilaiSiswa(nilai)
		assert.Equal(t, expected, actual, "error. hasil tidak sesuai")
	})
	t.Run("test nilai E", func(t *testing.T) {
		nilai := 0
		expected := "E"
		actual := CheckNilaiSiswa(nilai)
		assert.Equal(t, expected, actual, "error. hasil tidak sesuai")
	})
}
