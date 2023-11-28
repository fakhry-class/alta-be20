package main

import "testing"

func TestPenjumlahan(t *testing.T) {
	bil1 := 5
	bil2 := 4
	expected := 9
	actual := Penjumlahan(bil1, bil2)
	if actual != expected {
		t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
	}
}

func TestPenjumlahanBilNegatif(t *testing.T) {
	bil1 := -5
	bil2 := 4
	expected := -1
	actual := Penjumlahan(bil1, bil2)
	if actual != expected {
		t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
	}
}
func TestPenjumlahan2BilNegatif(t *testing.T) {
	bil1 := -5
	bil2 := -4
	expected := -9
	actual := Penjumlahan(bil1, bil2)
	if actual != expected {
		t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
	}
}

func TestPengurangan(t *testing.T) {
	bil1 := 5
	bil2 := 4
	expected := 1
	actual := Pengurangan(bil1, bil2)
	if actual != expected {
		t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
	}
}
