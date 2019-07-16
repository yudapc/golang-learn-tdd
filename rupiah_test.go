package tdd

import "testing"

func TestFormatRupiah(t *testing.T) {
	amount := float64(10000)
	expected := "Rp 10.000"
	res := FormatRupiah(amount)
	if res != expected {
		t.Errorf("Our FormatRupiah function doesn't work, %f isn't %s\n", amount, res)
	}
}

func TestGetAmount(t *testing.T) {
	rupiah := "Rp 10.000"
	expected := int64(10000)
	res := GetAmount(rupiah)
	if res != expected {
		t.Errorf("Our GetAmount function doesn't work, %s isn't %d\n", rupiah, res)
	}
}
